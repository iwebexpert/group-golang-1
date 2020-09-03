package main

import (
	"encoding/json"
	"html/template"
	"io/ioutil"
	"math"
	"net/http"
	"os"
	"strconv"

	"github.com/go-chi/chi"
)

type Server struct {
	Title string
	Tasks TaskItems
}

type TaskItems []TaskItem

type TaskItem struct {
	Text      string
	Completed bool
	Labels    []string
}

func (server *Server) GetIndexHandler(w http.ResponseWriter, r *http.Request) {
	file, err := os.Open("./www/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	templateMain := template.Must(template.New("managerList").Parse(string(data)))
	templateMain.ExecuteTemplate(w, "managerList", server)
}

func (server *Server) PostTaskStatusHandler(w http.ResponseWriter, r *http.Request) {
	taskIdString := chi.URLParam(r, "taskID")
	taskStatusString := chi.URLParam(r, "status")

	taskID, err := strconv.ParseInt(taskIdString, 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	taskStatus, err := strconv.ParseBool(taskStatusString)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	server.Tasks[taskID].Completed = taskStatus
	data, err := json.Marshal(server.Tasks[taskID])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Write(data)
}

func (tasks TaskItems) TaskWithStatus(completed bool) int {
	total := 0
	for _, task := range tasks {
		if task.Completed == completed {
			total++
		}
	}
	return total
}

func (tasks TaskItems) CompletedTasksPercent() float64 {
	percent := float64(tasks.TaskWithStatus(true)) / float64(len(tasks))
	return math.Round(percent * 100)
}

func main() {
	route := chi.NewRouter()
	server := Server{
		Title: "The Task Manager",
		Tasks: TaskItems{
			{Text: "Изучить Го", Completed: false, Labels: []string{"Go", "Lessons"}},
			{Text: "Create web-server", Completed: true, Labels: []string{"Go", "Server"}},
			{Text: "Xyz", Completed: false},
		},
	}

	route.Route("/", func(r chi.Router) {
		r.Get("/", server.GetIndexHandler)
		r.Post("/{taskID}/{status}", server.PostTaskStatusHandler)
	})
	http.ListenAndServe(":8080", route)
}
