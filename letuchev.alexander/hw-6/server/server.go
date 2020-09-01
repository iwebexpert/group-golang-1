package server

import (
	"blog/models"
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"

	"github.com/go-chi/chi"
	"go.mongodb.org/mongo-driver/mongo"
)

//BlogServer -
type BlogServer struct {
	Title   string
	Posts   models.BlogPostArray
	DBMongo *mongo.Database
	Ctx     context.Context
	router  *chi.Mux
}

//New -- Создание нового экземпляра сервера (требуется активное подключение к БД)
func New(ctx context.Context, mongodb *mongo.Database, title string) (*BlogServer, error) {
	posts, err := models.Retrieve(ctx, mongodb)
	if err != nil {
		return nil, err
	}

	srv := &BlogServer{
		Title:   title,
		DBMongo: mongodb,
		Posts:   *posts,
		router:  chi.NewRouter(),
		Ctx:     ctx,
	}
	srv.defineRoutes()
	return srv, nil
}

//Serve -- Сервер работать
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
