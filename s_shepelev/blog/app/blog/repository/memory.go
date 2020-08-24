package repository

import (
	"database/sql"
	"log"

	"github.com/Toringol/group-golang-1/tree/master/s_shepelev/blog/app/blog"
	"github.com/Toringol/group-golang-1/tree/master/s_shepelev/blog/app/model"
	"github.com/spf13/viper"
)

// Repository - Database methods with posts implemetation
type Repository struct {
	DB *sql.DB
}

// NewBlogMemoryRepository - create connection and return new repository
func NewBlogMemoryRepository() blog.Repository {
	dsn := viper.GetString("database")
	dsn += "&charset=utf8"
	dsn += "&interpolateParams=true"

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Println(err)
	}
	db.SetMaxOpenConns(10)

	err = db.Ping()
	if err != nil {
		log.Println("Error while Ping")
	}

	return &Repository{
		DB: db,
	}
}

// ListPosts - return all posts in DB
func (repo *Repository) ListPosts() ([]*model.Post, error) {
	posts := []*model.Post{}
	rows, err := repo.DB.Query("SELECT id, title, description, author FROM posts")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		record := &model.Post{}
		err = rows.Scan(&record.ID, &record.Title, &record.Description, &record.Author)
		if err != nil {
			return nil, err
		}
		posts = append(posts, record)
	}
	return posts, nil
}

// SelectPostByID - return post by id
func (repo *Repository) SelectPostByID(id int64) (*model.Post, error) {
	record := &model.Post{}
	err := repo.DB.
		QueryRow("SELECT id, title, description, author FROM posts WHERE id = ?", id).
		Scan(&record.ID, &record.Title, &record.Description, &record.Author)
	if err != nil {
		return nil, err
	}
	return record, nil
}

// CreatePost - create new record in DB posts
func (repo *Repository) CreatePost(post *model.Post) (int64, error) {
	result, err := repo.DB.Exec(
		"INSERT INTO posts (`title`, `description`, `author`) VALUES (?, ?, ?)",
		post.Title,
		post.Description,
		post.Author,
	)
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}

// UpdatePost - update post`s data in DataBase
func (repo *Repository) UpdatePost(post *model.Post) (int64, error) {
	result, err := repo.DB.Exec(
		"UPDATE posts SET"+
			"`title` = ?"+
			",`description` = ?"+
			",`author` = ? "+
			"WHERE id = ?",
		post.Title,
		post.Description,
		post.Author,
		post.ID,
	)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

// DeletePost - delete post`s record in DataBase
func (repo *Repository) DeletePost(id int64) (int64, error) {
	result, err := repo.DB.Exec(
		"DELETE FROM posts WHERE id = ?",
		id,
	)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}
