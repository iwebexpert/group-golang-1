package core

import (
	"encoding/json"
	"errors"
	"io/ioutil"
)

var (
	PostNotExistsError = errors.New("post does not exist")
	EmptyFieldsError = errors.New("one or more required fileds are empty")
)

type Post struct {
	Id			int		`json:"id"`
	Title		string	`json:"title"`
	Text		string	`json:"text"`
	// Created		string	`json:"created"
}

type BlogPosts struct {
	Posts 		[]Post 	`json:"posts"`
}

func (bp *BlogPosts) PostExists(id int) bool {
	if id < 0 || id >= len(bp.Posts) {
		return false
	}
	return true
}

func (bp *BlogPosts) GetPostById(id int) (Post, error) {
	id--
	if !bp.PostExists(id) {
		return Post{}, PostNotExistsError
	}

	return bp.Posts[id], nil
}

func (bp *BlogPosts) Load(src string) error {
	b, err := ioutil.ReadFile(src)
	if err != nil {
		return err
	}

	err = json.Unmarshal(b, &bp)
	if err != nil {
		return err
	}

	return nil
}

func (bp *BlogPosts) Save(dst string) error {
	data, err := json.Marshal(bp)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(dst, data, 0644)
	if err != nil {
		return err
	}

	return nil
}

func (bp *BlogPosts) Reverse() []Post {
	result := []Post{}
	for i := len(bp.Posts)-1; i >= 0; i-- {
		result = append(result, bp.Posts[i])
	}

	return result
}

func (bp *BlogPosts) AddPost(title, text string) (int, error) {

	if title == "" || text == "" {
		return -1, EmptyFieldsError
	}

	id := len(bp.Posts)+1
	post := Post{
		Id: id,
		Title: title,
		Text: text,
	}

	bp.Posts = append(bp.Posts, post)
	return id, nil
}

func (bp *BlogPosts) EditPost(id int, newTitle, newText string) error {
	id--
	if !bp.PostExists(id) {
		return PostNotExistsError
	}

	bp.Posts[id] = Post{
		Title: newTitle,
		Text: newText,
	}

	return nil
}