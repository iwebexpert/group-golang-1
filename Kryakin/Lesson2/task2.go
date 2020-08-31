package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	uuid "github.com/satori/go.uuid"
)

var db = map[string]*user{}

type user struct {
	Name string `json:"name"`
	ID   int    `json:"ID"`
}

func getHandler(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("uuid")
	if err != nil {
		log.Fatal(err)
	}

	user, ok := db[cookie.Value]
	if !ok {
		log.Fatal("user := db != ok")
	}

	result, err := json.Marshal(user)
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("content-type", "application/json")
	w.Write(result)
}

func postHandler(w http.ResponseWriter, r *http.Request) {
	byteArray, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	user := new(user)

	if err := json.Unmarshal(byteArray, user); err != nil {

	}

	id, err := uuid.NewV4()
	if err != nil {
		log.Fatal(err)
	}

	db[id.String()] = user

	cookie := http.Cookie{
		Name:  "uuid",
		Value: id.String(),
	}

	http.SetCookie(w, &cookie)
}
func main() {
	r := chi.NewRouter()

	r.Get("/get", getHandler)
	r.Post("/post", postHandler)

	log.Println("Server start")

	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatal(err)
	}
}
