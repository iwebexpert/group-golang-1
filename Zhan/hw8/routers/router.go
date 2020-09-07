package routers

import (
	"hw8/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.PostController{})
	beego.Router("/post/:id([)A-Za-z0-9(\"]+)", &controllers.PostController{}, "get:GetOnePost")
	beego.Router("/post/:id([)A-Za-z0-9(\"]+)", &controllers.PostController{})
	beego.Router("/newpost", &controllers.PostController{}, "get:NewPost")
}
