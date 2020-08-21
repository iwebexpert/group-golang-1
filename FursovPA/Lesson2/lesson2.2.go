package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	"github.com/go-chi/chi"
	uuid "github.com/satori/go.uuid"
)

type UserInfo struct {
	Name string `json:"name"`
	Age int64   `json:"age"`
}

var (
	sessionDB = map[string]*UserInfo{}
	mutex	  = &sync.Mutex{}
)

func getUserInfoHandler(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session")
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	UserInfo, ok := sessionDB[cookie.Value]
	if  !ok {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	output, err := json.Marshal(UserInfo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("countent-type", "application/json")
	w.Write(output)
}

func setUserInfoHeandler(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	userInfo := new(UserInfo)

	if err := json.Unmarshal(b, userInfo); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	id, err := uuid.NewV4()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	cookie := http.Cookie{
		Name: "session",
		Value: id.String(),
	}

	http.SetCookie(w, &cookie)
}

func main() {
	r := chi.NewRouter()

	r.Get("/getInfo", getUserInfoHandler)
	r.Post("/setInfo", setUserInfoHeandler)

	log.Println("Server start")

	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatal(err)
	}
}