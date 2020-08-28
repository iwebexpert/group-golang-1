package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

// TaskList - список задач
type TaskList struct {
	ID int
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

func viewList(w http.ResponseWriter, r *http.Request) {
	if err := tmpl.ExecuteTemplate(w, "list", simpleList); err != nil {
		log.Println(err)
	}
}

var database *sql.DB

// GetAllList - получение всех списков задачами
func GetAllList() ([]TaskList, error)  {
	res := []TaskList{}
	rows, err := database.Query("select * from task_list_app.lists")
	if err != nil {
		return res, err
	}

	for rows.Next() {
		list := TaskList{}

		err := rows.Scan(&list.ID, &list.Name, &list.Description)
		if err != nil {
			log.Println(err)
			continue
		}
		res = append(res, list)
	}
	return res, nil
}

// GetList - получение списка по айди
func GetList(id string) (TaskList, error) {
	list := TaskList{}

	row := database.QueryRow(fmt.Sprintf("select * from task_list_app.lists where lists.id = %v", id))
	err := row.Scan(&list.ID, &list.Name, &list.Description)
	if err != nil {
			return list, err
	}

	rows := database.Query(fmt.Sprintf("select * from task_list_app.tasks where tasks.list_id = %v", id))
	if err != nil {
		return list, err
	}
	defer rows.Close()
}

func main() {
	//подключение к БД
	db, err := sql.Open("mysql", "root:ваш_пароль_к_БД@/task_list_app")
	if err != nil {
		log.Fatal(err)
	}

	database = db
	defer database.Close()

	router := http.NewServeMux()
	router.HandleFunc("/", viewList)
	log.Fatal(http.ListenAndServe(":8080", router))

}
