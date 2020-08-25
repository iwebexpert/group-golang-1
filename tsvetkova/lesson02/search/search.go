package search

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"sync"
)

type SearchResult struct {
	Query string `json:"search"`
	Urls []string `json:"sites"`
	m 	 sync.Mutex `json:"-"`
}

func (r *SearchResult ) toJson() []byte {
	if len(r.Urls) == 0 {
		r.Urls = []string{}
	}

	b, err := json.Marshal(*r)
	if err != nil {
		log.Fatal(err)
	}
	
	return b
}

var PAGES = []string{
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

func SearchSites(query string) []byte {
	var wg sync.WaitGroup
	
	var result = SearchResult{
		Query: query,
	}

	for _, url := range PAGES {
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
				result.Urls = append(result.Urls, url)
				result.m.Unlock()
			}
		}(url)
	}
	wg.Wait()
	
	return result.toJson()
}