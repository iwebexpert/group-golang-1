package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"path"
	"text/template"

	"github.com/Toringol/group-golang-1/tree/master/s_shepelev/blog/tools"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

var (
	layoutTemplatePosts = template.Must(template.ParseFiles(path.Join("../templates", "layout.html"),
		path.Join("../templates", "posts.html")))
	layoutTemplateActualPost = template.Must(template.ParseFiles(path.Join("../templates", "layout.html"),
		path.Join("../templates", "post.html")))
	layoutTemplateAddPost = template.Must(template.ParseFiles(path.Join("../templates", "layout.html"),
		path.Join("../templates", "newpost.html")))
	posts = tools.NewPostArray()
)

type postText struct {
	Title string `json:"title"`
	Text  string `json:"text"`
}

func getAllPostsHandlers(w http.ResponseWriter, r *http.Request) {
	if err := layoutTemplatePosts.ExecuteTemplate(w, "layout", posts); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func getPostInfoHandler(w http.ResponseWriter, r *http.Request) {
	postName := chi.URLParam(r, "postName")
	postNamePrefix := "../posts/"
	postName = fmt.Sprint(postNamePrefix, postName)

	post, ok := posts.Items[postName]
	if !ok {
		http.Error(w, fmt.Errorf("No such post").Error(), http.StatusInternalServerError)
		return
	}

	if err := layoutTemplateActualPost.ExecuteTemplate(w, "layout", post); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func changePostHandler(w http.ResponseWriter, r *http.Request) {
	postName := chi.URLParam(r, "postName")
	postNamePrefix := "../posts/"
	postName = fmt.Sprint(postNamePrefix, postName)

	_, ok := posts.Items[postName]
	if !ok {
		http.Error(w, fmt.Errorf("No such post").Error(), http.StatusInternalServerError)
		return
	}

	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	msg := new(postText)

	if err := json.Unmarshal(b, msg); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := posts.ChangePost(postName, msg.Title, msg.Text); err != nil {
		http.Error(w, fmt.Errorf("No such post").Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func getNewPostHandler(w http.ResponseWriter, r *http.Request) {
	if err := layoutTemplateAddPost.ExecuteTemplate(w, "layout", ""); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func createNewPostHandler(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	msg := new(postText)

	if err := json.Unmarshal(b, msg); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := posts.CreatePost(msg.Title, msg.Text); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func main() {
	err := posts.InitPosts()
	if err != nil {
		log.Fatal(err)
	}

	router := chi.NewRouter()

	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	router.Get("/", getAllPostsHandlers)
	router.Get("/posts/{postName}", getPostInfoHandler)
	router.Post("/posts/{postName}", changePostHandler)
	router.Get("/addPost", getNewPostHandler)
	router.Post("/addPost", createNewPostHandler)

	log.Println("Server start")

	err = http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatal(err)
	}
}
