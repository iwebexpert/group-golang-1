package routers

import (
	"personalBlog/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/posts", &controllers.PostController{})
	beego.Router("/post/:id([0-9]+)", &controllers.PostController{}, "get:GetOnePost")
}
