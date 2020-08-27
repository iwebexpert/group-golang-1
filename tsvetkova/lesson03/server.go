package main

import (
	// "fmt"
	"html/template"
	"net/http"
	"log"
	"io/ioutil"
	"os"
	"os/signal"
	"strconv"

	"snowrill/blog/core"

	"github.com/go-chi/chi"
)

const DBpath = "static/data.json"

type BlogServer struct {
	Title string
	core.BlogPosts
}

func (b *BlogServer) IndexHandler(w http.ResponseWriter, r *http.Request) {
	data, err := loadTemplate("static/index.html")
	if err != nil {
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	temp := template.Must(template.New("blogIndex").Parse(data))
	temp.ExecuteTemplate(w, "blogIndex", b)
}

func (b *BlogServer) SinglePostHandler(w http.ResponseWriter, r *http.Request) {
	
	postId, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	post, err := b.GetPostById(postId)
	if err != nil {
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	data, err := loadTemplate("static/singlepost.html")
	if err != nil {
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	temp := template.Must(template.New("singlePost").Parse(data))
	temp.ExecuteTemplate(w, "singlePost", struct {
		BlogTitle string
		Post core.Post
		} {
			BlogTitle: b.Title,
			Post: post,
	})
}


func loadTemplate(path string) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		return "", err
	}

	data, err := ioutil.ReadAll(file)
	if err != nil {
		return "", err
	}

	return string(data), nil
}

func main() {
	stopCh := make(chan os.Signal)
	r := chi.NewRouter()

	blog := BlogServer{
		Title: "Simple blog",
	}

	blog.LoadPostsFrom(DBpath)
	r.Get("/", blog.IndexHandler)
	r.Get("/post/{id}", blog.SinglePostHandler)

	go func() {
		log.Println("Server is running")
		err := http.ListenAndServe(":8080", r)
		log.Fatal(err)
	}()

	signal.Notify(stopCh, os.Kill, os.Interrupt)
	<-stopCh

	log.Println("Server stopped")
}