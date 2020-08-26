package main

import (
	"io/ioutil"
	"net/http"
	"os"
	"text/template"

	"github.com/go-chi/chi"
)

type Server struct {
	Title string
	Posts PostItems
}

type PostItems []PostItem

type PostItem struct {
	Text   string
	Labels []string
}

func (server *Server) GetIndexHandler(w http.ResponseWriter, r *http.Request) {
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
	templateMain.ExecuteTemplate(w, "personalBlog", server)
}

func main() {
	route := chi.NewRouter()

	server := Server{
		Title: "Personal Blog!",
		Posts: PostItems{
			{Text: "Изучить Го", Labels: []string{"Go", "Lessons"}},
			{Text: "Create web-server", Labels: []string{"Go", "Server"}},
			{Text: "Xyz"},
		},
	}

	route.Route("/", func(r chi.Router) {
		r.Get("/", server.GetIndexHandler)
	})
	http.ListenAndServe(":8080", route)
}
