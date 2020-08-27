package models

import (
	"database/sql"
	"time"
)

type PostItem struct {
	ID     int    `jsosn: "id"`
	Header string `jsosn: "header"`
	Text   string `jsosn: "text"`
	Date   string `jsosn: "date"`
}

type PostItemSlice []PostItem

func GetAllPostItems(db *sql.DB) (PostItemSlice, error) {
	row, err := db.Query("SELECT ID, Header, Text, Date FROM posts")
	if err != nil {
		return nil, err
	}

	posts := make(PostItemSlice, 0, 10)
	for row.Next() {
		post := PostItem{}
		if err := row.Scan(&post.ID, &post.Header, &post.Text, &post.Date); err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	return posts, nil
}

func GetPostItem(db *sql.DB, id int) (PostItemSlice, error) {
	row, err := db.Query("SELECT ID, Header, Text, Date FROM posts WHERE ID = ?", id)
	if err != nil {
		return nil, err
	}

	posts := make(PostItemSlice, 0, 10)
	for row.Next() {
		post := PostItem{}
		if err := row.Scan(&post.ID, &post.Header, &post.Text, &post.Date); err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	return posts, nil
}

func (post *PostItem) Insert(db *sql.DB) error {
	_, err := db.Exec("INSERT INTO posts (Header, Text, Date) VALUES (?, ?, ?)", post.Header, post.Text, time.Now())
	return err
}
func (post *PostItem) Update(db *sql.DB) error {
	_, err := db.Exec("UPDATE posts SET Header=?, Text=?, Date=? WHERE ID = ?", post.Header, post.Text, time.Now(), post.ID)
	return err
}
func (post *PostItem) Delete(db *sql.DB) error {
	_, err := db.Exec("DELETE FROM posts WHERE ID = ?", post.ID)
	return err
}
