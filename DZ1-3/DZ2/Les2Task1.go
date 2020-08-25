package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type Poisk struct {
	search         string
	SitesForSearch []string
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/json", foo)
	http.ListenAndServe(":8080", mux)
	log.Println("Starting server on address")
}

func foo(w http.ResponseWriter, r *http.Request) {
	query := Poisk{"Поиск", []string{"https://ya.ru", "https://google.com"}}
	fmt.Println("looking for: ", *query.search)
	fmt.Println("looking at: ", query.SitesForSearch)
	by1page := []string{}
	for _, p := range query.SitesForSearch {
		response, _ := http.Get(p)
		defer response.Body.Close()

		bytesArray, _ := ioutil.ReadAll(response.Body)

		if strings.Contains(string(bytesArray), *query.search) {
			by1page = append(by1page, p)
		}
	}
	answer := Poisk{
		search:         *query.search,
		SitesForSearch: by1page,
	}
	js, err := json.Marshal(answer)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
