package models

import (
	"fmt"

	"github.com/astaxie/beego/orm"
)

type Posts struct {
	Id    uint64
	Title string
	Text  string
}

func (p *Posts) TableName() string {
	return "posts"
}

func NewPost(title, text string) (*Posts, error) {
	if title == "" {
		return nil, fmt.Errorf("Empty post title")
	}
	if text == "" {
		return nil, fmt.Errorf("Empty post text")
	}

	return &Posts{Title: title, Text: text}, nil
}

func UpdatePost(title, text string, id uint64) (*Posts, error) {
	if title == "" {
		return nil, fmt.Errorf("Empty post title")
	}
	if text == "" {
		return nil, fmt.Errorf("Empty post text")
	}

	return &Posts{Title: title, Text: text}, nil
}

func init() {
	orm.RegisterModel(new(Posts))
}
