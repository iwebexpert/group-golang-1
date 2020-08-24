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

func SearchLinks(queryString string, urls []string, wg *sync.WaitGroup) ([]string, error) {
	answer := struct {
		sync.Mutex
		urls []string
	}{
		urls: make([]string, 0, len(urls)),
	}
	for _, v := range urls {
		url := v

			res, err := http.Get(url)
			if err != nil {
				log.Fatal(err) // тут нельзя прокинуть ошибку выше в main, или нужен просто return, что неинформативно
			}

			body, err := ioutil.ReadAll(res.Body)
			if err != nil {
				log.Fatal(err)
			}


			if strings.Contains(string(body),queryString) {
				answer.Lock()
				answer.urls = append(answer.urls, url)
				answer.Unlock()
			}

	}
	defer wg.Done()
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

	var wg sync.WaitGroup
	wg.Add(1)
	resultSearch, err := SearchLinks(*querySearch, strings.Split(*urls, ","), &wg)
	if err != nil {
		log.Fatal(err)
	}
	wg.Wait()
	fmt.Println(strings.Join(resultSearch, "\n"))
}