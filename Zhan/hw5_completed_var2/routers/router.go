package routers

import (
	"hw5/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.PostController{})
	beego.Router("/post/:id([0-9]+)", &controllers.PostController{}, "get:GetOnePost;delete:Delete;put:Put")
	beego.Router("/newpost", &controllers.PostController{}, "get:NewPost")
}
