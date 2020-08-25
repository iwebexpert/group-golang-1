package main

import (
	"net/http"
	"log"
	"os"
	"os/signal"

	"searchserver/search"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Post("/search/{query}", SearchHandler)

	stopCh := make(chan os.Signal)

	go func() {
		log.Println("Server start")
		err := http.ListenAndServe(":8080", r)
		log.Fatal(err)
	}()

	signal.Notify(stopCh, os.Kill, os.Interrupt)
	<-stopCh

	log.Println("Server stop")
}

func SearchHandler(w http.ResponseWriter, r *http.Request) {
	query := chi.URLParam(r, "query")

	result := search.SearchSites(query)

	w.WriteHeader(http.StatusCreated)
	w.Write(result)
}