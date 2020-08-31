package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"sync"
)

// Create pagesArr and implement method String and Set
// because package flag doesn`t support arrayFlags
type pagesArr []string

func (i *pagesArr) String() string {
	return fmt.Sprintf("%s", *i)
}

func (i *pagesArr) Set(value string) error {
	*i = append(*i, value)
	return nil
}

// define flags
var (
	pages pagesArr
	query = flag.String("query", "apple", "this is query for parsing sites")
)

// Parse flags
func init() {
	flag.Var(&pages, "pages", "this pages will be parsed")
	flag.Parse()

	if len(pages) == 0 {
		pages = pagesArr{"https://ya.ru"}
	}
}

// get pageUrl and return html body
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

// Concurrent parsing sites
// If error appears return only error
// Without errors return sites that have subQuery
func sitesParser(query string, pages []string) ([]string, error) {
	resultPages := []string{}

	wg := sync.WaitGroup{}
	errCh := make(chan error)

	for _, page := range pages {
		wg.Add(1)

		go func(page, query string) {
			pageContent, err := fetch(page)
			if err != nil {
				wg.Done()
				errCh <- err
				return
			}

			if strings.Contains(pageContent, query) {
				resultPages = append(resultPages, page)
			}

			wg.Done()
		}(page, query)
	}

	wg.Wait()

	select {
	case err := <-errCh:
		return nil, err
	default:
		return resultPages, nil
	}
}

func main() {
	discoveredPages, err := sitesParser(*query, pages)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(discoveredPages)
}
