package routers

import (
	"lesson5/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/tasks", &controllers.TaskController{})
	beego.Router("/task/:id([0-9]+)", &controllers.TaskController{}, "get:GetOneTask")
}
