package tools

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"os"
	"path"
	"strings"
	"sync"

	"github.com/russross/blackfriday"
)

type post struct {
	Title   string
	Body    template.HTML
	ModTime int64
}

type postArray struct {
	Items map[string]post
	sync.RWMutex
}

func NewPostArray() *postArray {
	p := postArray{}
	p.Items = make(map[string]post)
	return &p
}

func (p *postArray) InitPosts() error {
	files, err := ioutil.ReadDir("../posts")
	if err != nil {
		return err
	}

	for _, file := range files {
		_, _, err := p.Load(path.Join("../posts", file.Name()))
		if err != nil {
			return err
		}
	}

	return nil
}

// Load markdown file and convert it to html
// Return object Post or error
func (p *postArray) Load(md string) (post, int, error) {
	info, err := os.Stat(md)
	if err != nil {
		if os.IsNotExist(err) {
			return post{}, 404, err
		}
		return post{}, 500, err
	}
	if info.IsDir() {
		return post{}, 404, fmt.Errorf("dir")
	}
	val, ok := p.Items[md]
	if !ok || (ok && val.ModTime != info.ModTime().UnixNano()) {
		p.RLock()
		defer p.RUnlock()
		fileread, _ := ioutil.ReadFile(md)
		lines := strings.Split(string(fileread), "\n")
		title := string(lines[0])
		body := strings.Join(lines[1:len(lines)], "\n")
		body = string(blackfriday.Run([]byte(body)))
		p.Items[md] = post{title, template.HTML(body), info.ModTime().UnixNano()}
	}
	post := p.Items[md]
	return post, 200, nil
}
