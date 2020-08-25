package blog

import (
	"github.com/Toringol/group-golang-1/tree/master/s_shepelev/blog/app/model"
)

// Usecase - funcs interact with Posts
type Usecase interface {
	ListPosts() ([]*model.Post, error)
	SelectPostByID(int64) (*model.Post, error)
	CreatePost(*model.Post) (int64, error)
	UpdatePost(*model.Post) (int64, error)
	DeletePost(int64) (int64, error)
}
