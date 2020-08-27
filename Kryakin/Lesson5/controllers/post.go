package controllers

import (
	"Lesson5/models"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type PostController struct {
	beego.Controller
}

func (c *PostController) Get() {
	beeOrm := orm.NewOrm()

	posts := []models.Posts{}

	_, err := beeOrm.QueryTable("Posts").All(&posts)
	if err != nil {
		c.Ctx.Output.SetStatus(400)
		c.Ctx.Output.Body([]byte("Error getting Posts from BD"))
	}
	c.Data["Title"] = "Test title"
	c.Data["Posts"] = posts
	c.TplName = "posts.tpl"
}

func (c *PostController) GetOnePost() {
	id := c.Ctx.Input.Param(":id")
	uid64, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.Ctx.Output.SetStatus(400)
		c.Ctx.Output.Body([]byte("Post id is incorrect1"))
		return
	}
	beeOrm := orm.NewOrm()
	post := models.Posts{Id: uid64}
	err2 := beeOrm.QueryTable("Posts").Filter("Id", uid64).One(&post)
	if err2 != nil {
		c.Ctx.Output.SetStatus(400)
		c.Ctx.Output.Body([]byte("Post id is incorrect2"))
		return
	}
	c.Data["Title"] = "Test title"
	c.Data["Posts"] = post
	c.TplName = "post.tpl"
}

func (c *PostController) Post() {
	req := struct {
		Header string `json: "header"`
		Text   string `json: "text"`
	}{}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &req); err != nil {
		c.Ctx.Output.SetStatus(400)
		c.Ctx.Output.Body([]byte("Body is empty"))
		return
	}

	post, err := models.NewPost(req.Header, req.Text)
	if err != nil {
		c.Ctx.Output.SetStatus(400)
		c.Ctx.Output.Body([]byte(err.Error()))
		return
	}

	beeOrm := orm.NewOrm()

	id, err := beeOrm.Insert(post)
	if err != nil {
		fmt.Println(err)
		c.Ctx.Output.SetStatus(400)
		c.Ctx.Output.Body([]byte("Error inserting post in BD"))
		return
	}
	_ = id
	c.Data["json"] = post
	c.ServeJSON()
}

func (c *PostController) UpdatePost() {
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

	post, err := models.ExPost(req.Header, req.Text, uid64)
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
func (c *PostController) DeletePost() {

	id := c.Ctx.Input.Param(":id")
	uid64, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.Ctx.Output.SetStatus(400)
		c.Ctx.Output.Body([]byte("Post id is incorrect1"))
		return
	}

	post, err := models.DelPost(uid64)
	if err != nil {
		c.Ctx.Output.SetStatus(400)
		c.Ctx.Output.Body([]byte(err.Error()))
		return
	}
	beeOrm := orm.NewOrm()
	fmt.Println(uid64)
	pid, err := beeOrm.Delete(post)
	if err != nil {
		c.Ctx.Output.SetStatus(400)
		c.Ctx.Output.Body([]byte("Error deleting post in BD"))
		return
	}
	_ = pid
	c.Data["json"] = post
	c.ServeJSON()
}
