package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func Search(query string, pages []string) ([]string, error) {
	var result []string
	for _, url := range pages {
		response, err := http.Get(url)
		if err != nil {
			return result, err
		}
		defer response.Body.Close()

		b, err := ioutil.ReadAll(response.Body)
		if err != nil {
			return result, err
		}

		if strings.Contains(string(b), query) {
			result = append(result, url)
		}
	}

	return result, nil
}

func parseArgs() (string, []string) {
	query := flag.String("query", "go version", "search query")
	pagesString := flag.String("pages", "https://golang.org/pkg/math/,https://golang.org/doc/install,https://stackoverflow.com/questions/42952979/go-version-command-shows-old-version-number-after-update-to-1-8", "pages to search")
	flag.Parse()

	pages := strings.Split(*pagesString, ",")
	return *query, pages
}

func main() {
	query, pages := parseArgs()
	fmt.Printf("Searching for: %q on %d pages\n", query, len(pages))
	result, err := Search(query, pages)
	if err != nil {
		log.Fatalf("Error occured: %s", err)
	}
	if len(result) == 0 {
		fmt.Println("No matches found\n")
	} else {
		fmt.Printf("Found match on %d pages\n", len(result))
		for idx, url := range result {
			fmt.Printf("%d. %s\n", idx+1, url)
		}
	}
}
