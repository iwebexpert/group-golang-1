package server

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"server/models"

	"github.com/go-chi/chi"
	"github.com/sirupsen/logrus"
)

// plus minus standart template for server

//Server ...
type Server struct {
	lg            *logrus.Logger
	db            *sql.DB
	rootDir       string
	templatesDir  string
	indexTemplate string
	Page          models.Page
}

//New create new server
func New(lg *logrus.Logger, rootDir string, db *sql.DB) *Server {
	return &Server{
		lg:            lg,
		db:            db,
		rootDir:       rootDir,
		templatesDir:  "/templates",
		indexTemplate: "index.html",
		Page: models.Page{
			Tasks: models.TaskItemSlice{
				{ID: "0", Text: "1234", Completed: false},
			},
		},
	}
}

//Start ...
func (serv *Server) Start(addr string) error {
	r := chi.NewRouter()
	serv.bindRoutes(r)
	serv.lg.Debug("server is started...")
	return http.ListenAndServe(addr, r)
}

//SendErr обрабатывает общие маршруты для сервера, есили что-то не обрабатывается, то возвращается ошибка
func (serv *Server) SendErr(w http.ResponseWriter, err error, code int, obj ...interface{}) {
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

//SendInternalErr ...
func (serv *Server) SendInternalErr(w http.ResponseWriter, err error, obj ...interface{}) {
	serv.SendErr(w, err, http.StatusInternalServerError, obj)
}