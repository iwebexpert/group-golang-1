package main

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

const yaCloudURL = "https://cloud-api.yandex.net/v1/disk/public/resources?public_key="

type dataJSON struct {
	Download string `json:"file"`
}

func getURLDownload(uri string) string {

	var link dataJSON

	resp, err := http.Get(yaCloudURL + uri)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(body, &link)

	return link.Download
}

func fileDownload(url string) {

	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	file, err := os.Create("my.JPG")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {

	url := os.Args[1]

	fileDownload(getURLDownload(url))

}
