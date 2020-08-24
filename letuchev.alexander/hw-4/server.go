package main

import (
	"blog/models"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/go-chi/chi"
)

//BlogServer -
type BlogServer struct {
	Title string
	Posts models.BlogPostArray
}

//GetIndexHandler -
func (srv *BlogServer) GetIndexHandler(w http.ResponseWriter, r *http.Request) {
	file, _ := os.Open("./templates/index.html")
	data, _ := ioutil.ReadAll(file)

	templateMain := template.Must(template.New("blogIndex").Parse(string(data)))

	templateMain.ExecuteTemplate(w, "blogIndex", srv)
}

//GetBlogHandler -
func (srv *BlogServer) GetBlogHandler(w http.ResponseWriter, r *http.Request) {
	file, _ := os.Open("./templates/blog.html")
	data, _ := ioutil.ReadAll(file)

	fmt.Println("Подготовка шаблона Blog")
	tmp, err := template.New("Blog").Parse(string(data))
	if err != nil {
		fmt.Println("Ошибка!", err)
		w.Write([]byte(fmt.Sprint(err)))
		return
	}
	fmt.Println("Наполнение шаблона Blog данными")
	tmp.ExecuteTemplate(w, "Blog", srv)
}

func main() {
	stopChannel := make(chan os.Signal)
	signal.Notify(stopChannel, os.Kill, os.Interrupt)

	server := BlogServer{
		Title: "Учебный блог ice65537",
		Posts: models.BlogPostArray{
			{About: "Навального отравили 20.08.2020", Text: "<h4>Навального</h4> отравили 20.08.2020 неизвестным ядом", PublicDate: time.Date(2020, 7, 15, 22, 12, 11, 30, time.UTC), Labels: []string{"Россия", "Политика"}},
			{About: "В Белоруссии продолжаются протесты", PublicDate: time.Date(2020, 7, 22, 22, 12, 11, 30, time.UTC), Labels: []string{"Белоруссия", "Политика"}},
			{About: "В Росии едят блины с лопаты", PublicDate: time.Date(2020, 8, 02, 18, 15, 11, 30, time.UTC)},
		},
	}

	router := chi.NewRouter()
	router.Route("/", func(r chi.Router) {
		r.Get("/", func(w http.ResponseWriter, r *http.Request) { server.GetIndexHandler(w, r) })
		r.Get("/blog/", func(w http.ResponseWriter, r *http.Request) { server.GetBlogHandler(w, r) })
	})

	go func() {
		fmt.Println("Server start")
		for {
			err := http.ListenAndServe(":8080", router)
			fmt.Println(err)
			fmt.Println("Не упал")
		}
	}()

	//Ждем сигнала от OS
	<-stopChannel
	fmt.Println("Server stop")
}
