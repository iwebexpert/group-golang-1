package models

import (
	"context"
	"fmt"
	"html/template"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

//Mongo -
type Mongo struct {
	ID primitive.ObjectID `json:"id"            bson:"_id,omitempty" `
}

//BlogPost -
type BlogPost struct {
	Mongo      `inline`
	About      string        `json:"about"      bson:"about"`
	Text       template.HTML `json:"text"       bson:"text"`
	Labels     []string      `json:"labels"     bson:"labels"`
	PublicDate time.Time     `json:"publicated" bson:"publicated"`
}

//BlogPostArray -
type BlogPostArray map[string]BlogPost

//Retrieve -
func Retrieve(ctx context.Context, db *mongo.Database) (*BlogPostArray, error) {
	cur, err := db.Collection("posts").Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	posts := make(BlogPostArray, 10)

	defer cur.Close(ctx)
	for cur.Next(ctx) {
		post := BlogPost{}
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

//NewBlogPost --
func (posts *BlogPostArray) NewBlogPost(ctx context.Context, db *mongo.Database, about string, text template.HTML, labels []string) (string, error) {
	tmp := BlogPost{
		About:      about,
		Text:       text,
		Labels:     labels,
		PublicDate: time.Now(),
	}

	rslt, err := db.Collection("posts").InsertOne(ctx, tmp)
	fmt.Println(rslt)
	if err != nil {
		return "", err
	}
	id := rslt.InsertedID.(primitive.ObjectID).Hex()
	(*posts)[id] = tmp

	return id, nil
}

//UpdateBlogPost --
func (posts *BlogPostArray) UpdateBlogPost(ctx context.Context, db *mongo.Database, id string, about string, text template.HTML, labels []string) error {
	tmp := BlogPost{
		About:      about,
		Text:       text,
		Labels:     labels,
		PublicDate: time.Now(),
	}

	objid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	tmp.ID = objid

	rslt, err := db.Collection("posts").ReplaceOne(ctx, bson.M{"_id": objid}, tmp)
	fmt.Println(rslt)
	if err != nil {
		return err
	}
	(*posts)[id] = tmp
	return nil
}

//DeleteBlogPost --
func (posts *BlogPostArray) DeleteBlogPost(ctx context.Context, db *mongo.Database, id string) error {
	objid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	rslt, err := db.Collection("posts").DeleteOne(ctx, bson.M{"_id": objid})
	fmt.Println(rslt)
	if err != nil {
		return err
	}
	delete((*posts), id)
	return nil
}
