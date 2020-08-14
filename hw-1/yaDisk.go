package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type downloadConfig struct{
	filePath string
	OAuthKey string
	APIPath string
}

func main(){
	config := downloadConfig{
		filePath: "%2FДокумент.docx",
		OAuthKey: "OAuth AgAAAAAmNiiYAAaM9Nuh_jjVV0spsX0IOckUBfQ",
		APIPath: "https://cloud-api.yandex.net/v1/disk/resources/download?path=",
	}

	client := &http.Client{}
	req, err := http.NewRequest(
		"GET",
		config.APIPath+config.filePath,
		//fmt.Sprintf("https://cloud-api.yandex.net/v1/disk/resources/download?path=#{config.myPath}", config.myPath),
		nil,
	)
	req.Header.Add("Autorization", config.OAuthKey)
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	io.Copy(os.Stdout, resp.Body)
}