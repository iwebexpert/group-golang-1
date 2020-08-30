package server

import (
	"blog/models"
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"

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
	srv.router.Route("/api/v1/reload/", func(r chi.Router) {
		r.Get("/", func(w http.ResponseWriter, r *http.Request) { srv.BlogReload(w, r) })
	})
}

//GetPage - Рендеринг страницы
func (srv *BlogServer) GetPage(w http.ResponseWriter, r *http.Request, pageName string) {
	file, _ := os.Open("./templates/" + pageName + ".html")
	data, _ := ioutil.ReadAll(file)

	fmt.Println("Подготовка шаблона " + pageName)
	tmp, err := template.New(pageName).Parse(string(data))
	if err != nil {
		fmt.Println("Ошибка!", err)
		w.Write([]byte(fmt.Sprint(err)))
		return
	}
	fmt.Println("Наполнение шаблона " + pageName + " данными")
	tmp.ExecuteTemplate(w, pageName, srv)
}

//BlogReload - Перезагрузка списка постов в памяти сервера из БД и возврат их полного списка
func (srv *BlogServer) BlogReload(w http.ResponseWriter, r *http.Request) {
	posts, err := models.Retrieve(srv.Ctx, srv.DBMongo)
	if err != nil {
		fmt.Println(err)
		return
	}
	srv.Posts = *posts
	data, _ := json.Marshal(posts)
	w.Write(data)
}

//PostCreate -
func (srv *BlogServer) PostCreate(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Запрос на вставку нового поста")
	data, _ := ioutil.ReadAll(r.Body)

	post := models.BlogPost{}
	err := json.Unmarshal(data, &post)
	if err != nil {
		fmt.Println("Ошибка парсинга json:", err)
		return
	}

	id, err := srv.Posts.NewBlogPost(srv.Ctx, srv.DBMongo, post.About, post.Text, post.Labels)
	if err != nil {
		fmt.Println(err)
		return
	}
	data, err = json.Marshal(srv.Posts[id])
	if err != nil {
		fmt.Println("Ошибка формирования json:", err)
		return
	}

	w.Write(data)
}

//PostGet -
func (srv *BlogServer) PostGet(w http.ResponseWriter, r *http.Request) {
	postID := chi.URLParam(r, "postID")
	post, ok := srv.Posts[postID]
	if !ok {
		fmt.Println("Запрошенный post.ID не найден!")
		return
	}

	data, err := json.Marshal(post)
	if err != nil {
		fmt.Println("Ошибка формирования json:", err)
		return
	}

	w.Write(data)
}

//PostDelete -
func (srv *BlogServer) PostDelete(w http.ResponseWriter, r *http.Request) {
	postID := chi.URLParam(r, "postID")
	fmt.Print("Запрос удаления post.ID=", postID, "\r\n")
	post, ok := srv.Posts[postID]
	if !ok {
		fmt.Print("Запрошенный для удаления post.ID=", postID, " не найден!\r\n")
		return
	}

	data, err := json.Marshal(post)
	if err != nil {
		fmt.Println("Ошибка формирования json:", err)
		return
	}

	if err := srv.Posts.DeleteBlogPost(srv.Ctx, srv.DBMongo, postID); err != nil {
		fmt.Println(err)
		return
	}

	w.Write(data)
}

//PostUpdate -
func (srv *BlogServer) PostUpdate(w http.ResponseWriter, r *http.Request) {
	postID := chi.URLParam(r, "postID")
	fmt.Print("Запрос обновления post.ID=", postID, "\r\n")
	_, ok := srv.Posts[postID]
	if !ok {
		fmt.Print("Запрошенный для обновления post.ID=", postID, " не найден!\r\n")
		return
	}
	data, _ := ioutil.ReadAll(r.Body)
	post := models.BlogPost{}
	err := json.Unmarshal(data, &post)
	if err != nil {
		fmt.Println("Ошибка парсинга json:", err)
		return
	}

	err = srv.Posts.UpdateBlogPost(srv.Ctx, srv.DBMongo, postID, post.About, post.Text, post.Labels)
	if err != nil {
		fmt.Println(err)
		return
	}

	data, err = json.Marshal(post)
	if err != nil {
		fmt.Println("Ошибка формирования json:", err)
		return
	}

	w.Write(data)
}
