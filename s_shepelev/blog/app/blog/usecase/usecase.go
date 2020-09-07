package usecase

import (
	"github.com/Toringol/group-golang-1/tree/master/s_shepelev/blog/app/blog"
	"github.com/Toringol/group-golang-1/tree/master/s_shepelev/blog/app/model"
)

// blogUsecase - connector database and post
type blogUsecase struct {
	repo blog.Repository
}

// NewBlogUsecase - create new usercase
func NewBlogUsecase(blogRepo blog.Repository) blog.Usecase {
	return blogUsecase{repo: blogRepo}
}

// ListPosts - return all posts in DB
func (bu blogUsecase) ListPosts() ([]*model.Post, error) {
	return bu.repo.ListPosts()
}

// SelectPostByID - return post by id
func (bu blogUsecase) SelectPostByID(id string) (*model.Post, error) {
	return bu.repo.SelectPostByID(id)
}

// CreatePost - create new record in DB posts
func (bu blogUsecase) CreatePost(post *model.Post) error {
	return bu.repo.CreatePost(post)
}

// UpdatePost - update post`s data in DataBase
func (bu blogUsecase) UpdatePost(post *model.Post) error {
	return bu.repo.UpdatePost(post)
}

// DeletePost - delete post`s record in DataBase
func (bu blogUsecase) DeletePost(id string) error {
	return bu.repo.DeletePost(id)
}
