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

//создание задачи
func (serv *Server) postTaskHandler(w http.ResponseWriter, r *http.Request) {
	data, _ := ioutil.ReadAll(r.Body)

	task := models.TaskItem{}
	_ = json.Unmarshal(data, &task)

	task.ID = uuid.NewV4().String()

	if err := task.Insert(serv.db); err != nil {
		serv.SendInternalErr(w, err)
		return
	}

	data, _ = json.Marshal(task)
	w.Write(data)
}

//редактирование задачи
func (serv *Server) putTaskHandler(w http.ResponseWriter, r *http.Request) {
	taskID := chi.URLParam(r, "id")

	data, _ := ioutil.ReadAll(r.Body)

	task := models.TaskItem{}
	_ = json.Unmarshal(data, &task)

	task.ID = taskID //TODO: по хорошему нужна проверка

	if err := task.Update(serv.db); err != nil {
		serv.SendInternalErr(w, err)
		return
	}

	data, _ = json.Marshal(task)
	w.Write(data)
}

//удаление задачи
func (serv *Server) deleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	taskID := chi.URLParam(r, "id")

	task := models.TaskItem{ID: taskID}

	if err := task.Delete(serv.db); err != nil {
		serv.SendInternalErr(w, err)
		return
	}

	w.Write(nil)
}
