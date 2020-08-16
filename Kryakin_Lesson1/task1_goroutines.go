package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type pageProc struct {
	url  string
	page []byte
}

func main() {
	//parse flags
	query := flag.String("q", "morda", "query to find")
	pages := flag.String("p", "https://ya.ru,https://mail.ru", "urls")
	flag.Parse()
	pagesParsed := strings.Split(*pages, ",")
	fmt.Println("looking for: ", *query)
	fmt.Println("looking at: ", pagesParsed)
	n := len(pagesParsed)

	//make channels
	pagesChan := make(chan pageProc, len(pagesParsed))
	resultsChan := make(chan []string, len(pagesParsed))

	//get pages
	go func() {
		for _, p := range pagesParsed {
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

	//look for querry
	go func() {
		var result []string
		var pageResult pageProc
		for i := 0; i < n; i++ {
			pageResult = <-pagesChan
			if strings.Contains(string(pageResult.page), *query) {
				fmt.Println("found at", pageResult.url)
				result = append(result, pageResult.url)
			}

		}
		resultsChan <- result
		fmt.Println("Lookup finished")
	}()
	//print result from channel
	fmt.Println(<-resultsChan)
}
