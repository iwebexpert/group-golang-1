package model

import (
	"html/template"
)

// Post - structure record in DataBase
type Post struct {
	ID          string        `json:"id,omitempty" bson:"_id,omitempty"`
	Title       string        `json:"title" bson:"title"`
	Description template.HTML `json:"description" bson:"description"`
	Author      string        `json:"author" bson:"author"`
}
