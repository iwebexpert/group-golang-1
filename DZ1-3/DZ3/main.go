package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"

	"github.com/go-chi/chi"
)

type MainPages struct {
	Title string
	Post  Posts
}

type Posts []PostStruct

type PostStruct struct {
	ID   int
	Name string
	Text string
}

type DetailedPage struct {
	Title string
	Post  PostStruct
}

var page = MainPages{
	Title: "Chat Of The Ring",
	Post: Posts{
		{ID: 1, Name: "Aragorn", Text: "Стобой будет мой меч!"},
		{ID: 2, Name: "Gimly", Text: "И моя секира!"},
		{ID: 3, Name: "Legolas", Text: "И мой лук!"},
		{ID: 4, Name: "Sauron", Text: "И моё кольцо!"},
	},
}

func main() {
	route := chi.NewRouter()
	route.Route("/", func(r chi.Router) {
		r.Get("/", page.GetIndexHandler)
		r.Get("/detail/", page.GetDetailHandler)
		r.Get("/post", page.GetPostHandler)
		r.Post("/post/newpost", page.PostPostsHandler)

	})
	http.ListenAndServe(":8080", route)
}

func (MainPage *MainPages) GetIndexHandler(w http.ResponseWriter, r *http.Request) {
	file, _ := os.Open("./www/index.html")
	data, _ := ioutil.ReadAll(file)
	tMain := template.Must(template.New("PostList").Parse(string(data)))
	//сама команда Must закладывает новосозданный шаблон "PostList" в шаблон tMain, я правильно понимаю?
	tMain.ExecuteTemplate(w, "PostList", MainPage)
}

func (MainPage *MainPages) GetDetailHandler(w http.ResponseWriter, r *http.Request) {
	file, _ := os.Open("./www/detailPage.html")
	data, _ := ioutil.ReadAll(file)
	PostID, _ := strconv.Atoi(r.URL.Query().Get("ID"))
	dPost := MainPage.Post[PostID-1]
	DPage := DetailedPage{
		Title: dPost.Name,
		Post:  dPost,
	}
	tMain := template.Must(template.New("Detail").Parse(string(data)))
	tMain.ExecuteTemplate(w, "Detail", DPage)
	//а можно тут напрямую передать dPost? я пробывал, но не работало, но может я что-то неправильно сделал когда пробывал
}

func (MainPage *MainPages) GetPostHandler(w http.ResponseWriter, r *http.Request) {
	file, _ := os.Open("./www/PostPage.html")
	data, _ := ioutil.ReadAll(file)
	tMain := template.Must(template.New("Post").Parse(string(data)))
	tMain.ExecuteTemplate(w, "Post", page)
}

func (MainPage *MainPages) PostPostsHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println(r.Form)
	newID := len(page.Post) + 1
	newName := r.Form["Name"][0]
	newText := r.Form["Text"][0]
	newPost := PostStruct{
		ID:   newID,
		Name: newName,
		Text: newText,
	}

	page.Post = append(page.Post, newPost)
	http.Redirect(w, r, "/", 301)
}
