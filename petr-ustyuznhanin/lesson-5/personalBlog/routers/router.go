package routers

import (
	"personalBlog/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.PostController{})
	beego.Router("/newpost", &controllers.PostController{}, "get:NewPost")
	beego.Router("/post/:id([0-9]+)", &controllers.PostController{}, "get:GetOnePost")
}
