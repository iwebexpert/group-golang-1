package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func search(arr []string, q string) (result []string) {
	for i := range arr {
		resp, err := http.Get(arr[i])
		if err != nil {
			fmt.Println(err)
			return
		}
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println(err)
			return
		}
		if strings.Contains(string(body), q) {
			result = append(result, arr[i])
		}
	}
	return
}

func main() {
	var query = flag.String("query", "", "Your query")
	var pages = flag.String("pages", "", "Pages for search")
	var links []string
	//"https://ya.ru,https://www.google.ru,https://mail.ru,https://www.rambler.ru,https://golang.org"
	flag.Parse()

	for _, link := range strings.Split(*pages, ",") {
		links = append(links, link)
	}
	fmt.Println(search(links, *query))
}
