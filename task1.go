package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	query := flag.String("q", "morda", "query to find")
	pages := flag.String("p", "https://ya.ru,https://mail.ru", "urls")
	flag.Parse()
	pagesParsed := strings.Split(*pages, ",")
	fmt.Println("looking for: ", *query)
	fmt.Println("looking at: ", pagesParsed)
	whoContains := []string{}
	for _, p := range pagesParsed {

		response, _ := http.Get(p)
		defer response.Body.Close()

		bytesArray, _ := ioutil.ReadAll(response.Body)

		if strings.Contains(string(bytesArray), *query) {
			whoContains = append(whoContains, p)
		}
	}
	fmt.Println("Found", *query, "at:", whoContains)
}

//For homework
//ya.ru
//google.com
//...

// query
//go run websocket.go --query "Red" --pages "http://ya.ru,http://1.ru"
// "flag", поиск через "strings"

//2 переменных - query, pages
//GET pages
//Ищем query
//Возвращаем и отображаем в консоль список страниц, которые содержат query
