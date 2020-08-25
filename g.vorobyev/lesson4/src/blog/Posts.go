package blog

import (
	"dbo"
	"errors"
	"github.com/gin-gonic/gin"
	"model"
	"net/http"
	"strconv"
)

func getId(rc *gin.Context) (uint, error) {
	var pt model.PageType
	if err := rc.ShouldBindUri(&pt); err != nil {
		return 0, errors.New("ERROR: Id is not set")
	}

	id, err := strconv.Atoi(pt.Id)
	if err != nil {
		return 0, errors.New("ERROR: Id is not integer")
	}

	return uint(id), nil
}

func ShowPosts(rc *gin.Context) {
	var Posts []model.Post

	dbo.DB.Find(&Posts)
	rc.HTML(http.StatusOK, "index.tmpl", gin.H{"title": "Home", "data": Posts})
}

func ShowPost(rc *gin.Context) {
	var Post model.Post

	id, err := getId(rc)
	if err != nil {
		rc.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	dbo.DB.Where(&model.Post{PostID: id}).First(&Post)

	// weird gostyle to check if its empty
	// WHY NOT JUST Post.isEmpty()
	if Post == (model.Post{}) {
		rc.JSON(http.StatusNotFound, gin.H{"post": "Not Found"})
		return
	}

	rc.HTML(http.StatusOK, "post.tmpl", gin.H{"title": "Show Post", "data": Post})
}

func AddPost(rc *gin.Context) {
	title := rc.PostForm("title")
	postMessage := rc.PostForm("post_message")

	PostType := model.Post{Title: title, PostData: postMessage}
	dbo.DB.Create(&PostType)

	rc.Redirect(http.StatusFound, "/")
}

func EditPost(rc *gin.Context) {
	var Post model.Post

	id, err := getId(rc)
	if err != nil {
		rc.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	title := rc.PostForm("title")
	postMessage := rc.PostForm("post_message")

	dbo.DB.Where(&model.Post{PostID: id}).First(&Post)

	// if we try edit unexist post
	if Post == (model.Post{}) {
		rc.JSON(http.StatusNotFound, gin.H{"post": "Not Found"})
		return
	}

	Post.Title = title
	Post.PostData = postMessage

	dbo.DB.Save(Post)

	rc.Redirect(301, "/")
}

func ShowAddPost(rc *gin.Context) {
	rc.HTML(http.StatusOK, "add_post.tmpl", gin.H{"title": "Add Post"})
}

func ShowEditPost(rc *gin.Context) {
	var Post model.Post
	id, err := getId(rc)
	if err != nil {
		rc.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	dbo.DB.Where(&model.Post{PostID: id}).First(&Post)

	if Post == (model.Post{}) {
		rc.JSON(http.StatusNotFound, gin.H{"post": "Not Found"})
		return
	}

	rc.HTML(http.StatusOK, "edit.tmpl", gin.H{"title": "Edit Post", "data": Post})
}
