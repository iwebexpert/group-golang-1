package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/go-chi/chi"
)

type TaskPage struct {
	Title string
	Posts Posts
}
type DetailedPage struct {
	Title string
	Post  PostType
}
type PostType struct {
	ID     int
	Header string
	Text   string
	Date   time.Time
}

type Posts []PostType

func (TaskPage *TaskPage) GetIndexHandler(w http.ResponseWriter, r *http.Request) {
	file, _ := os.Open("./www/index.html")
	data, _ := ioutil.ReadAll(file)

	templateMain := template.Must(template.New("PostList").Parse(string(data)))
	templateMain.ExecuteTemplate(w, "PostList", TaskPage)
}
func (TaskPage *TaskPage) GetDetailHandler(w http.ResponseWriter, r *http.Request) {
	file, _ := os.Open("./www/deatailPage.html")
	data, _ := ioutil.ReadAll(file)

	PostID, _ := strconv.Atoi(r.URL.Query().Get("ID"))
	dPost := TaskPage.Posts[PostID-1]
	DPage := DetailedPage{
		Title: dPost.Header,
		Post:  dPost,
	}
	templateMain := template.Must(template.New("Detail").Parse(string(data)))
	templateMain.ExecuteTemplate(w, "Detail", DPage)
}
func (TaskPage *TaskPage) GetPostHandler(w http.ResponseWriter, r *http.Request) {
	file, _ := os.Open("./www/PostPage.html")
	data, _ := ioutil.ReadAll(file)

	templateMain := template.Must(template.New("Post").Parse(string(data)))
	templateMain.ExecuteTemplate(w, "Post", page)
}

func (TaskPage TaskPage) NewPostHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println("New Post!")
	fmt.Println(r.Form)
	nHead := r.Form["Header"][0]
	nText := r.Form["text"][0]

	newID := len(page.Posts) + 1
	NewPost := PostType{
		ID:     newID,
		Text:   nText,
		Header: nHead,
		Date:   time.Now()}
	page.Posts = append(page.Posts, NewPost)
	http.Redirect(w, r, "/", 301)
}

var page = TaskPage{
	Title: "The Posts",
	Posts: Posts{
		{ID: 1, Header: "Post1", Text: "This is post number one!", Date: time.Date(2020, 8, 20, 10, 25, 21, 00, time.UTC)},
		{ID: 2, Header: "Interesting post", Text: "This is another interesting post", Date: time.Date(2020, 8, 21, 9, 50, 1, 01, time.UTC)},
	},
}

func main() {

	route := chi.NewRouter()
	route.Route("/", func(r chi.Router) {
		r.Get("/", page.GetIndexHandler)
		r.Get("/detail/", page.GetDetailHandler)
		r.Get("/post", page.GetPostHandler)
		r.Post("/post/newpost", page.NewPostHandler)
	})

	http.ListenAndServe(":8080", route)
}
