package models

import (
	"fmt"

	"github.com/astaxie/beego/orm"
)

// Posts - структура связанная с таблицей в БД
type Posts struct {
	ID     uint64
	Header string
	Text   string
}

// TableName - получаем имя таблицы для дальнейшего использования
func (p *Posts) TableName() string {
	return "posts"
}

// NewPost - создание нового поста
func NewPost(text string) (*Posts, error) {
	if text == "" {
		return nil, fmt.Errorf("Empty post title")
	}

	return &Posts{Text: text}, nil
}

func init() {
	orm.RegisterModel(new(Posts))
}
