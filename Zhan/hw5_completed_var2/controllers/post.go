package controllers

import (
	"encoding/json"
	"fmt"
	"hw5/models"
	"strconv"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

// PostController - стандартная структура
type PostController struct {
	beego.Controller
}

// Get - получение списка постов
func (p *PostController) Get() {
	o := orm.NewOrm()

	posts := []models.Posts{}

	_, err := o.QueryTable("posts").All(&posts)
	if err != nil {
		p.Ctx.Output.SetStatus(400)
		p.Ctx.Output.Body([]byte(err.Error()))
		return
	}

	p.Data["Title"] = "My blog"
	p.Data["Posts"] = posts
	p.TplName = "posts.tpl"
}

// GetOnePost - отображение выбранного поста
func (p *PostController) GetOnePost() {
	id := p.Ctx.Input.Param(":id")
	uid64, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		p.Ctx.Output.SetStatus(400)
		p.Ctx.Output.Body([]byte("Post ID is incorrect"))
		return
	}

	o := orm.NewOrm()
	post := models.Posts{Id: uid64}
	err = o.QueryTable("posts").Filter("ID", uid64).One(&post)
	if err != nil {
		p.Ctx.Output.SetStatus(400)
		p.Ctx.Output.Body([]byte("No posts in DB with current ID"))
		return
	}

	p.Data["Post"] = post
	p.TplName = "post.tpl"
}

// NewPost - страница создания и редактирования поста
func (p *PostController) NewPost() {
	p.Data["Title"] = "My blog"
	p.TplName = "newPost.tpl"
}

// Post - обработчик создания нового поста
func (p *PostController) Post() {
	req := struct {
		Header string `json:"header"`
		Text   string `json:"text"`
	}{}

	if err := p.ParseForm(&req); err != nil {
		p.Ctx.Output.SetStatus(400)
		p.Ctx.Output.Body([]byte("Body is empty"))
		return
	}

	// здесь для создания и редактирования вызывается одна и та же функция в зависимости от ситуации
	// в данном случае для создания поста передаётся id=0, что равносильно NULL для БД MySQL
	// значит база присвоит id в соответствии с autoincrement
	post, err := models.NewUpdPost(req.Header, req.Text, 0)
	if err != nil {
		p.Ctx.Output.SetStatus(400)
		p.Ctx.Output.Body([]byte(err.Error()))
		return
	}

	o := orm.NewOrm()
	_, err = o.Insert(post)
	if err != nil {
		p.Ctx.Output.SetStatus(400)
		p.Ctx.Output.Body([]byte("SQL insert error"))
		return
	}

	p.Redirect("/", 301)
}

// Delete - обработчик удаления поста
func (p *PostController) Delete() {
	id := p.Ctx.Input.Param(":id")
	uid64, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		p.Ctx.Output.SetStatus(400)
		p.Ctx.Output.Body([]byte("Post id is incorrect"))
		return
	}

	o := orm.NewOrm()
	post := models.Posts{Id: uid64}
	_, err = o.Delete(&post)
	if err != nil {
		p.Ctx.Output.SetStatus(400)
		p.Ctx.Output.Body([]byte("SQL delete error"))
		return
	}
	p.Data["json"] = post
	p.ServeJSON()
}

// Put - обработчик редактирования поста
func (p *PostController) Put() {
	req := struct {
		Header string `json:"header"`
		Text   string `json:"text"`
	}{}
	id := p.Ctx.Input.Param(":id")
	uid64, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		p.Ctx.Output.SetStatus(400)
		p.Ctx.Output.Body([]byte("Post id is incorrect"))
		return
	}
	if err := json.Unmarshal(p.Ctx.Input.RequestBody, &req); err != nil {
		p.Ctx.Output.SetStatus(400)
		p.Ctx.Output.Body([]byte("Body is empty"))
		return
	}

	// здесь для создания и редактирования вызывается одна и та же функция в зависимости от ситуации
	// в данном случае для редактирования поста передаётся нужный нам id
	post, err := models.NewUpdPost(req.Header, req.Text, uid64)
	if err != nil {
		p.Ctx.Output.SetStatus(400)
		p.Ctx.Output.Body([]byte(err.Error()))
		return
	}
	o := orm.NewOrm()

	_, err = o.Update(post)
	if err != nil {
		fmt.Println(err)
		p.Ctx.Output.SetStatus(400)
		p.Ctx.Output.Body([]byte("Error updating post in BD"))
		return
	}
	p.Data["json"] = post
	p.ServeJSON()
}
