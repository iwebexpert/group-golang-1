package main

import (
	"log"
	"net/http"
	"path"
	"text/template"

	"github.com/Toringol/group-golang-1/tree/master/s_shepelev/blog/tools"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

var (
	layoutTemplate = template.Must(template.ParseFiles(path.Join("../templates", "layout.html"),
		path.Join("../templates", "post.html")))
	posts = tools.NewPostArray()
)

func getAllPostsHandlers(w http.ResponseWriter, r *http.Request) {
	if err := layoutTemplate.ExecuteTemplate(w, "layout", posts); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func getPostHandler(w http.ResponseWriter, r *http.Request) {

}

func main() {
	err := posts.InitPosts()
	if err != nil {
		log.Fatal(err)
	}

	router := chi.NewRouter()
	router.Use(middleware.Logger)

	router.Get("/", getAllPostsHandlers)
	router.Get("/post/{postNumber}", getPostHandler)

	log.Println("Server start")

	err = http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatal(err)
	}
}
