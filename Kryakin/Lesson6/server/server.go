package server

import (
	"context"
	"fmt"
	"lesson6/models"
	"net/http"
	"os"
	"os/signal"

	"github.com/go-chi/chi"
	"go.mongodb.org/mongo-driver/mongo"
)

type Server struct {
	Title   string
	Posts   models.Posts
	DBMongo *mongo.Database
	Ctx     context.Context
	router  *chi.Mux
}

func New(ctx context.Context, mongodb *mongo.Database) (*Server, error) {
	posts, err := models.Get(ctx, mongodb)
	if err != nil {
		return nil, err
	}

	srv := &Server{
		DBMongo: mongodb,
		Posts:   *posts,
		router:  chi.NewRouter(),
		Ctx:     ctx,
	}
	srv.defineRoutes()
	return srv, nil
}

func (srv *Server) Serve(port string) {
	stopChannel := make(chan os.Signal)
	signal.Notify(stopChannel, os.Kill, os.Interrupt)

	go func() {
		fmt.Println("Server start")
		for {
			err := http.ListenAndServe(port, srv.router)
			fmt.Println(err)
		}
	}()
	<-stopChannel
	fmt.Println("Server stop")
}
