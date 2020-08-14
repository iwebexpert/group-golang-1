package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)




func main() {
	urls := flag.String("urls", "https://ya.ru,https://google.com", "urls")
	query := flag.String("query", "a", "query")
	flag.Parse()
	pagesParsed := strings.Split(*urls, ",")
	fmt.Println("Что ищем: ", *query)
	fmt.Println("Где ищем: ", pagesParsed)

	Cont := []string{}

	for _, urls := range pagesParsed {
		resp, _ := http.Get(urls)

		defer resp.Body.Close()
		body, _ := ioutil.ReadAll(resp.Body)

		if strings.Contains(string(body), *query) {

			Cont = append(Cont, urls)

		}
	}
fmt.Println("Найдено:", *query, "В:", Cont)
}





