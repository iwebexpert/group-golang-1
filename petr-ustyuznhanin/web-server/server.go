package main

import (
	"flag"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func handleGetQuery(w http.ResponseWriter, r *http.Request){
	query := flag.String("q", "empty", "query to find")
	pages := flag.String("p", "http://ya.ru", "pages for search")
	flag.Parse()

	splitPages := strings.Split(*pages, ",")

	matches := []string{}

	for _, p := range splitPages {
		res, err := http.Get(p)
		if err != nil {
			log.Fatal(err)
		}
		defer res.Body.Close()

		bytes, err := ioutil.ReadAll(res.Body)
		if err != nil {
			log.Fatal(err)
		}

		if strings.Contains(string(bytes), *query){
			matches = append(matches, p)
			w.Write([]byte(p))
		}
	}
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", handleGetQuery)

	log.Fatal(http.ListenAndServe(":8080", r))
}
