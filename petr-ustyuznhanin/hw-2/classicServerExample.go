package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	file, err := ioutil.ReadFile("./pages/index.html")
	if err != nil {
		log.Println(err)
	}
	w.Write(file)
}

func main() {
	stopChan := make(chan os.Signal)

	router := http.NewServeMux()

	router.HandleFunc("/", IndexHandler)

	go func() {
		log.Println("Server starts")
		err := http.ListenAndServe(":8080", router)
		log.Fatal(err)
	}()

	//ожидание сигнала на выключение сервера
	signal.Notify(stopChan, os.Kill, os.Interrupt)
	<- stopChan

	log.Println("Server stop")
}