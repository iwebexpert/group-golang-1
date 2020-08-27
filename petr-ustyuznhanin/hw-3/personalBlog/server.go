package main

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"text/template"
)

type HomePage struct {
	Title string
	Posts Posts
}

type DetailedPage struct {
	Title string
	Post  PostItem
}

type PostItem struct {
	ID     int
	Header string
	Text   string
	Labels []string
}

type Posts []PostItem

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
	file, err := os.Open("./www/detailPage.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	data, err := ioutil.ReadAll(file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	PostID, err := strconv.Atoi(r.URL.Query().Get("ID"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	dPost := home.Posts[PostID-1]
	DPage := DetailedPage{
		Title: dPost.Header,
		Post:  dPost,
	}
	templateMain := template.Must(template.New("Detail").Parse(string(data)))
	templateMain.ExecuteTemplate(w, "Detail", DPage)
}

/*// PostPostHandler Создайте роут и шаблон для редактирования и создания материала.
func (home *HomePage) GetCreatePostHandler(w http.ResponseWriter, r *http.Request) {
	file, err := os.Open("./www/newPost.html")
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

func (home *HomePage) PostNewPostHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println("New Post!")
	fmt.Println(r.Form)
	nHead := r.Form["Header"][0]
	nText := r.Form["text"][0]

	newID := len(home.Posts) + 1
	NewPost := PostItem{
		ID:     newID,
		Text:   nText,
		Header: nHead,
	}
	home.Posts = append(home.Posts, NewPost)
	http.Redirect(w, r, "/", 301)
}*/

var home = HomePage{
	Title: "Personal Blog!",
	Posts: Posts{
		{ID: 1, Header: "Как я догонял группу", Text: "Жоска", Labels: []string{"кул стори", "попа в мыле"}},
		{ID: 2, Header: "Получилось?", Text: "Если ты это читаешь, то да", Labels: []string{"победа", "успех"}},
	},
}

func main() {
	route := chi.NewRouter()

	route.Use(middleware.Logger)

	route.Route("/", func(r chi.Router) {
		r.Get("/", home.GetIndexHandler)
		r.Get("/detailPage/", home.GetDetailPost)
		/*r.Get("/create", home.GetCreatePostHandler)
		r.Post("/new", home.PostNewPostHandler)*/
	})
	http.ListenAndServe(":8080", route)
}
