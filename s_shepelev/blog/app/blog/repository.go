package blog

import (
	"github.com/Toringol/group-golang-1/tree/master/s_shepelev/blog/app/model"
)

// Repository - funcs interact with DataBase
type Repository interface {
	ListPosts() ([]*model.Post, error)
	SelectPostByID(string) (*model.Post, error)
	CreatePost(*model.Post) error
	UpdatePost(*model.Post) error
	DeletePost(string) error
}
