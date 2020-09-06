package server

import (
	"blog/models"
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"

	"github.com/go-chi/chi"
)

func (srv *BlogServer) defineRoutes() {
	srv.router.Route("/", func(r chi.Router) {
		r.Get("/", func(w http.ResponseWriter, r *http.Request) { srv.GetPage(w, r, "index") })
		r.Get("/blog/", func(w http.ResponseWriter, r *http.Request) { srv.GetPage(w, r, "blog") })
	})
	srv.router.Route("/api/v1/blog/", func(r chi.Router) {
		r.Post("/", func(w http.ResponseWriter, r *http.Request) { srv.PostCreate(w, r) })
		r.Put("/{postID}", func(w http.ResponseWriter, r *http.Request) { srv.PostUpdate(w, r) })
		r.Delete("/{postID}", func(w http.ResponseWriter, r *http.Request) { srv.PostDelete(w, r) })
		r.Get("/{postID}", func(w http.ResponseWriter, r *http.Request) { srv.PostGet(w, r) })
	})
	srv.router.Route("/api/v1/reload", func(r chi.Router) {
		r.Get("/", func(w http.ResponseWriter, r *http.Request) { srv.BlogReload(w, r) })
	})
}

//BlogReload -
func (srv *BlogServer) BlogReload(w http.ResponseWriter, r *http.Request) {
	srv.log.Debug("Запрошена принудительная перезагрузка постов из БД")
	posts, err := models.Retrieve(srv.DBlink, srv.log)
	if err != nil {
		srv.log.Error(err)
		return
	}
	srv.Posts = *posts
	data, _ := json.Marshal(posts)
	w.Write(data)
}

//GetPage -
func (srv *BlogServer) GetPage(w http.ResponseWriter, r *http.Request, pageName string) {
	srv.log.Debug("Запрошена страница ", pageName)
	file, err := os.Open("../templates/" + pageName + ".html")
	if err != nil {
		srv.log.Error("Ошибка!", err)
		w.Write([]byte(fmt.Sprintf("Error: %s", err)))
		return
	}
	data, err := ioutil.ReadAll(file)
	if err != nil {
		srv.log.Error("Ошибка!", err)
		w.Write([]byte(fmt.Sprintf("Error: %s", err)))
		return
	}
	srv.log.Debug("Подготовка шаблона " + pageName)
	tmp, err := template.New(pageName).Parse(string(data))
	if err != nil {
		srv.log.Error("Ошибка!", err)
		w.Write([]byte(fmt.Sprintf("Error: %s", err)))
		return
	}
	srv.log.Debug("Наполнение шаблона " + pageName + " данными")
	tmp.ExecuteTemplate(w, pageName, srv)
}

//PostCreate -
func (srv *BlogServer) PostCreate(w http.ResponseWriter, r *http.Request) {
	srv.log.Debug("Запрошено создание поста")
	data, _ := ioutil.ReadAll(r.Body)

	post := models.BlogPost{}
	_ = json.Unmarshal(data, &post)

	id, err := srv.Posts.NewBlogPost(srv.DBlink, post.About, post.Text, post.Labels)
	if err != nil {
		srv.log.Error(err)
		return
	}
	data, _ = json.Marshal(srv.Posts[id])
	w.Write(data)
}

//PostGet -
func (srv *BlogServer) PostGet(w http.ResponseWriter, r *http.Request) {
	srv.log.Debug("Запрошено чтение поста")
	postID, err := strconv.ParseInt(chi.URLParam(r, "postID"), 10, 32)
	if err != nil {
		srv.log.Error(err)
		return
	}
	data, _ := ioutil.ReadAll(r.Body)
	post := models.BlogPost{}
	_ = json.Unmarshal(data, &post)
	post.ID = int(postID)

	post2, ok := srv.Posts[post.ID]
	if !ok {
		srv.log.Debug("Запрошенный post.ID не найден!")
		return
	}
	data, _ = json.Marshal(post2)
	w.Write(data)
}

//PostDelete -
func (srv *BlogServer) PostDelete(w http.ResponseWriter, r *http.Request) {
	srv.log.Debug("Запрошено удаление поста")
	postID, err := strconv.ParseInt(chi.URLParam(r, "postID"), 10, 32)
	if err != nil {
		srv.log.Error(err)
		return
	}
	data, _ := ioutil.ReadAll(r.Body)
	post := models.BlogPost{}
	_ = json.Unmarshal(data, &post)
	post.ID = int(postID)

	if err := srv.Posts.DeleteBlogPost(srv.DBlink, post.ID); err != nil {
		srv.log.Error(err)
		return
	}
}

//PostUpdate -
func (srv *BlogServer) PostUpdate(w http.ResponseWriter, r *http.Request) {
	srv.log.Debug("Запрошено обновление поста")
	postID, err := strconv.ParseInt(chi.URLParam(r, "postID"), 10, 32)
	if err != nil {
		srv.log.Error(err)
		return
	}
	data, _ := ioutil.ReadAll(r.Body)
	post := models.BlogPost{}
	_ = json.Unmarshal(data, &post)
	post.ID = int(postID)

	err = srv.Posts.UpdateBlogPost(srv.DBlink, post.ID, post.About, post.Text, post.Labels)
	if err != nil {
		srv.log.Error(err)
		return
	}
	data, _ = json.Marshal(srv.Posts[post.ID])
	w.Write(data)
}
