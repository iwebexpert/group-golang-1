package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func webParser(s string, m []string) []string {
	var out []string
	for _, r := range m {
		resp, err := http.Get(r)
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}

		if strings.Contains(string(body), s) {
			out = append(out, r+" - cодержит заданную строку.")
		}

	}
	return out

}

func main() {
	searchString := flag.String("s", "Москва", "Строка которую будем искать")
	hosts := flag.String("h", "https://mail.ru,https://rambler.ru", "Список сайтов через зяпятую без пробелов")
	flag.Parse()
	fmt.Println(webParser(*searchString, strings.Split(*hosts, ",")))
}
