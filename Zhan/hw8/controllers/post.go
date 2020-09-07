package controllers

import (
	"encoding/json"
	"fmt"
	"hw8/models"

	"github.com/astaxie/beego"
)

// PostController - стандартная структура
type PostController struct {
	beego.Controller
}

// Get - получение списка постов
func (p *PostController) Get() {

	posts := []models.Post{}

	posts, err := models.GetAll(models.Ctx, models.Db)
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
	idString := p.Ctx.Input.Param(":id")
	id := idString[10 : len(idString)-2]

	post, err := models.GetOne(models.Ctx, models.Db, id)
	if err != nil {
		p.Ctx.Output.SetStatus(400)
		p.Ctx.Output.Body([]byte("No posts in DB with current ID"))
		return
	}

	p.Data["Id"] = id
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
	post := &models.Post{
		Header: req.Header,
		Text:   req.Text,
	}

	_, err := post.Insert(models.Ctx, models.Db)
	if err != nil {
		p.Ctx.Output.SetStatus(400)
		p.Ctx.Output.Body([]byte("DB insert error"))
		return
	}

	p.Redirect("/", 301)
}

// Delete - обработчик удаления поста
func (p *PostController) Delete() {
	id := p.Ctx.Input.Param(":id")

	post, err := models.GetOne(models.Ctx, models.Db, id)
	if err != nil {
		p.Ctx.Output.SetStatus(400)
		p.Ctx.Output.Body([]byte("No posts in DB with current ID"))
		return
	}

	_, err = post.Delete(models.Ctx, models.Db)
	if err != nil {
		p.Ctx.Output.SetStatus(400)
		p.Ctx.Output.Body([]byte("DB delete error"))
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

	if err := json.Unmarshal(p.Ctx.Input.RequestBody, &req); err != nil {
		p.Ctx.Output.SetStatus(400)
		p.Ctx.Output.Body([]byte("Body is empty"))
		return
	}

	post, err := models.GetOne(models.Ctx, models.Db, id)
	if err != nil {
		p.Ctx.Output.SetStatus(400)
		p.Ctx.Output.Body([]byte("No posts in DB with current ID"))
		return
	}

	post.Header = req.Header
	post.Text = req.Text

	_, err = post.Update(models.Ctx, models.Db)
	if err != nil {
		fmt.Println(err)
		p.Ctx.Output.SetStatus(400)
		p.Ctx.Output.Body([]byte("Error updating post in BD"))
		return
	}
	p.Data["json"] = post
	p.ServeJSON()
}
