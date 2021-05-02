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

func (c *PostController) Post() {
	req := struct {
		Text string `json:"text"`
	}{}
	fmt.Println(c.Ctx.Input.RequestBody)
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &req); err != nil {
		c.Ctx.Output.SetStatus(400)
		c.Ctx.Output.Body([]byte("Body is empty"))
		return
	}

	post, err := models.NewPost(req.Text)
	if err != nil {
		c.Ctx.Output.SetStatus(400)
		c.Ctx.Output.Body([]byte(err.Error()))
		return
	}

	beeOrm := orm.NewOrm()
	id, err := beeOrm.Insert(post)
	if err != nil {
		c.Ctx.Output.SetStatus(400)
		c.Ctx.Output.Body([]byte("SQL insert error"))
		return
	}
	_ = id

	c.Data["json"] = post
	c.ServeJSON()
}
