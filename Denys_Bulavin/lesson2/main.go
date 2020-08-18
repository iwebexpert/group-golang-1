package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
)

type Search struct {
	Search string   `json:"search"`
	Sites  []string `json:"sites"`
}
type Cookie struct {
	Name    string        `json:"name"`
	Value   string        `json:"value"`
	Expires time.Duration `json:"expire"`
}

var data Search

func webParser(s Search) []string {
	var out []string
	waitGroup := sync.WaitGroup{}
	for _, r := range s.Sites {

		waitGroup.Add(1)
		go func() {
			resp, err := http.Get(r)
			if err != nil {
				log.Fatal(err)
			}
			defer resp.Body.Close()

			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				log.Fatal(err)
			}

			if strings.Contains(string(body), s.Search) {
				out = append(out, r)

			}
			waitGroup.Done()
		}()
		waitGroup.Wait()
	}
	return out

}

func setCookie(w http.ResponseWriter, req *http.Request) {
	c := Cookie{"Name", "Denys", 30 * time.Minute}
	addCookie(w, c)

	io.WriteString(w, "add Cookie")
}

func addCookie(w http.ResponseWriter, c Cookie) {
	expire := time.Now().Add(c.Expires)
	cookie := http.Cookie{
		Name:    c.Name,
		Value:   c.Value,
		Expires: expire,
	}
	http.SetCookie(w, &cookie)
}

func reqTime(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "The current time is: %s\n", time.Now())
}

func getCookie(w http.ResponseWriter, r *http.Request) {
	c, _ := r.Cookie("Name")
	value := c.Value
	w.Write([]byte("Ваше Имя : " + value + "\n"))
}

func getParams(w http.ResponseWriter, r *http.Request) {

	body, _ := ioutil.ReadAll(r.Body)

	json.Unmarshal(body, &data)
	fmt.Println(data)

}

func searchResult(w http.ResponseWriter, r *http.Request) {

	render.JSON(w, r, webParser(data))
}

func main() {

	route := chi.NewRouter()
	route.Use(middleware.Logger)

	route.Get("/time", reqTime)

	route.Route("/search", func(r chi.Router) {
		r.Post("/get", getParams)
		r.Get("/result", searchResult)
	})

	route.Route("/cookie", func(r chi.Router) {
		r.Get("/set", setCookie)
		r.Get("/get", getCookie)

	})

	err := http.ListenAndServe(":8080", route)
	if err != nil {
		fmt.Println("ListenAndServe:", err)
	}

}
