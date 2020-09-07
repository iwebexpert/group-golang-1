package main

import (
	"flag"
	"fmt"
	"net/http"
)

// define flags:
// yandexToken allow access to cloud-api,
// downloadFile - path to file
var (
	yandexToken  = flag.String("token", "", "token of your yandex disk")
	downloadFile = flag.String("file", "", "saving file")
	yandexUrl    = "https://cloud-api.yandex.net/v1/disk"
)

// yaDiskSaver - save public file to user`s disk in download folder
func yaDiskSaver(publicFileUrl string) error {
	yandexSaveAPI := fmt.Sprintf("%s/%s", yandexUrl, "public/resources/save-to-disk")

	req, err := http.NewRequest("POST", yandexSaveAPI, nil)
	if err != nil {
		return err
	}

	q := req.URL.Query()
	q.Add("public_key", publicFileUrl)
	req.URL.RawQuery = q.Encode()

	req.Header.Add("Authorization", *yandexToken)

	client := &http.Client{}

	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	return nil
}

func main() {
	flag.Parse()

	if *yandexToken == "" || *downloadFile == "" {
		fmt.Println("You need to write yandexToken and public url to file")
		return
	}

	err := yaDiskSaver(*downloadFile)
	if err != nil {
		fmt.Println(err)
		return
	}
}
