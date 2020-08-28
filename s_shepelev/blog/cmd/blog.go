package main

import (
	"log"

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

	e := echo.New()

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${time_rfc3339} [${method}] ${remote_ip}, ${uri} ${status} 'error':'${error}'\n",
	}))

	bloghttp.NewBlogHandler(e, usecase.NewBlogUsecase(repository.NewBlogMemoryRepository()))

	e.Logger.Fatal(e.Start(listenAddr))
}
