package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

// task1
type searchReq struct {
	Search string   `json:"search"`
	Sites  []string `json:"sites"`
}

var sr searchReq

func postHandler(wr http.ResponseWriter, req *http.Request) {
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}

	defer req.Body.Close()

	err = json.Unmarshal(body, &sr)
	if err != nil {
		log.Panic(err)
	}

	result, err := json.Marshal(search(sr.Sites, sr.Search))
	if err != nil {
		log.Panic(err)
	}

	wr.Write(result)
}

func search(arr []string, query string) (result []string) {
	for i := range arr {
		resp, err := http.Get(arr[i])
		if err != nil {
			fmt.Println(err)
			return
		}
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println(err)
			return
		}
		if strings.Contains(string(body), query) {
			result = append(result, arr[i])
		}
	}
	return
}

// task2
func setCookie(wr http.ResponseWriter, req *http.Request) {
	cookie := http.Cookie{Name: "hw2_task2", Value: "hw2_Done!"}
	http.SetCookie(wr, &cookie)
	wr.Write([]byte("cookie set successfully"))
}

func getCookie(wr http.ResponseWriter, req *http.Request) {
	c, err := req.Cookie("hw2_task2")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Fprintf(wr, "Наша куки: %s", c.Value)
}

func main() {
	router := http.NewServeMux()

	// task1
	router.HandleFunc("/post", postHandler)

	//task2
	router.HandleFunc("/set_cookie", setCookie)
	router.HandleFunc("/get_cookie", getCookie)

	log.Fatal(http.ListenAndServe(":8080", router))
}

// curl -X POST -H 'Content-Type: application/json' -d "{\"search\":\"programming\",\"sites\":[\"https://ya.ru\",\"https://www.google.ru\",\"https://mail.ru\",\"https://www.rambler.ru\",\"https://golang.org\"]}" http://localhost:8080/post
// curl -X POST -H 'Content-Type: application/json' -d "{\"search\":\"programming\",\"sites\":[\"https://ya.ru\",\"https://www.google.ru\",\"https://mail.ru\",\"https://www.rambler.ru\",\"https://golang.org\"]}" https://postman-echo.com/post
