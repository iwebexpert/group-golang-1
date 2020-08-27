package main

import (
	// "fmt"
	"html/template"
	"net/http"
	"log"
	"io/ioutil"
	"os"
	"os/signal"

	"snowrill/blog/core"

	"github.com/go-chi/chi"
)

type BlogServer struct {
	Title string
	core.BlogPosts
}

func (b *BlogServer) IndexHandler(w http.ResponseWriter, r *http.Request) {
	file, err := os.Open("static/index.html")
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	data, err := ioutil.ReadAll(file)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	err = b.LoadPostsFrom("static/data.json")
	if err != nil {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	templateIndex := template.Must(template.New("blogIndex").Parse(string(data)))
	templateIndex.ExecuteTemplate(w, "blogIndex", b)
}


func main() {
	stopCh := make(chan os.Signal)
	r := chi.NewRouter()

	blog := BlogServer{
		Title: "Simple blog",
	}

	r.Route("/", func(r chi.Router) {
		r.Get("/", blog.IndexHandler)
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