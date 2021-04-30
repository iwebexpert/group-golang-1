package models

import (
	"fmt"

	"github.com/astaxie/beego/orm"
)

type Tasks struct {
	Id   uint64
	Text string
}

func (t *Tasks) TableName() string {
	return "tasks"
}

func NewTask(text string) (*Tasks, error) {
	if text == "" {
		return nil, fmt.Errorf("Empty task title")
	}

	return &Tasks{Text: text}, nil
}

func init() {
	orm.RegisterModel(new(Tasks))
}
