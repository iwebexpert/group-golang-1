package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
)

type searchJSON struct {
	Query string   `json:"search"`
	Urls  []string `json:"sites"`
}

type pageProc struct {
	url  string
	page []byte
}

func indexHandler(w http.ResponseWriter, r *http.Request) {

	//request check
	if r.Method == "POST" && r.Header.Get("Content-Type") == "application/json" {
		var bodyJSON searchJSON

		respBody, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Fatal(err)
		}
		err = json.Unmarshal(respBody, &bodyJSON)
		if err != nil {
			log.Fatal(err)
		}
		query := bodyJSON.Query
		urls := bodyJSON.Urls
		n := len(urls)

		pagesChan := make(chan pageProc, n)
		resultsChan := make(chan []string, n)

		//GET pages
		go func() {
			for _, p := range urls {
				response, err := http.Get(p)
				if err != nil {
					log.Println(err)
					log.Println("Looks like invalid URL. Skipping")
					n--
					continue
				}
				defer response.Body.Close()
				bytesArray, err := ioutil.ReadAll(response.Body)
				if err != nil {
					log.Fatal(err)
				}
				page := pageProc{url: p, page: bytesArray}
				pagesChan <- page
			}
			fmt.Println("GET pages done")
		}()

		//Look for query
		go func() {
			var result []string
			var pageResult pageProc
			for i := 0; i < n; i++ {
				pageResult = <-pagesChan
				if strings.Contains(string(pageResult.page), query) {
					fmt.Println("found at", pageResult.url)
					result = append(result, pageResult.url)
				}

			}
			resultsChan <- result
			fmt.Println("Lookup finished")
		}()
		//work with result
		var resultsJSON searchJSON
		resultsJSON.Query = query
		resultsJSON.Urls = <-resultsChan

		answerJSON, err := json.Marshal(resultsJSON)
		if err != nil {
			log.Fatal(err)
		}

		w.Header().Set("content-type", "application/json")
		w.Write(answerJSON)

	}
}

func main() {
	//postedChan := make(chan []byte)
	stopChan := make(chan os.Signal)

	router := http.NewServeMux()

	router.HandleFunc("/", indexHandler)

	go func() {
		log.Println("Server start")
		err := http.ListenAndServe(":8080", router)
		log.Fatal(err)
	}()

	//Wait signal
	signal.Notify(stopChan, os.Kill, os.Interrupt)
	<-stopChan

	log.Println("Server stop")

}
