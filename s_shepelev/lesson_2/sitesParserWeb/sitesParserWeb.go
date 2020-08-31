package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"sync"

	"github.com/go-chi/chi"
)

type InputMessage struct {
	Search string   `json:"search"`
	Sites  []string `json:"sites"`
}

type OutputMessage struct {
	Sites []string `json:"sites"`
}

// Get pageUrl and return html body
func fetch(page string) (string, error) {
	resp, err := http.Get(page)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

// Concurrently check every site for seach string match
// Return JSON with array of sites with search string
func sitesParserHandler(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	msg := new(InputMessage)

	if err := json.Unmarshal(b, msg); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resultPages := new(OutputMessage)

	wg := sync.WaitGroup{}
	mu := sync.Mutex{}
	errCh := make(chan error)

	for _, page := range msg.Sites {
		wg.Add(1)

		go func(page, query string) {
			pageContent, err := fetch(page)
			if err != nil {
				wg.Done()
				errCh <- err
				return
			}

			if strings.Contains(pageContent, query) {
				mu.Lock()
				resultPages.Sites = append(resultPages.Sites, page)
				mu.Unlock()
			}

			wg.Done()
		}(page, msg.Search)
	}

	wg.Wait()

	select {
	case err := <-errCh:
		http.Error(w, err.Error(), http.StatusInternalServerError)
	default:
		output, err := json.Marshal(resultPages)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("content-type", "application/json")
		w.Write(output)
	}
}

func main() {
	r := chi.NewRouter()

	r.Post("/", sitesParserHandler)

	log.Println("Server start")

	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatal(err)
	}
}
