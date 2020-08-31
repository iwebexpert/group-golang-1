package main

import (
	"context"
	"fmt"
	"lesson6/server"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	ctx := context.Background()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://root:1234@localhost:27017"))
	if err != nil {
		fmt.Println(err)
		return
	}

	defer client.Disconnect(ctx)
	db := client.Database("posts")
	srv, err := server.New(ctx, db)
	if err != nil {
		fmt.Println("Server start error", err)
		return
	}

	srv.Serve(":8080")
}
