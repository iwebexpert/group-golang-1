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

type SearchResult struct {
	urls []string
	m 	 sync.Mutex
}

func getPage(url string) (string, error) {
	response, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	b, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	return string(b), nil
}

func Search(query string, pages []string) []string {
	var wg sync.WaitGroup
	var result SearchResult

	for _, url := range pages {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()

			text, err := getPage(url)
			if err != nil {
				log.Println(err)
				return
			}

			if strings.Contains(text, query) {
				result.m.Lock()
				result.urls = append(result.urls, url)
				result.m.Unlock()
			}
		}(url)
	}

	wg.Wait()
	return result.urls
}

func parseArgs() (string, []string) {
	query := flag.String("query", "файл", "search query")
	pagesString := flag.String("pages", "", "pages to search")
	flag.Parse()

	var pages []string
	if len(*pagesString) == 0 {
		pages = []string{
			"https://geekbrains.ru/posts/perspektivy-java-ai-i-tvorcheskie-podrostki",					// false
			"https://geekbrains.ru/posts/kak-podklyuchat-fajly-php-i-zachem-ehto-voobshche-nuzhno",		// true
			"https://geekbrains.ru/posts/data_recovery",												// true
			"https://geekbrains.ru/posts/kurs-dlya-detej-po-razrabotke-igr-na-python",					// false
			"https://geekbrains.ru/posts/tts_python",													// true
			"https://geekbrains.ru/posts/postroenie-nebolshogo-centra-sertifikacii-na-osnove-openssl",	// true
			"https://geekbrains.ru/posts/programming_types",											// false
			"https://geekbrains.ru/posts/npp_java",														// true
			"https://geekbrains.ru/posts/100-goryachih-klavish-photoshop-dlya-ognennogo-rezultata",		// true
			"https://geekbrains.ru/posts/brosat-razvitie-sobstvennyh-produktov-ya-ne-sobiralsya",		// false
			"https://geekbrains.ru/posts/chto-takoe-css-obyasnyaem-prostymi-slovami",					// true
			"https://geekbrains.ru/posts/kak-rabotat-iz-doma-bolee-ehffektivno",						// false
			"https://geekbrains.ru/posts/kak-zashifrovat-pochtu-i-sdelat-ehlektronnuyu-podpis",			// true
			"https://geekbrains.ru/posts/kak-za-pyat-shagov-nastroit-skvoznuyu-analitiku",				// false
			"https://geekbrains.ru/posts/9-horoshih-servisov-proverki-koda-dlya-razrabotchikov",		// false
		}
	} else {
		pages = strings.Split(*pagesString, ",")
	}

	return *query, pages
}

func main() {
	query, pages := parseArgs()
	fmt.Printf("Searching for %q on %d pages\n", query, len(pages))

	result := Search(query, pages)

	if len(result) == 0 {
		fmt.Println("No matches found\n")
	} else {
		fmt.Printf("Found match on %d page(s)\n", len(result))
		for idx, url := range result {
			fmt.Printf("%d. %s\n", idx+1, url)
		}
	}
}
