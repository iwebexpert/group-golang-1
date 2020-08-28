package routers

import (
	"Lesson5/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/posts", &controllers.PostController{})
	beego.Router("/post/:id([0-9]+)", &controllers.PostController{}, "get:GetOnePost;put:UpdatePost;delete:DeletePost")
}
