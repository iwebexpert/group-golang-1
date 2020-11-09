package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type FileT struct {
	FileLink string `json:"file"`
	Error *string `json:"error"`
	ErrorMessage *string `json:"message"`
}

const linkBase = "https://cloud-api.yandex.net/v1/disk/public/resources?public_key="

func GetFile(info *FileT) (io.ReadCloser, error) {
	if info.Error != nil {
		return nil, fmt.Errorf("disk error: %s: %s", info.Error, info.ErrorMessage)
	}

	response, err := http.Get(info.FileLink)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// GetData заполняет структуру FileT
func GetData(key string) (*FileT, error) {
	link := fmt.Sprintf("%s%s", linkBase, key)

	response, err := http.Get(link)
	if err != nil {
		return nil, err
	}

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	file := &FileT{}
	if err := json.Unmarshal(data, file); err != nil {
		return nil, err
	}

	return file, nil
}

func SaveFile(key, path string) error {
	info, err := GetData(key)
	if err != nil {
		return err
	}

	stream, err := GetFile(info)
	if err != nil {
		return err
	}
	defer stream.Close()

	file, err := os.Create(path)
	if err != nil {
		return err
	}

	if _, err := io.Copy(file, stream); err != nil {
		return err
	}
	return nil
}

func main() {
	key := flag.String("key", "https://yadi.sk/d/Sbjmcqfgl4wZZQ", "path or filekey")
	name := flag.String("name", "./file.txt", "path where file stored")
	flag.Parse()
	if err := SaveFile(*key, *name); err != nil {
		log.Fatalf("error: %s", err)
	}
	log.Println("Everything is ok!")
}
