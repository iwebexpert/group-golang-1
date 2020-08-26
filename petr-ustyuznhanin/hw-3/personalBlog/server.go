package main

import (
	"io/ioutil"
	"net/http"
	"os"
	"text/template"

	"github.com/go-chi/chi"
)

type HomePage struct {
	Title string
	Posts PostItems
}

type PostPage struct {
	Title string
	Post  PostItem
}

type PostItems []PostItem

type PostItem struct {
	ID     int
	Header string
	Text   string
	Labels []string
}

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

func main() {
	route := chi.NewRouter()

	home := HomePage{
		Title: "Personal Blog!",
		Posts: PostItems{
			{ID: 1, Header: "Как я догонял группу", Text: "Жоска", Labels: []string{"кул стори", "попа в мыле"}},
			{ID: 2, Header: "Получилось?", Text: "Если ты это читаешь, то да", Labels: []string{"победа", "успех"}},
		},
	}

	route.Route("/", func(r chi.Router) {
		r.Get("/", home.GetIndexHandler)
	})
	http.ListenAndServe(":8080", route)
}
