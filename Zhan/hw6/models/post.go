package models

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// Post - структура связанная с таблицей в БД
type Post struct {
	Mongo  `inline`
	Header string `bson:"header"`
	Text   string `bson:"text"`
}

// GetCollectionName - получаем имя коллекции для дальнейшего использования
func (p *Post) GetCollectionName() string {
	return "posts"
}

// Insert - добавление нового поста
func (p *Post) Insert(ctx context.Context, db *mongo.Database) (*Post, error) {
	collection := db.Collection(p.GetCollectionName())
	_, err := collection.InsertOne(ctx, p)
	if err != nil {
		return nil, err
	}
	return p, nil
}

// Update - обновление поста
func (p *Post) Update(ctx context.Context, db *mongo.Database) (*Post, error) {
	collection := db.Collection(p.GetCollectionName())
	_, err := collection.ReplaceOne(ctx, bson.M{"_id": p.Id}, p)
	return p, err
}

// Delete - удаление поста
func (p *Post) Delete(ctx context.Context, db *mongo.Database) (*Post, error) {
	collection := db.Collection(p.GetCollectionName())
	_, err := collection.DeleteOne(ctx, bson.M{"_id": p.Id})
	return p, err
}

// GetOne - получение одного поста
func GetOne(ctx context.Context, db *mongo.Database, id string) (*Post, error) {
	p := Post{}
	collection := db.Collection(p.GetCollectionName())

	docId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	result := collection.FindOne(ctx, bson.M{"_id": docId})
	if err := result.Decode(&p); err != nil {
		return nil, err
	}
	return &p, nil
}

// GetAll - получение всех постов
func GetAll(ctx context.Context, db *mongo.Database) ([]Post, error) {
	p := Post{}
	posts := []Post{}
	collection := db.Collection(p.GetCollectionName())

	result, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	for result.Next(ctx) {
		err := result.Decode(&p)
		if err != nil {
			return nil, err
		}
		posts = append(posts, p)
	}
	if err := result.Err(); err != nil {
		return nil, err
	}
	return posts, nil
}
