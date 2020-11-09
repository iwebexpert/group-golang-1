package server

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"server/models"

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

func (serv *Server) SendError(w http.ResponseWriter, err error, code int, obj ...interface{}) {
	serv.lg.WithField("data", obj).WithError(err).Error("server error")
	w.WriteHeader(code)
	errModel := models.ErrorModel{
		Code:     code,
		Err:      err.Error(),
		Desc:     "server error",
		Internal: obj,
	}

	data, _ := json.Marshal(errModel)
	w.Write(data)
}

func (serv *Server) SendInternalError(w http.ResponseWriter, err error, obj ...interface{}) {
	serv.SendError(w, err, http.StatusInternalServerError, obj)
}
