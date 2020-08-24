package http

import (
	"encoding/json"
	"html/template"
	"io/ioutil"
	"net/http"
	"path"
	"strconv"

	"github.com/Toringol/group-golang-1/tree/master/s_shepelev/blog/app/blog"
	"github.com/Toringol/group-golang-1/tree/master/s_shepelev/blog/app/model"
	"github.com/go-chi/chi"
	"github.com/russross/blackfriday"
)

var (
	layoutTemplatePosts = template.Must(template.ParseFiles(path.Join("../templates", "layout.html"),
		path.Join("../templates", "posts.html")))
	layoutTemplateActualPost = template.Must(template.ParseFiles(path.Join("../templates", "layout.html"),
		path.Join("../templates", "post.html")))
	layoutTemplateAddPost = template.Must(template.ParseFiles(path.Join("../templates", "layout.html"),
		path.Join("../templates", "newpost.html")))
)

// blogHandlers - http handlers structure
type blogHandlers struct {
	usecase blog.Usecase
}

// NewBlogHandler - deliver our handlers in http
func NewBlogHandler(router *chi.Mux, us blog.Usecase) {
	handlers := blogHandlers{usecase: us}

	// Blog handlers
	router.Get("/", handlers.getAllPostsHandlers)
	router.Get("/posts/{postID}", handlers.getPostInfoHandler)
	router.Post("/posts/{postID}", handlers.changePostHandler)
	router.Get("/addPost", handlers.getNewPostHandler)
	router.Post("/addPost", handlers.createNewPostHandler)
}

func (bh *blogHandlers) getAllPostsHandlers(w http.ResponseWriter, r *http.Request) {
	posts, err := bh.usecase.ListPosts()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := layoutTemplatePosts.ExecuteTemplate(w, "layout", posts); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (bh *blogHandlers) getPostInfoHandler(w http.ResponseWriter, r *http.Request) {
	postID, err := strconv.ParseInt(chi.URLParam(r, "postID"), 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	post, err := bh.usecase.SelectPostByID(postID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := layoutTemplateActualPost.ExecuteTemplate(w, "layout", post); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (bh *blogHandlers) changePostHandler(w http.ResponseWriter, r *http.Request) {
	postID, err := strconv.ParseInt(chi.URLParam(r, "postID"), 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	oldPostData, err := bh.usecase.SelectPostByID(postID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	postData := new(model.Post)
	postData.ID = postID

	if err := json.Unmarshal(body, postData); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if postData.Author == "" {
		postData.Author = oldPostData.Author
	}
	if postData.Description == "" {
		postData.Description = oldPostData.Description
	}
	if postData.Title == "" {
		postData.Title = oldPostData.Title
	}

	postData.Description = template.HTML(blackfriday.Run([]byte(postData.Description)))

	_, err = bh.usecase.UpdatePost(postData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (bh *blogHandlers) getNewPostHandler(w http.ResponseWriter, r *http.Request) {
	if err := layoutTemplateAddPost.ExecuteTemplate(w, "layout", ""); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (bh *blogHandlers) createNewPostHandler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	postData := new(model.Post)

	if err := json.Unmarshal(body, postData); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	postData.Description = template.HTML(blackfriday.Run([]byte(postData.Description)))

	_, err = bh.usecase.CreatePost(postData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
