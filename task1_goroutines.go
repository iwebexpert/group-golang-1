package main

import (
	"flag"
	"fmt"
	"io/ioutil"
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
			response, _ := http.Get(p)
			defer response.Body.Close()
			bytesArray, _ := ioutil.ReadAll(response.Body)
			page := pageProc{url: p, page: bytesArray}
			pagesChan <- page
		}
		fmt.Println("pages done")
	}()
	//look for querry
	fmt.Println(n)
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
		fmt.Println("lookup finished")
	}()
	//print result from channel
	fmt.Println(<-resultsChan)
}
