package server

import (
	"blog/models"
	"fmt"
	"net/http"
	"os"
	"os/signal"

	"github.com/go-chi/chi"
	"github.com/jinzhu/gorm"
)

//BlogServer -
type BlogServer struct {
	Title     string
	Posts     models.BlogPostArray
	DBlinkORM *gorm.DB
	router    *chi.Mux
}

//New --
func New(title string, gormdb *gorm.DB) (*BlogServer, error) {
	posts, err := models.Retrieve(gormdb)
	if err != nil {
		return nil, err
	}

	srv := &BlogServer{
		Title:     title,
		DBlinkORM: gormdb,
		Posts:     *posts,
		router:    chi.NewRouter(),
	}
	srv.defineRoutes()
	return srv, nil
}

//Serve --
func (srv *BlogServer) Serve(port string) {
	stopChannel := make(chan os.Signal)
	signal.Notify(stopChannel, os.Kill, os.Interrupt)

	go func() {
		fmt.Println("Server start")
		for {
			err := http.ListenAndServe(port, srv.router)
			fmt.Println(err)
			fmt.Println("Не упал")
		}
	}()

	//Ждем сигнала от OS
	<-stopChannel
	fmt.Println("Server stop")
}
