package models

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Mongo - структура для получения ID
type Mongo struct {
	Id primitive.ObjectID `bson:"_id,omitempty"`
}

// Ctx - контекст
var Ctx = context.Background()

// Db - переменная для БД
var Db *mongo.Database

// GetCollectionName - для перегрузки метода
func (m *Mongo) GetCollectionName() string {
	panic("Метод не был перегружен GetCollectionName()")
	return ""
}

func init() {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://root:1@192.168.1.9:27017"))
	if err != nil {
		log.Fatal(err)
	}
	err = client.Connect(Ctx)
	if err != nil {
		log.Fatal(err)
	}
	Db = client.Database("my_blog")
}
