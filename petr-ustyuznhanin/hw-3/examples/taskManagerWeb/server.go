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

func (server *Server) GetIndexHandler(w http.ResponseWriter, r *http.Request) {
	file, _ := os.Open("./www/index.html")
	data, _ := ioutil.ReadAll(file)

	templateMain := template.Must(template.New("managerList").Parse(string(data)))
	templateMain.ExecuteTemplate(w, "managerList", server)
}

func (server *Server) PostTaskStatusHandler(w http.ResponseWriter, r *http.Request) {
	taskIdString := chi.URLParam(r, "taskID")
	taskStatusString := chi.URLParam(r, "status")

	taskID, _ := strconv.ParseInt(taskIdString, 10, 64)
	taskStatus, _ := strconv.ParseBool(taskStatusString)

	server.Tasks[taskID].Completed = taskStatus

	data, _ := json.Marshal(server.Tasks[taskID])
	w.Write(data)
}

type TaskItem struct {
	Text      string
	Completed bool
	Labels    []string
}

type TaskItems []TaskItem

func (tasks TaskItems) TasksWithStatus(completed bool) int {
	total := 0
	for _, task := range tasks {
		if task.Completed == completed {
			total++
		}
	}
	return total
}

func (tasks TaskItems) CompletedTasksPercent() float64 {
	percent := float64(tasks.TasksWithStatus(true)) / float64(len(tasks))
	return math.Round(percent * 100)
}

func main() {
	route := chi.NewRouter()

	server := Server{
		Title: "The Task Manager",
		Tasks: TaskItems{
			{Text: "Изучить Go", Completed: false, Labels: []string{"Go", "Lessons"}},
			{Text: "Создать веб-сервер", Completed: true, Labels: []string{"Go", "Server"}},
			{Text: "Еще что-то", Completed: false},
		},
	}

	route.Route("/", func(r chi.Router) {
		r.Get("/", server.GetIndexHandler)
		r.Post("/{taskID}/{status}", server.PostTaskStatusHandler)
	})

	http.ListenAndServe(":8080", route)
}