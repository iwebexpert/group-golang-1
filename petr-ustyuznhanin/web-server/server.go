package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type QueryT struct {
	Search string `json:"search"`
	Sites []string `json:"sites"`
}

var data QueryT

func QueryContains(s QueryT) ([]string, error){
	var answer []string
	for _, v := range s.Sites {
		res, err := http.Get(v)
		if err != nil {
			return nil, err
		}
		defer res.Body.Close()

		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return nil, err
		}

		if strings.Contains(string(body), s.Search); err != nil{
			answer = append(answer, v)
		}
	}
	return answer, nil
}

func GetQueryParams(w http.ResponseWriter, r *http.Request)  {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil { //не понимаю как в таких случаях обработать ошибку
		return
	}
	err = json.Unmarshal(body, &data)
	if err != nil { //не понимаю как в таких случаях обработать ошибку
		return
	}
}

func handleGetQuery (w http.ResponseWriter, r *http.Request) {
	QueryContains(data)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/search", handleGetQuery)
	log.Fatal(http.ListenAndServe(":8080", r))
}
