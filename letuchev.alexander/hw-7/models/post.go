package models

import (
	"database/sql"
	"fmt"
	"html/template"
	"strings"
	"time"
)

//BlogPost -
type BlogPost struct {
	ID         int           `json:"id"`
	About      string        `json:"about"`
	Text       template.HTML `json:"text"`
	Labels     []string      `json:"labels"`
	PublicDate time.Time     `json:"publicated"`
}

//BlogPostArray -
type BlogPostArray map[int]BlogPost

//Retrieve -
func Retrieve(db *sql.DB) (*BlogPostArray, error) {
	row, err := db.Query("SELECT nID, coalesce(sAbout,''), coalesce(sText,''), coalesce(sLabels,''), dtPublic FROM blog.t_post")
	if err != nil {
		return nil, err
	}

	posts := make(BlogPostArray, 10)
	for row.Next() {
		post := BlogPost{}
		labels := ""
		if err = row.Scan(&post.ID, &post.About, &post.Text, &labels, &post.PublicDate); err != nil {
			return nil, err
		}
		post.Labels = strings.Split(labels, ",")
		posts[post.ID] = post
	}
	return &posts, nil
}

//NewBlogPost --
func (posts *BlogPostArray) NewBlogPost(db *sql.DB, about string, text template.HTML, labels []string) (int, error) {
	row, err := db.Query("INSERT INTO blog.t_post (nID, sAbout, sText, sLabels, dtPublic)"+
		" values (nextval('blog.seq_post_id'),$1,$2,$3,current_timestamp) returning nID, dtPublic", about, string(text), strings.Join(labels, ","))
	if err != nil {
		return 0, err
	}
	if !row.Next() {
		return 0, fmt.Errorf("Однострочный запрос не вернул ни одной записи")
	}
	tmp := BlogPost{
		About:  about,
		Text:   text,
		Labels: labels,
	}
	row.Scan(&tmp.ID, &tmp.PublicDate)
	(*posts)[tmp.ID] = tmp

	return tmp.ID, nil
}

//UpdateBlogPost --
func (posts *BlogPostArray) UpdateBlogPost(db *sql.DB, id int, about string, text template.HTML, labels []string) error {
	row, err := db.Query("UPDATE blog.t_post SET"+
		" sAbout=$1,sText=$2,sLabels=$3,dtPublic=current_timestamp"+
		" WHERE nID=$4"+
		" returning dtPublic", about, string(text), strings.Join(labels, ","), id)
	if err != nil {
		return err
	}
	if !row.Next() {
		return fmt.Errorf("Однострочный запрос не вернул ни одной записи")
	}
	tmp := BlogPost{
		ID:     id,
		About:  about,
		Text:   text,
		Labels: labels,
	}
	row.Scan(&tmp.PublicDate)
	(*posts)[id] = tmp

	return nil
}

//DeleteBlogPost --
func (posts *BlogPostArray) DeleteBlogPost(db *sql.DB, id int) error {
	_, err := db.Exec("DELETE FROM blog.t_post"+
		" WHERE nID=$1", id)
	if err != nil {
		return err
	}
	delete((*posts), id)

	return nil
}
