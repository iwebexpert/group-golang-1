package server

import (
	"context"
	"lesson6/models"
	"net/http"
	"os"
	"os/signal"

	"github.com/go-chi/chi"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
)

type Server struct {
	lg      *logrus.Logger
	Title   string
	Posts   models.Posts
	DBMongo *mongo.Database
	Ctx     context.Context
	router  *chi.Mux
}

func New(lg *logrus.Logger, ctx context.Context, mongodb *mongo.Database) (*Server, error) {
	posts, err := models.Get(ctx, mongodb)
	if err != nil {
		lg.WithError(err).Fatal("Models get err")
		return nil, err
	}

	srv := &Server{
		lg:      lg,
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

		for {
			err := http.ListenAndServe(port, srv.router)
			srv.lg.WithError(err).Fatal("Serve err")
		}
	}()
	<-stopChannel
	srv.lg.Warnln("Server stoped")
}
