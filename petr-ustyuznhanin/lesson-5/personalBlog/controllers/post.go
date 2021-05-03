package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"personalBlog/models"
	"strconv"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type PostController struct {
	beego.Controller
}

// Get - получить все посты
func (c *PostController) Get() {
	beeOrm := orm.NewOrm()

	posts := []models.Posts{}

	_, err := beeOrm.QueryTable("posts").All(&posts)
	if err != nil {
		log.Fatal(err)
	}

	c.Data["Title"] = "My personal blog"
	c.Data["Posts"] = posts
	c.TplName = "posts.tpl"
}

func (c *PostController) GetOnePost() {
	id := c.Ctx.Input.Param(":id")
	uid64, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.Ctx.Output.SetStatus(400)
		c.Ctx.Output.Body([]byte("Post id is incorrect"))
		return
	}

	beeOrm := orm.NewOrm()
	post := models.Posts{Id: uid64}
	err2 := beeOrm.QueryTable("posts").Filter("Id", uid64).One(&post)
	if err2 != nil {
		c.Ctx.Output.SetStatus(400)
		c.Ctx.Output.Body([]byte("Post id is incorrect"))
		return
	}

	c.Data["Post"] = post
	c.TplName = "post.tpl"
}

// NewPost - перейти на страницу создания нового поста
func (p *PostController) NewPostPage() {
	p.Data["Title"] = "My blog"
	p.TplName = "newpost.tpl"
}

func (c *PostController) Post() {
	req := struct {
		Title string `json:"title"`
		Text  string `json:"text"`
	}{}
	if err := c.ParseForm(&req); err != nil {
		c.Ctx.Output.SetStatus(400)
		c.Ctx.Output.Body([]byte("Body is empty"))
		return
	}

	post, err := models.NewPost(req.Title, req.Text)
	if err != nil {
		c.Ctx.Output.SetStatus(400)
		c.Ctx.Output.Body([]byte(err.Error()))
		return
	}

	beeOrm := orm.NewOrm()
	_, err = beeOrm.Insert(post)
	if err != nil {
		c.Ctx.Output.SetStatus(400)
		c.Ctx.Output.Body([]byte("SQL insert error"))
		return
	}

	c.Redirect("/", 301)
	//c.Data["json"] = post
	//c.ServeJSON()
}

// Put - редактирование поста
func (c *PostController) Put() {
	req := struct {
		Header string `json: "header"`
		Text   string `json: "text"`
	}{}
	id := c.Ctx.Input.Param(":id")
	uid64, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.Ctx.Output.SetStatus(400)
		c.Ctx.Output.Body([]byte("Post id is incorrect1"))
		return
	}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &req); err != nil {
		c.Ctx.Output.SetStatus(400)
		c.Ctx.Output.Body([]byte("Body is empty"))
		return
	}

	post, err := models.UpdatePost(req.Header, req.Text, uid64)
	if err != nil {
		c.Ctx.Output.SetStatus(400)
		c.Ctx.Output.Body([]byte(err.Error()))
		return
	}
	beeOrm := orm.NewOrm()

	pid, err := beeOrm.Update(post)
	if err != nil {
		fmt.Println(err)
		c.Ctx.Output.SetStatus(400)
		c.Ctx.Output.Body([]byte("Error updating post in BD"))
		return
	}
	_ = pid
	c.Data["json"] = post
	c.ServeJSON()
}

func (c *PostController) Delete() {
	id := c.Ctx.Input.Param(":id")
	uid64, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.Ctx.Output.SetStatus(400)
		c.Ctx.Output.Body([]byte("Post id is incorrect"))
		return
	}

	beeOrm := orm.NewOrm()
	post := models.Posts{Id: uid64}
	_, err = beeOrm.Delete(&post)
	if err != nil {
		c.Ctx.Output.SetStatus(400)
		c.Ctx.Output.Body([]byte("SQL delete error"))
		return
	}
	c.Data["json"] = post
	c.ServeJSON()
}
