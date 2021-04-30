package controllers

import (
	"encoding/json"
	"fmt"
	"lesson5/models"
	"log"
	"strconv"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type TaskController struct {
	beego.Controller
}

func (c *TaskController) Get() {
	beeOrm := orm.NewOrm()

	tasks := []models.Tasks{}

	_, err := beeOrm.QueryTable("tasks").All(&tasks)
	if err != nil {
		log.Fatal(err)
	}

	c.Data["Title"] = "Test title"
	c.Data["Tasks"] = tasks
	c.TplName = "tasks.tpl"
}

func (c *TaskController) GetOneTask() {
	id := c.Ctx.Input.Param(":id")
	uid64, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.Ctx.Output.SetStatus(400)
		c.Ctx.Output.Body([]byte("Task id is incorrect"))
		return
	}

	beeOrm := orm.NewOrm()
	task := models.Tasks{Id: uid64}
	err2 := beeOrm.QueryTable("tasks").Filter("Id", uid64).One(&task)
	if err2 != nil {
		c.Ctx.Output.SetStatus(400)
		c.Ctx.Output.Body([]byte("Task id is incorrect"))
		return
	}

	c.Data["Task"] = task
	c.TplName = "task.tpl"
}

func (c *TaskController) Post() {
	req := struct {
		Text string `json:"text"`
	}{}
	fmt.Println(c.Ctx.Input.RequestBody)
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &req); err != nil {
		c.Ctx.Output.SetStatus(400)
		c.Ctx.Output.Body([]byte("Body is empty"))
		return
	}

	task, err := models.NewTask(req.Text)
	if err != nil {
		c.Ctx.Output.SetStatus(400)
		c.Ctx.Output.Body([]byte(err.Error()))
		return
	}

	beeOrm := orm.NewOrm()
	id, err := beeOrm.Insert(task)
	if err != nil {
		c.Ctx.Output.SetStatus(400)
		c.Ctx.Output.Body([]byte("SQL insert error"))
		return
	}
	_ = id

	c.Data["json"] = task
	c.ServeJSON()
}
