package controllers

import (
	"fmt"
	"lesson5/models"
	"log"
	"strconv"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type ArticleController struct {
	beego.Controller
}

func (a *ArticleController) Get() {
	o := orm.NewOrm()

	articles := []models.Articles{}
	_, err := o.QueryTable("articles").All(&articles)
	if err != nil {
		log.Fatal(err)
	}

	a.Data["Title"] = "Test title"
	a.Data["Articles"] = articles
	a.TplName = "articles.tpl"
}

func (a *ArticleController) GetOneArticle() {
	id := a.Ctx.Input.Param(":id")
	uid64, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		a.Ctx.Output.SetStatus(400)
		a.Ctx.Output.Body([]byte("Article id is incorect"))
	}
	o := orm.NewOrm()

	article := models.Articles{Id: uid64}
	err2 := o.QueryTable("articles").Filter("Id", uid64).One(&article)
	if err2 != nil {
		a.Ctx.Output.SetStatus(400)
		a.Ctx.Output.Body([]byte("Article id is incorect"))
	}

	a.Data["Articles"] = article
	a.TplName = "article.tpl"
}

func (c *ArticleController) GetAddArticle() {
	c.TplName = "add.tpl"
}

func (c *ArticleController) PostAddArticle() {

	var req models.GetForm

	if err := c.ParseForm(&req); err != nil {
		c.Ctx.Output.SetStatus(400)
		c.Ctx.Output.Body([]byte("Body is empty"))
		return
	}

	article, err := models.NewArticle(&req)
	if err != nil {
		c.Ctx.Output.SetStatus(400)
		c.Ctx.Output.Body([]byte(err.Error()))
		return
	}

	o := orm.NewOrm()
	id, err := o.Insert(article)
	if err != nil {
		c.Ctx.Output.SetStatus(400)
		c.Ctx.Output.Body([]byte(err.Error()))
		return
	}

	_ = id

	c.Redirect("/articles", 301)
}

func (c *ArticleController) Delete() {
	id := c.Ctx.Input.Param(":id")
	uid64, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.Ctx.Output.SetStatus(400)
		c.Ctx.Output.Body([]byte("DELETE: Article id is incorect"))
	}
	o := orm.NewOrm()

	article := models.Articles{Id: uid64}

	if exist := o.QueryTable(article.TableName()).Filter("Id", uid64).Exist(); exist {
		if num, err := o.Delete(&models.Articles{Id: uid64}); err == nil {
			beego.Info("Record Deleted. ", num)
		} else {
			beego.Error("Record couldn't be deleted. Reason: ", err)
		}
	} else {
		beego.Info("Record Doesn't exist.")
	}

	c.Redirect("/articles", 301)
}

func (c *ArticleController) GetUpdateArticle() {
	id := c.Ctx.Input.Param(":id")
	uid64, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.Ctx.Output.SetStatus(400)
		c.Ctx.Output.Body([]byte("Article id is incorect"))
	}

	c.Data["id"] = uid64
	c.TplName = "updatepost.tpl"
}

func (c *ArticleController) PostUpdateArticle() {
	id := c.Ctx.Input.Param(":id")
	uid64, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.Ctx.Output.SetStatus(400)
		c.Ctx.Output.Body([]byte("Article id is incorect"))
	}

	o := orm.NewOrm()

	article := models.Articles{Id: uid64}

	if err := c.ParseForm(&article); err != nil {
		c.Ctx.Output.SetStatus(400)
		c.Ctx.Output.Body([]byte("Body is empty"))
		return
	}
	fmt.Println(article)

	num, err := o.Update(&article)
	if err != nil {
		c.Ctx.Output.SetStatus(400)
		c.Ctx.Output.Body([]byte(err.Error()))
		return
	}

	_ = num
	c.Redirect("/articles", 301)
	/*
		var req models.Articles

		if err := c.ParseForm(&req); err != nil {
			c.Ctx.Output.SetStatus(400)
			c.Ctx.Output.Body([]byte("Body is empty"))
			return
		}

		article, err := models.NewArticle(&req)
		if err != nil {
			c.Ctx.Output.SetStatus(400)
			c.Ctx.Output.Body([]byte(err.Error()))
			return
		}
		o := orm.NewOrm()
		id, err := o.Update(&article)
		if err != nil {
			c.Ctx.Output.SetStatus(400)
			c.Ctx.Output.Body([]byte(err.Error()))
			return
		}

		_ = id

		c.Redirect("/articles", 301)
	*/
}
