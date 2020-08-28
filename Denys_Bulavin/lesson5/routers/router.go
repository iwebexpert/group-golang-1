package routers

import (
	"lesson5/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/articles", &controllers.ArticleController{})
	beego.Router("/article/:id", &controllers.ArticleController{}, "get:GetOneArticle")
	beego.Router("/add", &controllers.ArticleController{}, "get:GetAddArticle;post:PostAddArticle")
	beego.Router("/article/delete/:id([0-9]+)", &controllers.ArticleController{}, "get:Delete")
	//	beego.Router("/del", &controllers.ArticleController{}, "")
}
