package main

import (
	"context"
	"lesson6/server"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	lg := NewLogger()

	ctx := context.Background()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://root:1234@localhost:27017"))
	if err != nil {
		lg.WithError(err).Fatal("Не удалось соединиться с БД")
		return
	}

	defer client.Disconnect(ctx)
	db := client.Database("posts")
	srv, err := server.New(lg, ctx, db)
	if err != nil {
		lg.WithError(err).Fatal("Server start err.")
		return
	}
	lg.Debug("Server is started ...")

	srv.Serve(":8080")
}
