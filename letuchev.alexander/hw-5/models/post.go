package models

import (
	"database/sql"
	"fmt"
	"html/template"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
)

//BlogPost -
type BlogPost struct {
	//gorm.Model
	/*ID         int           `json:"id"    gorm:"AUTO_INCREMENT;primary_key"`
	About      string        `json:"about" gorm:"type:varchar(200)"`
	Text       template.HTML `json:"text"  gorm:"type:varchar(5000)"`
	Labels     []string      `json:"labels" gorm:"-"`
	PublicDate time.Time     `json:"publicated"`*/
	ID         int           `gorm:"column:nid;AUTO_INCREMENT;primary_key"`
	About      string        `gorm:"column:sabout;type:varchar(300)"`
	Text       template.HTML `gorm:"column:stext;type:varchar(20000)"`
	Labels     []string      `gorm:"-"`
	PublicDate time.Time     `gorm:"column:dtpublic"`
}

//TableName -
func (BlogPost) TableName() string {
	return "blog.t_post"
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
func (posts *BlogPostArray) NewBlogPost(db *gorm.DB, about string, text template.HTML, labels []string) (int, error) {
	tmp := BlogPost{
		ID:         -666,
		About:      about,
		Text:       text,
		Labels:     labels,
		PublicDate: time.Now(),
	}
	if err := db.Create(&tmp).Error; err != nil {
		fmt.Println(err)
		return 0, nil
	}
	if tmp.ID == -666 {
		fmt.Println("ID не был присвоен!")
	}
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
	_, err := db.Exec("DELETE FROM blog.t_post WHERE nID=$1", id)
	if err != nil {
		return err
	}
	delete((*posts), id)

	return nil
}
