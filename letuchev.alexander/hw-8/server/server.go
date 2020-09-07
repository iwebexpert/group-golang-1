package server

import (
	"blog/models"
	"database/sql"
	"io"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/go-chi/chi"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

//BlogServer -
type BlogServer struct {
	Title  string
	Posts  models.BlogPostArray
	DBlink *sql.DB
	router *chi.Mux
	log    *logrus.Logger
}

//New --
func New(title string, db *sql.DB) (*BlogServer, error) {
	var mwarr []io.Writer

	lg := logrus.New()
	if viper.GetBool("log.file_enable") {
		file, err := os.OpenFile(viper.GetString("log.filepath")+time.Now().Format("blog20060102-150405")+".log",
			os.O_RDWR|os.O_CREATE|os.O_APPEND, 0664)
		if err != nil {
			return nil, err
		}
		mwarr = append(mwarr, file)
	}
	if viper.GetBool("log.stdout_enable") {
		mwarr = append(mwarr, os.Stdout)
	}
	/* не поддерживается в windows
	if viper.GetBool("log.syslog_enable") {
		syslogW, err := syslog.New(syslog.LOG_DEBUG, "go_server_app")
		if err != nil {
			return nil, err
		}
		mwarr = append(mwarr, syslogW)
	}*/
	lg.SetOutput(io.MultiWriter(mwarr...))

	level, err := logrus.ParseLevel(viper.GetString("log.level"))
	if err != nil {
		lg.Error(err)
		lg.SetLevel(logrus.InfoLevel)
	}
	lg.SetLevel(level)

	posts, err := models.Retrieve(db, lg)
	if err != nil {
		lg.Error(err)
		return nil, err
	}

	srv := &BlogServer{
		Title:  title,
		DBlink: db,
		Posts:  *posts,
		router: chi.NewRouter(),
		log:    lg,
	}
	srv.defineRoutes()
	return srv, nil
}

//Serve --
func (srv *BlogServer) Serve(port string) {
	stopChannel := make(chan os.Signal)
	signal.Notify(stopChannel, os.Kill, os.Interrupt)

	go func() {
		srv.log.Debug("Server start")
		for {
			err := http.ListenAndServe(port, srv.router)
			srv.log.Error(err)
			srv.log.Debug("Не упал")
		}
	}()

	//Ждем сигнала от OS
	<-stopChannel
	srv.log.Debug("Server stop")
}
