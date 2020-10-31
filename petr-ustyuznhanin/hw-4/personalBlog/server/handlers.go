package server

import (
	"encoding/json"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"server/models"

	"github.com/go-chi/chi"
	uuid "github.com/satori/go.uuid"
)

func (serv *Server) getTemplateHandler(w http.ResponseWriter, r *http.Request) {
	templateName := chi.URLParam(r, "template")

	if templateName == "" {
		templateName = serv.indexTemplate
	}

	file, err := os.Open(path.Join(serv.rootDir, serv.templatesDir, templateName))
	if err != nil {
		if err != os.ErrNotExist {
			w.WriteHeader(http.StatusNotFound)
		}
		serv.SendInternalError(w, err)
		return
	}

	data, err := ioutil.ReadAll(file)
	if err != nil {
		serv.SendInternalError(w, err)
		return
	}

	templ, err := template.New("Page").Parse(string(data))
	if err != nil {
		serv.SendInternalError(w, err)
		return
	}

	posts, err := models.GetAllPosts(serv.db)
	if err != nil {
		serv.SendInternalError(w, err)
		return
	}

	serv.Page.Posts = posts

	if err := templ.Execute(w, serv.Page); err != nil {
		serv.SendInternalError(w, err)
		return
	}
}

func (serv *Server) postPostHandler(w http.ResponseWriter, r *http.Request) {
	data, _ := ioutil.ReadAll(r.Body)

	post := models.PostItem{}
	_ = json.Unmarshal(data, &post)

	post.ID = uuid.NewV4().String()

	if err := post.Insert(serv.db); err != nil {
		serv.SendInternalError(w, err)
		return
	}

	data, _ = json.Marshal(post)
	w.Write(data)
}
