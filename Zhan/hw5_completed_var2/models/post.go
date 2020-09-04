package models

import (
	"fmt"

	"github.com/astaxie/beego/orm"
)

// Posts - структура связанная с таблицей в БД
type Posts struct {
	Id     uint64
	Header string
	Text   string
}

// TableName - получаем имя таблицы для дальнейшего использования
func (p *Posts) TableName() string {
	return "posts"
}

// NewUpdPost - создание нового или редактирование существующего поста
func NewUpdPost(header, text string, id uint64) (*Posts, error) {
	if header == "" {
		return nil, fmt.Errorf("Empty post header")
	}
	return &Posts{Id: id, Text: text, Header: header}, nil
}

func init() {
	orm.RegisterModel(new(Posts))
}
