package main

import (
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
	BlogTitle string
	core.BlogPosts
}

func (b *BlogServer) IndexHandler(w http.ResponseWriter, r *http.Request) {
	b.Load(DBpath)
	data, err := loadTemplate("static/index.html")
	if err != nil {
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}


	temp := template.Must(template.New("blogIndex").Parse(data))
	temp.ExecuteTemplate(w, "blogIndex", b)
}

func (b *BlogServer) SinglePostHandler(w http.ResponseWriter, r *http.Request) {
	b.Load(DBpath)
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
			BlogTitle: b.BlogTitle,
			Post: post,
	})
}

func (b *BlogServer) LoadAddForm(w http.ResponseWriter, r *http.Request) {
	b.Load(DBpath)
	data, err := loadTemplate("static/addpost.html")
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	temp := template.Must(template.New("addPost").Parse(data))
	temp.ExecuteTemplate(w, "addPost", b)
}

func (b *BlogServer) LoadEditForm(w http.ResponseWriter, r *http.Request) {
	b.Load(DBpath)
	data, err := loadTemplate("static/editpost.html")
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

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

	temp := template.Must(template.New("editPost").Parse(data))
	temp.ExecuteTemplate(w, "editPost", struct{
		BlogTitle string
		Post 	  core.Post
	}{
		BlogTitle: b.BlogTitle,
		Post: post,
	})
}

func (b *BlogServer) EditPostHandler(w http.ResponseWriter, r *http.Request) {
	b.Load(DBpath)
	postId, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	r.ParseForm()
	title := r.Form.Get("title")
	text := r.Form.Get("text")

	err = b.EditPost(postId, title, text)
	if err != nil {
		return
	}

	b.Save(DBpath)
	http.Redirect(w, r, "/", 301)
}

func (b *BlogServer) AddPostHandler(w http.ResponseWriter, r *http.Request) {
	b.Load(DBpath)
	r.ParseForm()
	title := r.Form.Get("title")
	text := r.Form.Get("text")

	_, err := b.AddPost(title, text) 
	if err != nil {
		http.Redirect(w, r, "/posts/add", 301)
		return
	}

	b.Save(DBpath)
	http.Redirect(w, r, "/", 301)
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
		BlogTitle: "Simple blog",
	}
	blog.Load(DBpath)

	r.Get("/", blog.IndexHandler)

	r.Route("/posts/{id}", func(r chi.Router){
		r.Get("/", blog.SinglePostHandler)
		r.Get("/edit", blog.LoadEditForm)
		r.Post("/edit", blog.EditPostHandler)
	})


	r.Route("/posts/add", func(r chi.Router) {
		r.Get("/", blog.LoadAddForm)
		r.Post("/", blog.AddPostHandler)
	})

	go func() {
		log.Println("Server is running")
		err := http.ListenAndServe(":8080", r)
		log.Fatal(err)
	}()

	signal.Notify(stopCh, os.Kill, os.Interrupt)
	<-stopCh

	log.Println("Server stopped")
}