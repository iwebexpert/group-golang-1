package service

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

func SearchV1(search_str string, urls []string) []string {
	result := []string{}
	for i := 0; i < len(urls); i++ {
		u, err_valid_url := url.ParseRequestURI(urls[i])
		if err_valid_url != nil {
			log.Println("Error on parsing url")
			continue
		}

		log.Println(fmt.Sprintf("Processing - %s", u.String()))

		resp, err_get := http.Get(u.String())
		if err_get != nil {
			log.Println("Error: Get request failed")
			continue
		}

		defer resp.Body.Close()
		body, err_body := ioutil.ReadAll(resp.Body)
		if err_body != nil {
			log.Println("ERROR on reading from buffer")
			continue
		}
		if strings.Contains(string(body), search_str) {
			result = append(result, u.String())
		} else {
			fmt.Println("Nope. Not found")
		}
	}
	return result
}
