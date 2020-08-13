package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

func DownloadFile(path string, url string) error {
	// Get data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create file
	out, err := os.Create(path)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write to file
	_, err = io.Copy(out, resp.Body)
	return err
}
func main() {
	var data map[string]interface{}
	timeout := time.Duration(5 * time.Second)
	client := http.Client{
		Timeout: timeout,
	}

	//reqest to get link from YaDisk

	req, err := http.NewRequest("GET", "https://cloud-api.yandex.net/v1/disk/resources/download?path=%2FYaDiskFile", nil)
	req.Header.Set("Authorization", "AgAEA7qjSNZVAAaMT3ljqBoh0kI_h2GoTA9w1dM")
	if err != nil {
		log.Fatal(err)
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	//Parsing response from YaDisk to map
	bData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	if err := json.Unmarshal(bData, &data); err != nil {
		log.Fatal(err)
	}
	//Place link from map to var
	href := data["href"].(string)
	//fmt.Println(href)

	errDownload := DownloadFile("fileDownloaded", href)
	if errDownload != nil {
		log.Fatal(err)
	}
	fmt.Println("File downloaded")
}
