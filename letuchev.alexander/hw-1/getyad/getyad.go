package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

//YaFileInfo -
type YaFileInfo struct {
	Link string `json:"file"`
	Name string `json:"name"`
}

func main() {
	var fileInfo YaFileInfo

	const yaAPIlink = "https://cloud-api.yandex.net/v1/disk/public/resources?public_key="

	resp, err := http.Get(yaAPIlink + os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	respByte, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = json.Unmarshal(respByte, &fileInfo)
	if err != nil {
		fmt.Println(err)
		return
	}

	resp, err = http.Get(fileInfo.Link)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	outFile, err := os.Create(fileInfo.Name)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer outFile.Close()

	_, err = io.Copy(outFile, resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Файл", fileInfo.Name, "скачан успешно")
	return
}
