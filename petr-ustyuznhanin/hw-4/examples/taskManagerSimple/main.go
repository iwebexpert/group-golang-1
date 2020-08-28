package main

import (
	"html/template"
	"log"
	"net/http"
)

// TaskList - список задач
type TaskList struct {
	Name string
	Description string
	List []Task
}

// Task - задача и ее статус
type Task struct {
	ID string
	Text string
	Complete bool
}

var tmpl = template.Must(template.New("MyTemplate").ParseFiles("tmpl.html"))

var simpleList = TaskList{
	Name: "Название листа",
	Description: "Описание листа с задачами",
	List: []Task{
		{"first", "Первая задача", false},
		{"second", "Вторая задача", false},
		{"thrid", "Третья задача", true},
	},
}

func main() {
	router := http.NewServeMux()

	router.HandleFunc("/", viewList)

	log.Fatal(http.ListenAndServe(":8080", router))
}

func viewList(w http.ResponseWriter, r *http.Request) {
	if err := tmpl.ExecuteTemplate(w, "list", simpleList); err != nil {
		log.Println(err)
	}
}