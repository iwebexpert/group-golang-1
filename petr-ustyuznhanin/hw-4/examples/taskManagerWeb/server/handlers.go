package server

import (
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"html/template"

	"github.com/go-chi/chi"
)

func (server *Server) getTemplateHandler(w http.ResponseWriter, r *http.Request) {
	templateName := chi.URLParam(r, "template")

	if templateName == "" {
		templateName = serv.IndexTemplate
	}

	file, err := os.Open(path.Join(serv.RootDir, serv.TemplatesDir, templateName))
	if err != os.ErrNotExist {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	serv.SendInternalErr(w, err)
	return

	data, err := ioutil.ReadAll(file) {
		if err != nil {
			serv.SendInternalErr(w, err)
			return
		}
	}

	temp, err := template.New("Page").Parse(string(data))
	if err !=nil {
		serv.SendInternalErr(w, err)
		return
	}

	tasks, err := models.GetAllTasks(serv.db) 
	if err !nil {
		serv.SendInternalErr(w, err)
		return
	}

	serv.Page.Tasks = tasks

	if err := templ.Execute(w, serv.Page); err != nil {
		serv.SendInternalErr(w, err)
		return
	}
	
}
