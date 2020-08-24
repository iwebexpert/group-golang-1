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
	uuid "github.com/satori/go.uuid"
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
		fileread, err := ioutil.ReadFile(md)
		if err != nil {
			return post{}, 500, err
		}
		lines := strings.Split(string(fileread), "\n")
		title := string(lines[0])
		body := strings.Join(lines[1:len(lines)], "\n")
		body = string(blackfriday.Run([]byte(body)))
		p.Items[md] = post{title, template.HTML(body), info.ModTime().UnixNano()}
	}
	post := p.Items[md]
	return post, 200, nil
}

// ChangePost - write new data to file and change data in posts structure
func (p *postArray) ChangePost(postName, title, text string) error {
	data := []byte(title + "\n" + text)

	err := ioutil.WriteFile(postName, data, 0644)
	if err != nil {
		return err
	}

	p.RLock()
	defer p.RUnlock()

	info, err := os.Stat(postName)
	if err != nil {
		return err
	}

	body := string(blackfriday.Run([]byte(text + "\n")))

	p.Items[postName] = post{title, template.HTML(body), info.ModTime().UnixNano()}

	return nil
}

func (p *postArray) CreatePost(title, text string) error {

	postNamePrefix := "../posts/"
	fileName := uuid.NewV4()
	postName := postNamePrefix + fileName.String() + ".md"

	fd, err := os.Create(postName)
	if err != nil {
		return err
	}

	data := title + "\n" + text

	_, err = fd.WriteString(data)
	if err != nil {
		return err
	}

	p.RLock()
	defer p.RUnlock()

	info, err := os.Stat(postName)
	if err != nil {
		return err
	}

	body := string(blackfriday.Run([]byte(text + "\n")))

	p.Items[postName] = post{title, template.HTML(body), info.ModTime().UnixNano()}

	return nil
}
