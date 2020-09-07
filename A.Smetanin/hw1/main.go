package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

var searchQueryFlag = flag.String("query", "help", "Enter query to search")
var sitesListFlag = flag.String("sites", "https://ya.ru, https://google.com", "Enter URLs to search from (http://example.com) separated by comma")
var result []string

func main() {
	flag.Parse()
	sites := strings.Split(*sitesListFlag, ", ")
	for i := 0; i < len(sites); i += 1 {
		resp, _ := http.Get(sites[i])
		body, _ := ioutil.ReadAll(resp.Body)
		if strings.Contains(string(body), *searchQueryFlag) {
			result = append(result, sites[i])
		}
	}
	fmt.Println(result)
}
