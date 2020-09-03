package main

import (
	"blog/config"
	"blog/server"
	"database/sql"
	"fmt"

	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/spf13/viper"
)

func main() {
	if err := config.Parse(); err != nil {
		fmt.Println("Ошибка чтения файла конфигурации", err)
		return
	}

	db, err := sql.Open("pgx", fmt.Sprintf("user=%s password=%s host=%s port=%s database=%s",
		viper.GetString("user"), viper.GetString("password"), viper.GetString("host"),
		viper.GetString("port"), viper.GetString("database")))
	if err != nil {
		fmt.Println("Не удалось соединиться с БД", err)
		return
	}
	defer db.Close()

	srv, err := server.New("Учебный блог ice65537", db)
	if err != nil {
		fmt.Println("Ошибка создания сервера", err)
		return
	}

	srv.Serve(":" + viper.GetString("listenPort"))
}
