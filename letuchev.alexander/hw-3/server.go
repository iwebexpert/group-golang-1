package main

import (
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
	About string
	Posts BlogPostArray
}

//GetIndexHandler -
func (srv *BlogServer) GetIndexHandler(w http.ResponseWriter, r *http.Request) {
	file, _ := os.Open("./templates/index.html")
	data, _ := ioutil.ReadAll(file)

	templateMain := template.Must(template.New("blogIndex").Parse(string(data)))
	templateMain.ExecuteTemplate(w, "blogIndex", srv)
}

//GetUniHandler -
func (srv *BlogServer) GetUniHandler(w http.ResponseWriter, r *http.Request, tmplt string) {
	file, _ := os.Open("./templates/" + tmplt + ".html")
	data, _ := ioutil.ReadAll(file)

	templateMain := template.Must(template.New(tmplt).Parse(string(data)))
	templateMain.ExecuteTemplate(w, tmplt, srv)
}

//BlogPost -
type BlogPost struct {
	About      string
	Text       template.HTML
	Labels     []string
	PublicDate time.Time
}

//BlogPostArray -
type BlogPostArray []BlogPost

func main() {
	stopChannel := make(chan os.Signal)
	signal.Notify(stopChannel, os.Kill, os.Interrupt)

	server := BlogServer{
		Title: "Учебный блог ice65537",
		About: "Домашнее задание №3. Шаблонизация",
		Posts: BlogPostArray{
			{About: "Навального отравили 20.08.2020", Text: "Навального отравили 20.08.2020 неизвестным ядом", PublicDate: time.Date(2020, 7, 15, 22, 12, 11, 30, time.UTC), Labels: []string{"Россия", "Политика"}},
			{About: "В Белоруссии продолжаются протесты", PublicDate: time.Date(2020, 7, 22, 22, 12, 11, 30, time.UTC), Labels: []string{"Белоруссия", "Политика"}},
			{About: "В Росии едят блины с лопаты", PublicDate: time.Date(2020, 8, 02, 18, 15, 11, 30, time.UTC)},
		},
	}

	router := chi.NewRouter()
	router.Route("/", func(r chi.Router) {
		r.Get("/", func(w http.ResponseWriter, r *http.Request) { server.GetUniHandler(w, r, "index") })
		r.Get("/blog/", func(w http.ResponseWriter, r *http.Request) { server.GetUniHandler(w, r, "blog") })
		r.Get("/blogPost/", func(w http.ResponseWriter, r *http.Request) { server.GetUniHandler(w, r, "blog") })
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
