package main

import (
	"blog/server"
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {
	gormdb, err := gorm.Open("postgres", "user=blogdb_adm password=sys host=193.168.0.99 port=5432 database=blogdb")
	if err != nil {
		fmt.Println("Не удалось соединиться с БД", err)
		return
	}
	defer gormdb.Close()

	srv, err := server.New("Учебный блог ice65537", gormdb)
	if err != nil {
		fmt.Println("Ошибка создания сервера", err)
		return
	}

	srv.Serve(":8080")
}
