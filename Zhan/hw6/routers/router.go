package routers

import (
	"hw6/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.PostController{})
	beego.Router("/post/:id([)A-Za-z0-9(\"]+)", &controllers.PostController{}, "get:GetOnePost;delete:Delete;put:Put")
	beego.Router("/newpost", &controllers.PostController{}, "get:NewPost")
}
