package models

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Mongo struct {
	ID primitive.ObjectID `json:"id"            bson:"_id,omitempty" `
}

type Post struct {
	Mongo  `inline`
	Header string `json:"Header"      bson:"Header"`
	Text   string `json:"Text"       bson:"Text"`
	Date   string `json:"Date" bson:"Date"`
}
type Posts map[string]Post

//Get Posts
func Get(ctx context.Context, db *mongo.Database) (*Posts, error) {
	cur, err := db.Collection("posts").Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	posts := make(Posts, 10)

	defer cur.Close(ctx)
	for cur.Next(ctx) {
		post := Post{}
		err := cur.Decode(&post)
		if err != nil {
			return nil, err
		}
		posts[post.ID.Hex()] = post
	}
	if err := cur.Err(); err != nil {
		return nil, err
	}
	return &posts, nil
}
func (posts *Posts) NewPost(ctx context.Context, db *mongo.Database, header string, text string) (string, error) {
	t := time.Now()
	tmp := Post{
		Header: header,
		Text:   text,
		Date:   t.Format("2006-01-02 15:04:05"),
	}

	done, err := db.Collection("posts").InsertOne(ctx, tmp)
	fmt.Println(done)
	if err != nil {
		return "", err
	}
	id := done.InsertedID.(primitive.ObjectID).Hex()
	(*posts)[id] = tmp
	return id, nil
}

//UpdatePost
func (posts *Posts) UpdatePost(ctx context.Context, db *mongo.Database, id string, header string, text string) error {
	t := time.Now()
	tmp := Post{
		Header: header,
		Text:   text,
		Date:   t.Format("2006-01-02 15:04:05"),
	}

	newID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	tmp.ID = newID

	rslt, err := db.Collection("posts").ReplaceOne(ctx, bson.M{"_id": newID}, tmp)
	fmt.Println(rslt)
	if err != nil {
		return err
	}
	(*posts)[id] = tmp
	return nil
}

//DeletePost
func (posts *Posts) DeletePost(ctx context.Context, db *mongo.Database, id string) error {
	postID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	rslt, err := db.Collection("posts").DeleteOne(ctx, bson.M{"_id": postID})
	fmt.Println(rslt)
	if err != nil {
		return err
	}
	delete((*posts), id)
	return nil
}
