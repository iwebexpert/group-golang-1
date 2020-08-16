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

const diskAPIPublicURL = "https://cloud-api.yandex.net/v1/disk/public/resources?public_key="
const diskAPIPathURL = "https://cloud-api.yandex.net/v1/disk/resources/download?path="
const OAuthKey = "AgAEA7qjSNZVAAaMT3ljqBoh0kI_h2GoTA9w1dM" //Так делать нельзя! UNSAFE!

var url = "https://yadi.sk/d/Sbjmcqfgl4wZZQ"
var path = "%2FYaDiskFile"

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
func getPublicLink(url string) string {
	var link map[string]interface{}

	resp, err := http.Get(diskAPIPublicURL + url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(respBody, &link)
	if err != nil {
		log.Fatal(err)
	}
	srtURL := link["file"].(string)

	return srtURL

}
func getPathLink(filePath string, key string) string {
	var data map[string]interface{}
	timeout := time.Duration(5 * time.Second)
	client := http.Client{
		Timeout: timeout,
	}

	//reqest to get link from YaDisk

	req, err := http.NewRequest("GET", diskAPIPathURL+filePath, nil)
	req.Header.Set("Authorization", key)
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

	return data["href"].(string)
}
func main() {

	filename := "YaDiskFile"
	choice := 0

	fmt.Println("New file name:")
	fmt.Scanln(&filename)
	for choice > 2 || choice < 1 {
		fmt.Println("type 1 for downloading from link")
		fmt.Println("type 2 for downloading from path on your disk")
		fmt.Scanln(&choice)
	}
	switch choice {
	case 1:
		fmt.Println("Place public link:")
		fmt.Scanln(&url)
		err := DownloadFile(filename, getPublicLink(url))
		if err != nil {
			log.Fatal(err)
		}
	case 2:
		fmt.Println("Name of file on disk? (Try YaDiskFile):") //YaDiskFile
		fmt.Scanln(&path)
		err := DownloadFile(filename, getPathLink(path, OAuthKey))
		if err != nil {
			log.Fatal(err)
		}
	}
	fmt.Println("File", filename, "downloaded")
}
