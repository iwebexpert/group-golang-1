package server

import (
	"database/sql"
	"net/http"

	"github.com/go-chi/chi"

	"github.com/sirupsen/logrus"
)

type Server struct {
	lg            *logrus.Logger
	db            *sql.DB
	rootDir       string
	templatesDir  string
	indexTemplate string
	Page          models.Page
}

func New(lg *logrus.Logger, rootDir string, db *sql.DB) *Server {
	return &Server{
		lg:            lg,
		db:            db,
		rootDir:       rootDir,
		templatesDir:  "/templates",
		indexTemplate: "index.html",
		Page: models.Page{
			Posts: models.PostItemSlice{
				{ID: "0", Title: "titleTest", Text: "textTest"},
			},
		},
	}
}

func (serv *Server) Start(addr string) error {
	r := chi.NewRouter()
	serv.bindRoutes(r)
	serv.lg.Debug("server is started...")
	return http.ListenAndServe(addr, r)
}
