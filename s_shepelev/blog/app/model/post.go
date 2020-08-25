package model

import "html/template"

// Post - structure record in DataBase
type Post struct {
	ID          int64         `json:"id,omitempty"`
	Title       string        `json:"title"`
	Description template.HTML `json:"description"`
	Author      string        `json:"author"`
}
