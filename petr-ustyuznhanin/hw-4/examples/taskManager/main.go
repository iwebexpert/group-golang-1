package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
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

// обработчик роута для конкретного листа
func viewList(w http.ResponseWriter, r *http.Request) {
	list, err := GetList(r.URL.Query().Get("id"))
	if err != nil {
		log.Fatal(err)
		w.WriteHeader(404)
		return
	}

	if err := tmpl.ExecuteTemplate(w, "list", list); err != nil {
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
	defer rows.Close()

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

	rows, err := database.Query(fmt.Sprintf("select * from task_list_app.tasks where tasks.list_id = %v", id))
	if err != nil {
		return list, err
	}
	defer rows.Close()

	for rows.Next() {
		task := Task{}

		err := rows.Scan(&task.ID, new(int), &task.Text, &task.Complete)
		if err != nil {
			log.Println(err)
			continue
		}

		list.List = append(list.List, task)
	}
	return list,nil
}

// роут для показа всех списков задач
func viewLists(w http.ResponseWriter, r *http.Request){
	lists, err := GetAllList()
	if err != nil {
		log.Fatal(err)
		w.WriteHeader(500)
		return
	}

	if err := tmpl.ExecuteTemplate(w, "alllists", lists); err != nil {
		log.Println(err)
	}
}

func main() {
	//подключение к БД
	db, err := sql.Open("mysql", "root:1234@/task_list_app")
	if err != nil {
		log.Fatal(err)
	}

	database = db
	defer database.Close()

	router := http.NewServeMux()
	router.HandleFunc("/", viewList)
	router.HandleFunc("/", viewLists)
	log.Fatal(http.ListenAndServe(":8080", router))

}
