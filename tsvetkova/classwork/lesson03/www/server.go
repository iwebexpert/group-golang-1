package main

import (
	"net/http"
	"os"
	"io/ioutil"
	"html/template"
	"math"

	"github.com/go-chi/chi"
)

type Server struct {
	Title 	string
	Tasks 	TaskItems
}



func (server *Server) GetIndexHandler(w http.ResponseWriter, r *http.Request) {
	file, _ := os.Open("./static/index.html")
	data, _ := ioutil.ReadAll(file)

	templateMain := template.Must(template.New("managerList").Parse(string(data)))
	templateMain.ExecuteTemplate(w, "managerList", server)
}

type TaskItem struct {
	Text 		string
	Completed 	bool
	Labels 		[]string
}
type TaskItems []TaskItem

func (tasks TaskItems) TasksWithStatus(completed bool) int {
	total := 0
	for _, task := range tasks {
		if (task.Completed == completed) {
			total++
		}
	}
	return total
}

func (tasks TaskItems) CompletedTasksPercent() float64 {
	percent := float64(tasks.TasksWithStatus(true)) / float64(len(tasks))
	return math.Round(percent * 100 )
}

func main() {
	router := chi.NewRouter()

	server := Server{
		Title: "Task manager",
		Tasks: TaskItems{
			{"Изучить Go", false, []string{"go", "lessons", "обучение"}},
			{"Создать веб-сервер", true, []string{"go", "lessons", "обучение", "сервер"}},
			{"Еще что-то", false, []string{}},
		},
	}

	router.Route("/", func(r chi.Router) {
		r.Get("/", server.GetIndexHandler)
	})

	http.ListenAndServe(":8080", router)
}