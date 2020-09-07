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
		panic(fmt.Sprintf("Ошибка чтения файла конфигурации %s", err))
	}

	db, err := sql.Open("pgx", fmt.Sprintf("user=%s password=%s host=%s port=%s database=%s",
		viper.GetString("database.user"), viper.GetString("database.password"), viper.GetString("database.host"),
		viper.GetString("database.port"), viper.GetString("database.database")))
	if err != nil {
		panic(fmt.Sprintf("Не удалось соединиться с БД: %s", err))
	}
	defer db.Close()

	srv, err := server.New("Учебный блог ice65537", db)
	if err != nil {
		panic(fmt.Sprintf("Ошибка создания сервера: %s", err))
	}

	srv.Serve(":" + viper.GetString("server.listen_port"))
}
