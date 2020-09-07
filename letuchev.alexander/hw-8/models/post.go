package models

import (
	"database/sql"
	"fmt"
	"html/template"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
)

var log *logrus.Logger

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

//SetLogger -
func SetLogger(lg *logrus.Logger) {
	log = lg
}

//Retrieve -
func Retrieve(db *sql.DB, lg *logrus.Logger) (*BlogPostArray, error) {
	SetLogger(lg) //Подцепили логгер в глобальную переменную пакета
	log.Debug("Загрузка списка постов из БД")
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
	log.Debug(fmt.Sprintf("Создание нового поста с заголовком [%s]", about))
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
	log.Debug(fmt.Sprintf("Присвоен id [%d]", tmp.ID))

	return tmp.ID, nil
}

//UpdateBlogPost --
func (posts *BlogPostArray) UpdateBlogPost(db *sql.DB, id int, about string, text template.HTML, labels []string) error {
	log.Debug("Обновление поста [", id, "] с заголовком [", about, "]")
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
	log.Debug("Удаление поста [", id, "] с заголовком [", (*posts)[id].About, "]")
	_, err := db.Exec("DELETE FROM blog.t_post"+
		" WHERE nID=$1", id)
	if err != nil {
		return err
	}
	delete((*posts), id)
	return nil
}
