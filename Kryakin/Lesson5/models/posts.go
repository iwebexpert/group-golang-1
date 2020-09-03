package models

import (
	"fmt"
	"time"

	"github.com/astaxie/beego/orm"
)

type Posts struct {
	Id     uint64
	Header string
	Text   string
	Date   string
}

func (p *Posts) TableName() string {
	return "posts"
}

func NewPost(header, text string) (*Posts, error) {
	if text == "" {
		return nil, fmt.Errorf("Empty post header")
	}
	t := time.Now()
	return &Posts{Text: text, Header: header, Date: t.Format("2006-01-02 15:04:05")}, nil
}
func ExPost(header, text string, id uint64) (*Posts, error) {
	if text == "" {
		return nil, fmt.Errorf("Empty post header")
	}
	t := time.Now()
	return &Posts{Id: id, Text: text, Header: header, Date: t.Format("2006-01-02 15:04:05")}, nil
}
func DelPost(id uint64) (*Posts, error) {
	return &Posts{Id: id}, nil
}

func init() {
	orm.RegisterModel(new(Posts))
}
