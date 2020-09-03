package server

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"lesson6/models"
	"net/http"
	"os"

	"github.com/go-chi/chi"
)

func (srv *Server) defineRoutes() {
	srv.router.Route("/", func(r chi.Router) {
		r.Get("/", func(w http.ResponseWriter, r *http.Request) { srv.GetPage(w, r) })
	})
	srv.router.Route("/api/v1/posts/", func(r chi.Router) {
		r.Post("/", func(w http.ResponseWriter, r *http.Request) { srv.PostCreate(w, r) })
		r.Put("/{postID}", func(w http.ResponseWriter, r *http.Request) { srv.PostUpdate(w, r) })
		r.Delete("/{postID}", func(w http.ResponseWriter, r *http.Request) { srv.PostDelete(w, r) })
		r.Get("/{postID}", func(w http.ResponseWriter, r *http.Request) { srv.PostGet(w, r) })
	})
}
func (srv *Server) GetPage(w http.ResponseWriter, r *http.Request) {
	file, _ := os.Open("./templates/posts.html")
	data, _ := ioutil.ReadAll(file)

	tmp, err := template.New("posts").Parse(string(data))
	if err != nil {
		w.Write([]byte(fmt.Sprint(err)))
		return
	}
	NewPosts, err := models.Get(srv.Ctx, srv.DBMongo)
	if err != nil {
		srv.lg.WithError(err).Fatal("NewPost err")
		return
	}
	srv.Posts = *NewPosts
	tmp.ExecuteTemplate(w, "posts", srv)

}
func (srv *Server) PostCreate(w http.ResponseWriter, r *http.Request) {
	data, _ := ioutil.ReadAll(r.Body)

	post := models.Post{}
	_ = json.Unmarshal(data, &post)

	id, err := srv.Posts.NewPost(srv.Ctx, srv.DBMongo, post.Header, post.Text)
	if err != nil {
		srv.lg.WithError(err).Fatal("NewPost err")
		return
	}
	data, _ = json.Marshal(srv.Posts[id])

	w.Write(data)
}

func (srv *Server) PostGet(w http.ResponseWriter, r *http.Request) {
	postID := chi.URLParam(r, "postID")
	post, ok := srv.Posts[postID]
	if !ok {
		srv.lg.Warningln("no such ID=", postID)
		return
	}

	data, err := json.Marshal(post)
	if err != nil {
		srv.lg.WithError(err).Fatal("PostGet Marshal err")
		return
	}

	w.Write(data)
}

//PostDelete -
func (srv *Server) PostDelete(w http.ResponseWriter, r *http.Request) {
	postID := chi.URLParam(r, "postID")
	srv.lg.Debug("Del", postID)
	post, ok := srv.Posts[postID]
	if !ok {
		srv.lg.Warningln("no such ID=", postID)
		return
	}

	data, _ := json.Marshal(post)

	if err := srv.Posts.DeletePost(srv.Ctx, srv.DBMongo, postID); err != nil {
		fmt.Println(err)
		return
	}

	w.Write(data)
}

func (srv *Server) PostUpdate(w http.ResponseWriter, r *http.Request) {
	postID := chi.URLParam(r, "postID")
	srv.lg.Debug("upd ID=", postID)
	_, ok := srv.Posts[postID]
	if !ok {
		srv.lg.Warningln("no such ID=", postID)
		return
	}
	data, _ := ioutil.ReadAll(r.Body)
	post := models.Post{}
	_ = json.Unmarshal(data, &post)

	err := srv.Posts.UpdatePost(srv.Ctx, srv.DBMongo, postID, post.Header, post.Text)
	if err != nil {
		srv.lg.WithError(err).Fatal("UpdatePost err")
		return
	}

	data, _ = json.Marshal(post)

	w.Write(data)
}
