package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Message struct {
	Search string   `json:"Search"`
	Url    []string `json:"url"`
}

func ParsJson(w http.ResponseWriter, r *http.Request) {

	b, _ := ioutil.ReadAll(r.Body) // доделать проверку на ошибку
	defer r.Body.Close()

	x := Message{}
	_ = json.Unmarshal(b, &x) // доделать проверку на ошибку

	fmt.Fprintf(w, "%+v", x)

	output, err := json.Marshal(x)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.Header().Set("content-type", "application/json")
	w.Write(output)

}

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/json", ParsJson)
	http.ListenAndServe(":8080", mux)
	log.Println("Starting server on address")
}
