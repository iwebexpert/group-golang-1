package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"sync"

	uuid "github.com/satori/go.uuid"
)

//Sessions -
var Sessions map[string]string

const COOKIE_SEEKER_KEY = "SEEKER"

func main() {
	stopChannel := make(chan os.Signal)
	signal.Notify(stopChannel, os.Kill, os.Interrupt)

	Sessions = make(map[string]string)

	http.HandleFunc("/", getRootHandler)
	http.HandleFunc("/seek", getSeekHandler)

	go func() {
		fmt.Println("Server start")
		err := http.ListenAndServe(":3333", nil)
		fmt.Println(err)
	}()

	//Ждем сигнала от OS
	<-stopChannel
	fmt.Println("Server stop")
}

func getRootHandler(w http.ResponseWriter, r *http.Request) {
	sid := cookieCheck(w, r)
	if sid == "" {
		query := r.URL.Query()
		name, ok := query["name"]
		if ok {
			sid = cookieSet(w, name[0])
		}
	}
	w.Write([]byte("Hello " + getSeekerName(sid)))
}

func getSeekerName(sid string) string {
	if sid == "" {
		return "ANONIMUS"
	}
	return Sessions[sid]
}

func getSeekHandler(w http.ResponseWriter, r *http.Request) {
	sid := cookieCheck(w, r)
	query := r.URL.Query()
	phrase, ok := query["phrase"]
	if !ok || len(phrase) != 1 {
		w.Write([]byte("Phrase is mandatory. Only one phrase per request"))
	}
	links, ok := query["links"]
	if !ok || len(links) == 0 {
		w.Write([]byte("At least one link must be declared"))
	}
	fmt.Println("Пользователь", getSeekerName(sid))
	fmt.Println("Поиск", phrase[0])
	fmt.Println("На ресурсах", links)
	s := seek(phrase[0], links)
	if len(s) > 0 {
		w.Write([]byte(strings.Join(s, ",")))
	} else {
		w.Write([]byte("Not found"))
	}
}

func cookieCheck(w http.ResponseWriter, r *http.Request) string {
	cookie, _ := r.Cookie(COOKIE_SEEKER_KEY)
	if cookie == nil {
		return ""
	}
	_, ok := Sessions[cookie.Value]
	if !ok {
		cookieSet(w, "")
		return ""
	}
	return cookie.Value
}

func cookieSet(w http.ResponseWriter, seekername string) string {
	var newkey string
	if seekername != "" {
		newkey = uuid.Must(uuid.NewV4()).String()
		Sessions[newkey] = seekername
	} else {
		newkey = ""
	}
	http.SetCookie(w, &http.Cookie{
		Name:  COOKIE_SEEKER_KEY,
		Value: newkey,
	})
	return newkey
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
