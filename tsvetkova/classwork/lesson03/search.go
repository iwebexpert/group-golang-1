  package main

import (
	"flag"
	"strings"
	"sync"
	"net/http"
	"io/ioutil"
	"log"
	"fmt"

	"golang.org/x/sync/errgroup"
)

func main() {
	defaultLinks := []string{
		"http://google.com",
		"http://ya.ru",
	}

	urls := flag.String("urls", strings.Join(defaultLinks, ","), "list of sites")
	query := flag.String("query", "ya", "search query")
	flag.Parse()

	result, err := SearchLinks(*query, strings.Split(*urls, ","))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(result)
}

func SearchLinks(query string, urls []string) ([]string, error) {
	group := struct {
		errgroup.Group // запуск горутины с отловом ошибок
		sync.Mutex	// синхронизация горутин
		urls []string
	}{
		urls: make([]string, 0, len(urls)),
	}

	for _, url := range urls {

		group.Go(func() error {
			resp, err := http.Get(url)
			if err != nil {
				return err
			}
			defer resp.Body.Close()

			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return err
			}

			// поиск с блокировкой
			if strings.Contains(string(body), query) {
				group.Lock()
				group.urls = append(group.urls, url)
				group.Unlock()
			}

			return nil
		})

	}
	err := group.Wait()
	return group.urls, err
}