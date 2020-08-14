package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	query := flag.String("query", "find", "query to find")
	pages := flag.String("urls", "https://ya.ru,https://google.com", "urls")
	flag.Parse()
	pagesParsed := strings.Split(*pages, ",")
	fmt.Println("looking for: ", *query)
	fmt.Println("looking at: ", pagesParsed)
	by1page := []string{}
	for _, p := range pagesParsed {
		response, _ := http.Get(p)
		defer response.Body.Close()

		bytesArray, _ := ioutil.ReadAll(response.Body)

		if strings.Contains(string(bytesArray), *query) {
			by1page = append(by1page, p)
		}
	}
	fmt.Println("Found", *query, "at:", by1page)
}
