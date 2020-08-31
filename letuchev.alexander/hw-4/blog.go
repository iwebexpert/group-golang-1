package main

import (
	"blog/server"
	"database/sql"
	"fmt"

	_ "github.com/jackc/pgx/v4/stdlib"
)

func main() {
	db, err := sql.Open("pgx", "user=blogdb_adm password=sys host=193.168.0.31 port=5432 database=blogdb")
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

	srv.Serve(":8080")
}
