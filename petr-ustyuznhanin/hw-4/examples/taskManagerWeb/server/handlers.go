package server

import (
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"server/models"

	"github.com/go-chi/chi"
)

func (serv *Server) getTemplateHandler(w http.ResponseWriter, r *http.Request) {
	templateName := chi.URLParam(r, "template")

	if templateName == "" {
		templateName = serv.indexTemplate
	}

	file, err := os.Open(path.Join(serv.rootDir, serv.templatesDir, templateName))
	if err != os.ErrNotExist {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	serv.SendInternalErr(w, err)
	return

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

	tasks, err := models.GetAllTasks(serv.db)
	if err != nil {
		serv.SendInternalErr(w, err)
		return
	}

	serv.Page.Tasks = tasks

	if err := templ.Execute(w, serv.Page); err != nil {
		serv.SendInternalErr(w, err)
		return
	}

}
