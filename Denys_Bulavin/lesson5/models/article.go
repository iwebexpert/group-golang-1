package models

import (
	"github.com/astaxie/beego/orm"
)

type Articles struct {
	Id      uint64
	Title   string
	Article string
	Tags    string
}

type GetForm struct {
	Title   string `form:"art-title"`
	Article string `form:"art-article"`
	Tags    string `form:"art-tags"`
}

func (a *Articles) TableName() string {
	return "articles"
}

func NewArticle(a *GetForm) (*Articles, error) {
	//	if s.Article == "" {
	//		return nil, fmt.Errorf("Empty article title")
	//	}

	return &Articles{Title: a.Title, Article: a.Article, Tags: a.Tags}, nil
}

func init() {
	orm.RegisterModel(new(Articles))
}
