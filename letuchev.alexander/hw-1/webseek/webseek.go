/*1. Напишите функцию, которая будет получать на вход строку с поисковым запросом (string) и массив ссылок на страницы,
по которым стоит произвести поиск ([]string). Результатом работы функции должен быть массив строк со ссылками на страницы,
 на которых обнаружен поисковый запрос. Функция должна искать точное соответствие фразе в тексте ответа от сервера по каждой
 из ссылок.*/

package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"sync"
)

func main() {
	linkarr := os.Args[2:]
	/*linkarr := []string{
		"https://ru.wikiedia.org",
		"https://ru.wikipedia.org/wiki/%D0%90%D0%B1%D1%80%D0%B0%D0%BA%D0%B0%D0%B4%D0%B0%D0%B1%D1%80%D0%B0",
		"https://ru.wikipedia.org/wiki",
		"https://geekbrains.ru",
		"https://rbc.ru",
		"http://lib.ru",
		"http://libqerqwerqwerqwer.ru",
	}*/
	strFind := os.Args[1]
	//strFind := "Абракадабра"
	strArr := seek(strFind, linkarr)
	fmt.Printf("Найдено %d вхождений строки [%s]:\r\n", len(strArr), strFind)
	fmt.Println(strArr)
}

func seek(findstr string, pagearr []string) []string {
	stopChannel := make(chan int)
	resultChannel := make(chan string)
	resultArray := []string{}

	//горутина-сборщик результата
	go func(resultChannel *chan string, resultArray *[]string, stopChannel *chan int) {
		for {
			select {
			case s := <-*resultChannel:
				*resultArray = append(*resultArray, s)
			case <-*stopChannel:
				return
			}
		}
	}(&resultChannel, &resultArray, &stopChannel)

	//запуск горутин поиска строки
	var wg sync.WaitGroup
	for i := 0; i < len(pagearr); i++ {
		wg.Add(1)
		go func(wg *sync.WaitGroup, stream int, addr string, findstr string, resultChannel *chan string) {
			var cnt int64
			defer wg.Done()
			fmt.Println("Seek stream:", stream, "opening", addr)
			resp, err := http.Get(addr)
			if err != nil {
				fmt.Println("Seek stream:", stream, err)
				return
			}
			fmt.Println("Seek stream:", stream, "response code", resp.StatusCode)
			defer resp.Body.Close()
			buff := bytes.NewBuffer([]byte(""))
			cnt, err = io.Copy(buff, resp.Body)
			if err != nil {
				fmt.Println("Seek stream:", stream, err)
				return
			}
			fmt.Println("Seek stream:", stream, "bytes read", cnt)
			if strings.Contains(buff.String(), findstr) {
				fmt.Println("Seek stream:", stream, "FOUND")
				*resultChannel <- addr
			} else {
				fmt.Println("Seek stream:", stream, "NOT FOUND")
			}
		}(&wg, i, pagearr[i], findstr, &resultChannel)
	}
	//Ожидание окончания работы горутин поиска
	wg.Wait()
	//Остановка горутины-сборщика
	stopChannel <- 0
	return resultArray
}
