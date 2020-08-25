package main

import (
	"html/template"
	"io/ioutil"
	"math"
	"net/http"
	"os"
)

type Server struct {
	Title string
	Tasks TaskItems
}

type TaskItems []TaskItem

type TaskItem struct {
	Text string
	Completed bool
	Labels []string
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
	percent := float64(tasks.TaskWithStatus(true))/float64(len(tasks))
	return math.Round(percent * 100)
}