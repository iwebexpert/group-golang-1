package main

import (
	"flag"
	"fmt"
	"golang.org/x/sync/errgroup" // библиотека для упрощенной работы с горутинами и возможностью возврата ошибки из горутины
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"sync"
)


func SearchLinksX(queryString string, urls []string) ([]string, error) {
	group := struct {
		errgroup.Group //запуск горутин с использованием ошибок
		sync.Mutex //синхронизация горутин
		urls []string
	}{
		urls: make([]string, 0, len(urls)),
	}

	for _, url := range urls {
		group.Go(func() error { // запускает функцию внутри горутины
			res, err := http.Get(url)
			if err != nil {
				return err
			}
			defer res.Body.Close()

			body, err := ioutil.ReadAll(res.Body)
			if err != nil{
				return err
			}

			if strings.Contains(string(body), queryString) {
				group.Lock() //блокировка на случай, если одновременно будет проходить запись в group.url
				group.urls = append(group.urls, url)
				group.Unlock()
			}
			return nil
		})
	}
	err := group.Wait()
	return group.urls, err
}

func main() {
	defaultLinks := []string{
		"https://google.com/",
		"https://ya.ru/",
		"https://yandex.ru/",
	}
	querySearch := flag.String("query", "ya", "query for search")
	urls := flag.String("urls", strings.Join(defaultLinks, ","), "list of sites")
	flag.Parse()

	resultSearch, err := SearchLinksX(*querySearch, strings.Split(*urls, ","))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(strings.Join(resultSearch, "\n"))
}