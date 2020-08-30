package main

import (
	"blog/server"
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	ctx := context.Background()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://sys:sys@193.168.0.99:27017"))
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			fmt.Println(err)
		}
	}()
	db := client.Database("blog")

	srv, err := server.New(ctx, db, "Учебный блог ice65537")
	if err != nil {
		fmt.Println("Ошибка создания сервера", err)
		return
	}

	srv.Serve(":8080")
}
