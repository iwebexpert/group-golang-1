package server

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"server/models"
	"text/template"

	"github.com/go-chi/chi"
	uuid "github.com/satori/go.uuid"
)

func (s *Server) getTemplateHandler(w http.ResponseWriter, r *http.Request) {
	templateName := chi.URLParam(r, "template")

	if templateName == "" {
		templateName = s.indexTemplate
	}

	file, err := os.Open(path.Join(s.rootDir, s.templatesDir, templateName))
	if err != nil {
		if err == os.ErrNotExist {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		s.SendInternalErr(w, err)
		return
	}

	data, err := ioutil.ReadAll(file)
	if err != nil {
		s.SendInternalErr(w, err)
		return
	}

	templ, err := template.New("Page").Parse(string(data))
	if err != nil {
		s.SendInternalErr(w, err)
		return
	}

	articles, err := models.GetAllArticleItems(s.db)
	if err != nil {
		s.SendInternalErr(w, err)
		return
	}

	s.Page.Articles = articles

	if err := templ.Execute(w, s.Page); err != nil {
		s.SendInternalErr(w, err)
		return
	}
}

func (s *Server) postArticleHandler(w http.ResponseWriter, r *http.Request) {
	data, _ := ioutil.ReadAll(r.Body)

	article := models.Article{}
	_ = json.Unmarshal(data, &article)

	article.ID = uuid.NewV4().String()

	if err := article.Insert(s.db); err != nil {
		s.SendInternalErr(w, err)
		return
	}

	data, _ = json.Marshal(article)
	w.Write(data)
}

/*
func (s *Server) putArticleHandler(w http.ResponseWriter, r *http.Request) {
	articleID := chi.URLParam(r, "id")

	data, _ := ioutil.ReadAll(r.Body)

	article := models.Article{}
	_ = json.Unmarshal(data, &article)

	article.ID = articleID

	if err := article.Update(s.db); err != nil {
		s.SendInternalErr(w, err)
		return
	}

	data, _ = json.Marshal(article)
	w.Write(data)
}
*/
func (s *Server) deleteArticleHandler(w http.ResponseWriter, r *http.Request) {
	articleID := chi.URLParam(r, "id")

	article := models.Article{ID: articleID}

	if err := article.Delete(s.db); err != nil {
		s.SendInternalErr(w, err)
		return
	}

	w.Write(nil)
}
