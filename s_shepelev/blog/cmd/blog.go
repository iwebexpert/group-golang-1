package main

import (
	"html/template"
	"log"
	"path"

	"github.com/Toringol/group-golang-1/tree/master/s_shepelev/blog/views"

	bloghttp "github.com/Toringol/group-golang-1/tree/master/s_shepelev/blog/app/blog/delivery/http"
	"github.com/Toringol/group-golang-1/tree/master/s_shepelev/blog/app/blog/repository"
	"github.com/Toringol/group-golang-1/tree/master/s_shepelev/blog/app/blog/usecase"
	"github.com/Toringol/group-golang-1/tree/master/s_shepelev/blog/config"
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/viper"
)

func main() {
	if err := config.InitConfig(); err != nil {
		log.Fatalf("%s", err.Error())
	}

	listenAddr := viper.GetString("listenAddr")

	templates := make(map[string]*template.Template)

	templates["posts"] = template.Must(template.ParseFiles(path.Join("../views", "layout.html"),
		path.Join("../views", "posts.html")))
	templates["post"] = template.Must(template.ParseFiles(path.Join("../views", "layout.html"),
		path.Join("../views", "post.html")))
	templates["newPost"] = template.Must(template.ParseFiles(path.Join("../views", "layout.html"),
		path.Join("../views", "newPost.html")))

	e := echo.New()

	e.Renderer = &views.TemplateRegistry{
		Templates: templates,
	}

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${time_rfc3339} [${method}] ${remote_ip}, ${uri} ${status} 'error':'${error}'\n",
	}))

	bloghttp.NewBlogHandler(e, usecase.NewBlogUsecase(repository.NewBlogMemoryRepository()))

	e.Logger.Fatal(e.Start(listenAddr))
}
