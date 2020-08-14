package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

// HowTo: run 'go build findOnSite.go'
// run './findOnSite -query="somestring" -pages="space-separated urls for search"'

//findOnSite returns a list of sites where found s string
func findOnSite (s string, searchList []string) (res []string) {
	for _, v := range searchList{
		response, err := http.Get(v)
		if err != nil {
			panic(err) // how to handle the error correctly?
		}
		bytesArray, err := ioutil.ReadAll(response.Body)
		if err != nil {
			panic(err)
		}

		if strings.Contains(string(bytesArray), s) {
			res = append(res, v)
		}
	}
	return res
}

type urls []string

func(u *urls) String() string {
	return fmt.Sprint(*u)
}

func(u *urls) Set(value string) error {
	for _, pages := range strings.Split(value, " ") {
		*u = append(*u, pages)
	}
	return nil
}


func main() {
	query := flag.String("query", "empte", "query for search")
	var pages urls
	flag.Var(&pages, "pages", "pages for query search")
	flag.Parse()

	fmt.Println(findOnSite(*query, pages))
}