package server

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"server/models"
	"strconv"

	"github.com/go-chi/chi"
)

func (serv *Server) getTemplateHandler(w http.ResponseWriter, r *http.Request) {
	templateName := chi.URLParam(r, "template")
	postID := r.URL.Query().Get("ID")
	if templateName == "" {
		templateName = serv.indexTemplate
	}

	file, err := os.Open(path.Join(serv.rootDir, serv.templatesDir, templateName))
	if err != nil {
		if err == os.ErrNotExist {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		serv.SendInternalErr(w, err)
		return
	}

	data, err := ioutil.ReadAll(file)
	if err != nil {
		serv.SendInternalErr(w, err)
		return
	}

	templ, err := template.New("Page").Parse(string(data))
	if err != nil {
		serv.SendInternalErr(w, err)
		return
	}
	if postID != "" {
		postID, _ := strconv.Atoi(postID)
		Posts, err := models.GetPostItem(serv.db, postID)
		if err != nil {
			serv.SendInternalErr(w, err)
			return
		}
		serv.Page.Posts = Posts
		if err := templ.Execute(w, serv.Page); err != nil {
			serv.SendInternalErr(w, err)
			return
		}
	} else {
		Posts, err := models.GetAllPostItems(serv.db)
		if err != nil {
			serv.SendInternalErr(w, err)
			return
		}
		serv.Page.Posts = Posts
		if err := templ.Execute(w, serv.Page); err != nil {
			serv.SendInternalErr(w, err)
			return
		}
	}

}

//NewPostHandler creates new Post in DB
func (serv *Server) NewPostHandler(w http.ResponseWriter, r *http.Request) {
	data, _ := ioutil.ReadAll(r.Body)

	post := models.PostItem{}
	_ = json.Unmarshal(data, &post)

	if err := post.Insert(serv.db); err != nil {
		serv.SendInternalErr(w, err)
		return
	}

	data, _ = json.Marshal(post)
	w.Write(data)
}

//Редактирование задачи
func (serv *Server) putPostHandler(w http.ResponseWriter, r *http.Request) {
	postID, _ := strconv.Atoi(chi.URLParam(r, "id"))
	data, _ := ioutil.ReadAll(r.Body)

	post := models.PostItem{}
	_ = json.Unmarshal(data, &post)
	post.ID = postID
	fmt.Println(post)
	if err := post.Update(serv.db); err != nil {
		serv.SendInternalErr(w, err)
		return
	}

	data, _ = json.Marshal(post)
	w.Write(data)
}

//Удаление задачи
func (serv *Server) deletePostHandler(w http.ResponseWriter, r *http.Request) {
	postID, _ := strconv.Atoi(chi.URLParam(r, "id"))
	//fmt.Println(taskID, "deleteTaskHandler")

	post := models.PostItem{ID: postID}

	if err := post.Delete(serv.db); err != nil {
		serv.SendInternalErr(w, err)
		return
	}

	w.Write(nil)
}
