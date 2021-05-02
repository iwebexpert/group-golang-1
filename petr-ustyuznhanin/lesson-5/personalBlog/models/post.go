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

func NewPost(text string) (*Posts, error) {
	if text == "" {
		return nil, fmt.Errorf("Empty post text")
	}

	return &Posts{Text: text}, nil
}

func init() {
	orm.RegisterModel(new(Posts))
}
