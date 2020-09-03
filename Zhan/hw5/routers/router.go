package routers

import (
	"hw5/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.PostController{})
	beego.Router("/post/:id([0-9]+)", &controllers.PostController{}, "get:SelectPost")
}
