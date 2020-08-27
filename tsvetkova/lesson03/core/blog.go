package core

import (
	"encoding/json"
	"io/ioutil"
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
