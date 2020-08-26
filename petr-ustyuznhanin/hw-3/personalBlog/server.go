package main

import (
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"text/template"

	"github.com/go-chi/chi"
)

type HomePage struct {
	Title string
	Posts Posts
}

type PostPage struct {
	Title string
	Post  PostItem
}

type Posts []PostItem

type PostItem struct {
	ID     int
	Header string
	Text   string
	Labels []string
}

// GetIndexHandler : 1. Создайте роут и шаблон для отображения всех постов в блоге.
func (home *HomePage) GetIndexHandler(w http.ResponseWriter, r *http.Request) {
	file, err := os.Open("./www/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	templateMain := template.Must(template.New("personalBlog").Parse(string(data)))
	templateMain.ExecuteTemplate(w, "personalBlog", home)
}

// GetDetailPost : 2. Создайте роут и шаблон для просмотра конкретного поста в блоге.
func (home *HomePage) GetDetailPost(w http.ResponseWriter, r *http.Request) {
	file, err := os.Open("./www/post.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	postID, err := strconv.Atoi(r.URL.Query().Get("ID"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	dPost := home.Posts[postID-1]
	DetailedPost := PostPage{
		Title: dPost.Header,
		Post:  dPost,
	}

	templateMain := template.Must(template.New("post").Parse(string(data)))
	templateMain.ExecuteTemplate(w, "post", DetailedPost)
}

// PostPostHamdler Создайте роут и шаблон для редактирования и создания материала.
func (home *HomePage) PostPostHandler(w http.ResponseWriter, r *http.Request) {

}

var home = HomePage{
	Title: "Personal Blog!",
	Posts: Posts{
		{ID: 1, Header: "Как я догонял группу", Text: "Жоска", Labels: []string{"кул стори", "попа в мыле"}},
		{ID: 2, Header: "Получилось?", Text: "Если ты это читаешь, то да", Labels: []string{"победа", "успех"}},
	},
}

func main() {
	route := chi.NewRouter()

	route.Route("/", func(r chi.Router) {
		r.Get("/", home.GetIndexHandler)
		r.Get("/detail/", home.GetDetailPost)
		r.Post("", home.PostPostHandler)
	})
	http.ListenAndServe(":8080", route)
}
