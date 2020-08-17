package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type searchReq struct {
	search string   `json:"search"`
	sites  []string `json:"sites"`
}

// var obj = []byte(`{"search":"programming","sites":["https://ya.ru","https://www.google.ru","https://mail.ru","https://www.rambler.ru","https://golang.org"]}`)

func postHandler(rw http.ResponseWriter, req *http.Request) {
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}

	defer req.Body.Close()

	log.Println(string(body))

	var SR searchReq

	err = json.Unmarshal(body, &SR)
	if err != nil {
		log.Panic(err)
	}

	log.Println(SR.search)
	log.Println(SR.sites)
}

func main() {
	router := http.NewServeMux()

	router.HandleFunc("/post", postHandler)

	log.Fatal(http.ListenAndServe(":8080", router))
}

// curl -X POST -H 'Content-Type: application/json' -d "{\"search\":\"programming\",\"sites\":[\"https://ya.ru\",\"https://www.google.ru\",\"https://mail.ru\",\"https://www.rambler.ru\",\"https://golang.org\"]}" http://localhost:8080/post
// curl -X POST -H 'Content-Type: application/json' -d "{\"search\":\"programming\",\"sites\":[\"https://ya.ru\",\"https://www.google.ru\",\"https://mail.ru\",\"https://www.rambler.ru\",\"https://golang.org\"]}" https://postman-echo.com/post
