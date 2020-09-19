package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"sync"
)

var searchQueryFlag = flag.String("query", "help", "Enter query to search")
var sitesListFlag = flag.String("sites", "https://ya.ru, https://google.com", "Enter URLs to search from (http://example.com) separated by comma")
var result []string
var wg sync.WaitGroup


func f(url string, query string, result *[]string )  {
	resp, _ := http.Get(url)
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	if strings.Contains(string(body), query) {
		*result = append(*result, url)
	}
	wg.Done()
}

func main() {
	flag.Parse()
	sites := strings.Split(*sitesListFlag, ", ")
	wg.Add(len(sites))
	for i := 0; i < len(sites); i += 1 {
		go f(sites[i], *searchQueryFlag, &result)
	}
	wg.Wait()
	fmt.Println(result)
}
