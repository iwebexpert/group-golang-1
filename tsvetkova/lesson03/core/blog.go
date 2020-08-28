package core

import (
	"encoding/json"
	"errors"
	"io/ioutil"
)

var PostNotExistsError = errors.New("post does not exist")

type Post struct {
	Id			int		`json:"id"`
	Title		string	`json:"title"`
	Text		string	`json:"text"`
	// Created		string	`json:"created"
}

type BlogPosts struct {
	Posts 		[]Post 	`json:"posts"`
}

func (bp *BlogPosts) GetPostById(id int) (Post, error) {
	id--
	if id < 0 || id >= len(bp.Posts) {
		return Post{}, PostNotExistsError
	}

	return bp.Posts[id], nil
}

func (bp *BlogPosts) LoadPostsFrom(src string) error {
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

func (bp *BlogPosts) AddPost(title, text string) int {
	id := len(bp.Posts)
	post := Post{
		Id: id,
		Title: title,
		Text: text,
	}

	bp.Posts = append(bp.Posts, post)
	return id+1
}