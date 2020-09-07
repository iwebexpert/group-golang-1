package blog

import (
	"github.com/Toringol/group-golang-1/tree/master/s_shepelev/blog/app/model"
)

// Usecase - funcs interact with Posts
type Usecase interface {
	ListPosts() ([]*model.Post, error)
	SelectPostByID(string) (*model.Post, error)
	CreatePost(*model.Post) error
	UpdatePost(*model.Post) error
	DeletePost(string) error
}
