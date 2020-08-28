package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"sync"
)

func SearchLinks(queryString string, urls []string) ([]string, error) {
	answer := struct {
		sync.Mutex
		sync.WaitGroup
		urls []string
	}{
		urls: make([]string, 0, len(urls)),
	}
	for _, url := range urls {
			answer.Add(1)
			go func(url string) {
				res, err := http.Get(url)
				if err != nil {
					log.Fatal(err) // тут нельзя прокинуть ошибку выше в main, или нужен просто return, что неинформативно
				}

				body, err := ioutil.ReadAll(res.Body)
				if err != nil {
					log.Fatal(err)
				}

				if strings.Contains(string(body), queryString) {
					answer.Lock()
					answer.urls = append(answer.urls, url)
					answer.Unlock()
				}
				answer.Done()
			}(url)
	}
	answer.Wait()
	return answer.urls, nil
}

func main() {
	defaultLinks := []string{
		"https://google.com/",
		"https://ya.ru/",
		"https://yandex.ru/",
	}

	querySearch := flag.String("query", "ya", "query for search")
	urls := flag.String("urls", strings.Join(defaultLinks, ","), "list for search")
	flag.Parse()


	resultSearch, err := SearchLinks(*querySearch, strings.Split(*urls, ","))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(strings.Join(resultSearch, "\n"))
}