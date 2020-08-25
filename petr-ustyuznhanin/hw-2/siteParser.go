package main

import (
	"encoding/json"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"sync"
)

type QueryT struct {
	Search string `json:"search"`
	Sites []string `json:"sites"`
}

type AnswerT struct {
	Sites []string `json:"sites"`
}

var (
	data QueryT
)

func handleGetQuery(w http.ResponseWriter, r *http.Request){
	//читаем тело запроса с Json и парсим его
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.Unmarshal(b, &data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var (
		wg sync.WaitGroup
		mu sync.Mutex
		)

	var answer AnswerT
	for _, v := range data.Sites {
		wg.Add(1)
		go func() {
			res, err := http.Get(v)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			defer res.Body.Close()

			body, err := ioutil.ReadAll(res.Body)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			if strings.Contains(string(body), data.Search) {
				mu.Lock()
				answer.Sites = append(answer.Sites, v)
				mu.Unlock()
			}
			wg.Done()
		}()
		wg.Wait()
	}

	output, err := json.Marshal(answer)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(output)
}

func main() {
	stopChan := make(chan os.Signal)


	router := chi.NewRouter()

	router.Use(middleware.Logger)

	router.Post("/", handleGetQuery)

	go func() {
		log.Println("Server starts")
		err := http.ListenAndServe(":8080", router)
		log.Fatal(err)
	}()

	//ожидание сигнала на выключение сервера
	signal.Notify(stopChan, os.Kill, os.Interrupt)
	<- stopChan

	log.Println("Server stop")
}
