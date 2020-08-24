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
		//go func() {
			res, err := http.Get(v)
			if err != nil {
				log.Fatal(err) // тут нельзя прокинуть ошибку выше в main, или нужен просто retrun, что неинформативно
			}

			body, err := ioutil.ReadAll(res.Body)
			if err != nil {
				log.Fatal(err)
			}


			if strings.Contains(string(body),queryString) {
				answer.Lock()
				answer.urls = append(answer.urls, v)
				answer.Unlock()
			}
		//}()
	}
	wg.Done()
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

	fmt.Println(strings.Join(resultSearch, "\n"))
}