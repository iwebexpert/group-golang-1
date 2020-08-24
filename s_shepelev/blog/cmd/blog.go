package main

import (
	"log"
	"net/http"

	bloghttp "github.com/Toringol/group-golang-1/tree/master/s_shepelev/blog/app/blog/delivery/http"
	"github.com/Toringol/group-golang-1/tree/master/s_shepelev/blog/app/blog/repository"
	"github.com/Toringol/group-golang-1/tree/master/s_shepelev/blog/app/blog/usecase"
	"github.com/Toringol/group-golang-1/tree/master/s_shepelev/blog/config"
	"github.com/go-chi/chi"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
)

func main() {
	if err := config.InitConfig(); err != nil {
		log.Fatalf("%s", err.Error())
	}

	router := chi.NewRouter()

	bloghttp.NewBlogHandler(router, usecase.NewBlogUsecase(repository.NewBlogMemoryRepository()))

	log.Println("Server start")

	err := http.ListenAndServe(viper.GetString("listenPort"), router)
	if err != nil {
		log.Fatal(err)
	}
}
