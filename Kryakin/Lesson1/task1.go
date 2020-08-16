package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
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

		response, err := http.Get(p)
		if err != nil {
			log.Fatal(err)
		}
		defer response.Body.Close()

		bytesArray, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Fatal(err)
		}

		if strings.Contains(string(bytesArray), *query) {
			whoContains = append(whoContains, p)
		}
	}
	fmt.Println("Found", *query, "at:", whoContains)
}
