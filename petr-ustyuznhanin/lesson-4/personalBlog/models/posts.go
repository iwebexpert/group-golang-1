package models

import "database/sql"

// PostItem ...
type PostItem struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	Text  string `json:"text"`
}

// PostItemSlice ...
type PostItemSlice []PostItem

func GetAllPosts(db *sql.DB) (PostItemSlice, error) {
	row, err := db.Query("SELECT ID, Title, Text FROM PostItems")
	if err != nil {
		return nil, err
	}

	posts := make(PostItemSlice, 0, 10)
	for row.Next() {
		post := PostItem{}
		if err := row.Scan(&post.ID, &post.Title, &post.Text); err != nil {
			return nil, err
		}

		posts = append(posts, post)

	}
	return posts, nil
}

func (post *PostItem) Insert(db *sql.DB) error {
	_, err := db.Exec("INSERT INTO PostItems (ID, Title, Text) VALUES (?, ?, ?)",
		post.ID, post.Title, post.Text)
	return err
}

func (post *PostItem) Update(db *sql.DB) error {
	_, err := db.Exec("UPDATE PostItems SET Title = ?, Text = ? WHERE ID = ?",
		post.Title, post.Text, post.ID)
	return err
}

func (post *PostItem) Delete(db *sql.DB) error {
	_, err := db.Exec("DELETE FROM PostItems WHERE ID = ?", post.ID)
	return err
}
