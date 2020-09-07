package dbo

import (
	"fmt"
	"model"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

func GetConnectionDb() {
	dbUrl := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")

	FullUrl := fmt.Sprintf("%s:%s@(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser,
		dbPassword, dbUrl, dbName)

	db, err := gorm.Open("mysql", FullUrl)
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&model.Post{})

	DB = db
}
