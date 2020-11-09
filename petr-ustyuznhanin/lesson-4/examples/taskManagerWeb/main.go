package main

import (
	"database/sql"
	"os"
	"os/signal"
	"server/server"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	lg := NewLogger()

	db, err := sql.Open("mysql", "taskmanager:1234@/taskmanager")
	if err != nil {
		lg.WithError(err).Fatal("Не удалось соединиться с БД")
	}
	defer db.Close()

	serv := server.New(lg, "./www", db)

	go func() {
		err := serv.Start("localhost:8080")
		if err != nil {
			lg.WithError(err).Fatal("Не удалось запустить сервер")
		}
	}()

	stopSignal := make(chan os.Signal)
	signal.Notify(stopSignal, os.Interrupt, os.Kill)
	<-stopSignal // ожидаем
}
