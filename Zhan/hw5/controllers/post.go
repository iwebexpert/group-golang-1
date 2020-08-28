package controllers

import (
	"encoding/json"
	"fmt"
	"hw5/models"
	"log"
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
		log.Println(err)
	}

	p.Data["Title"] = "My blog"
	p.Data["Posts"] = posts
	p.TplName = "allPosts.tpl"
}

// SelectPost - отображение выбранного поста
func (p *PostController) SelectPost() {
	id := p.Ctx.Input.Param(":id")
	uid64, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		p.Ctx.Output.SetStatus(400)
		p.Ctx.Output.Body([]byte("Post id is incorrect"))
		return
	}

	o := orm.NewOrm()
	post := models.Posts{ID: uid64}
	err2 := o.QueryTable("posts").Filter("ID", uid64).One(&post)
	if err2 != nil {
		p.Ctx.Output.SetStatus(400)
		p.Ctx.Output.Body([]byte("Post id is incorrect"))
		return
	}

	p.Data["Post"] = post
	p.TplName = "selectPost.tpl"
}

// Post - создание задачи
func (p *PostController) Post() {
	req := struct {
		Text string `json:"text"`
	}{}
	fmt.Println(p.Ctx.Input.RequestBody)
	if err := json.Unmarshal(p.Ctx.Input.RequestBody, &req); err != nil {
		p.Ctx.Output.SetStatus(400)
		p.Ctx.Output.Body([]byte("Body is empty"))
		return
	}

	post, err := models.NewPost(req.Text)
	if err != nil {
		p.Ctx.Output.SetStatus(400)
		p.Ctx.Output.Body([]byte(err.Error()))
		return
	}

	o := orm.NewOrm()
	id, err := o.Insert(post)
	if err != nil {
		p.Ctx.Output.SetStatus(400)
		p.Ctx.Output.Body([]byte("SQL insert error"))
		return
	}
	_ = id

	p.Data["json"] = post
	p.ServeJSON()
}
