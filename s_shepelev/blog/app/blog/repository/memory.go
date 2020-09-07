package repository

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/Toringol/group-golang-1/tree/master/s_shepelev/blog/app/blog"
	"github.com/Toringol/group-golang-1/tree/master/s_shepelev/blog/app/model"
	"github.com/spf13/viper"
)

// Repository - Database methods with posts implemetation
type Repository struct {
	DB *mongo.Database
}

// NewBlogMemoryRepository - create connection and return new repository
func NewBlogMemoryRepository() blog.Repository {
	dsn := viper.GetString("database")

	clientOptions := options.Client().ApplyURI(dsn)
	ctx := context.TODO()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	db := client.Database("blog_app")

	return &Repository{
		DB: db,
	}
}

func (repo *Repository) collection() *mongo.Collection {
	return repo.DB.Collection("posts")
}

// ListPosts - return all posts in DB
func (repo *Repository) ListPosts() ([]*model.Post, error) {
	posts := []*model.Post{}

	ctx := context.TODO()

	cur, err := repo.collection().Find(ctx, bson.M{})
	if err != nil {
		return posts, err
	}

	for cur.Next(ctx) {
		record := &model.Post{}

		err := cur.Decode(record)
		if err != nil {
			return posts, err
		}

		posts = append(posts, record)
	}

	if err := cur.Err(); err != nil {
		return posts, err
	}

	cur.Close(ctx)

	return posts, nil
}

// SelectPostByID - return post by id
func (repo *Repository) SelectPostByID(id string) (*model.Post, error) {
	record := &model.Post{}

	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	err = repo.collection().FindOne(context.TODO(), bson.M{"_id": oid}).Decode(record)
	if err != nil {
		return nil, err
	}

	return record, nil
}

// CreatePost - create new record in DB posts
func (repo *Repository) CreatePost(post *model.Post) error {
	_, err := repo.collection().InsertOne(context.TODO(), post)

	return err
}

// UpdatePost - update post`s data in DataBase
func (repo *Repository) UpdatePost(post *model.Post) error {
	oid, err := primitive.ObjectIDFromHex(post.ID)
	if err != nil {
		return err
	}

	filter := bson.M{"_id": oid}

	update := bson.D{
		{"$set", bson.D{
			{"title", post.Title},
			{"description", post.Description},
			{"author", post.Author},
		}},
	}

	result := repo.collection().FindOneAndUpdate(context.TODO(), filter, update)

	return result.Err()
}

// DeletePost - delete post`s record in DataBase
func (repo *Repository) DeletePost(id string) error {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	_, err = repo.collection().DeleteOne(context.TODO(), bson.M{"_id": oid})

	return err
}
