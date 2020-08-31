package models

import "database/sql"

type Article struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	Text  string `json:"text"`
	Tags  string `json:"tags"`
}

type ArticleItemSlice []Article

func GetAllArticleItems(db *sql.DB) (ArticleItemSlice, error) {
	row, err := db.Query("SELECT ID, Title, Text, Tags FROM articles")
	if err != nil {
		return nil, err
	}

	articles := make(ArticleItemSlice, 0, 10)
	for row.Next() {
		article := Article{}
		if err := row.Scan(&article.ID, &article.Title, &article.Text, &article.Tags); err != nil {
			return nil, err
		}

		articles = append(articles, article)
	}
	return articles, nil
}

func (a *Article) Insert(db *sql.DB) error {
	_, err := db.Exec("INSERT INTO articles (ID, Title, Text, Tags) VALUES (?, ?, ?, ?)",
		a.ID, a.Title, a.Text, a.Tags)

	return err
}

/*
func (a *Article) Update(db *sql.DB) error {
	_, err := db.Exec("UPDATE articles SET Title = ?, Text = ?, Tags = ? WHERE ID = ?",
		a.Title, a.Text, a.Tags, a.ID)

	return err
}
*/

func (a *Article) Delete(db *sql.DB) error {
	_, err := db.Exec("DELETE FROM articles WHERE ID = ?",
		a.ID)

	return err
}
