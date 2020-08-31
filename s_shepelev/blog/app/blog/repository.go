package blog

import (
	"github.com/Toringol/group-golang-1/tree/master/s_shepelev/blog/app/model"
)

// Repository - funcs interact with DataBase
type Repository interface {
	ListPosts() ([]*model.Post, error)
	SelectPostByID(int64) (*model.Post, error)
	CreatePost(*model.Post) (int64, error)
	UpdatePost(*model.Post) (int64, error)
	DeletePost(int64) (int64, error)
}
