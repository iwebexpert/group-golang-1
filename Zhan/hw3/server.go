package main

import (
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
)

type Server struct {
	Title string
	Posts PostItems
}

func (server *Server) allPosts(w http.ResponseWriter, r *http.Request) {
	file, _ := os.Open("./www/index.html")
	data, _ := ioutil.ReadAll(file)

	templateMain := template.Must(template.New("myBlog").Parse(string(data)))
	templateMain.ExecuteTemplate(w, "myBlog", server)
}

func (server *Server) selectedPost(w http.ResponseWriter, r *http.Request) {
	file, _ := os.Open("./www/2.html")
	data, _ := ioutil.ReadAll(file)

	templateMain := template.Must(template.New("2").Parse(string(data)))
	templateMain.ExecuteTemplate(w, "2", server)
}

func (server *Server) createPost(w http.ResponseWriter, r *http.Request) {
	file, _ := os.Open("./www/create.html")
	data, _ := ioutil.ReadAll(file)

	templateMain := template.Must(template.New("createPost").Parse(string(data)))
	templateMain.ExecuteTemplate(w, "createPost", server)
}

type PostItem struct {
	ID       uint
	Header   string
	Text     string
	Comments []string
}

type PostItems []PostItem

func main() {
	route := chi.NewRouter()

	server := Server{
		Title: "My blog",
		Posts: PostItems{
			{ID: 1, Header: "Мой первый пост", Text: "Жили-были дед да баба.", Comments: []string{"Супер!", "Интересное начало..."}},
			{ID: 2, Header: "Мой второй пост", Text: "И была у них Курочка Ряба.", Comments: []string{"Заинтриговал...", "Жги!"}},
			{ID: 3, Header: "Мой третий пост", Text: "Снесла курочка яичко, да не простое — золотое.", Comments: []string{"Вот это новость!", "Интересно, что же будет дальше..."}},
		},
	}

	route.Route("/", func(r chi.Router) {
		r.Get("/", server.allPosts)
		r.Get("/2", server.selectedPost)
		r.Get("/create", server.createPost)
	})

	log.Fatal(http.ListenAndServe(":8080", route))
}
