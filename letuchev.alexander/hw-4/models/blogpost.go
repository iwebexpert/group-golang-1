package models

import (
	"html/template"
	"time"
)

//BlogPost -
type BlogPost struct {
	ID         int           `json:"id"`
	About      string        `json:"about"`
	Text       template.HTML `json:"text"`
	Labels     []string      `json:"labels"`
	PublicDate time.Time     `json:"publicated"`
}

//BlogPostArray -
type BlogPostArray []BlogPost
