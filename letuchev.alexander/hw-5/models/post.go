package models

import (
	"html/template"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
)

//BlogPost -
type BlogPost struct {
	ID         int           `gorm:"column:nid;AUTO_INCREMENT;primary_key" json:"id"`
	About      string        `gorm:"column:sabout;type:varchar(300)"       json:"about"`
	Text       template.HTML `gorm:"column:stext;type:varchar(20000)"      json:"text"`
	Labels     []string      `gorm:"-"                                     json:"labels"`
	LabelsJ    string        `gorm:"column:slabels;type:varchar(10000)"    json:"-"`
	PublicDate time.Time     `gorm:"column:dtpublic"                       json:"publicated"`
}

//TableName -
func (BlogPost) TableName() string {
	return "blog.t_post"
}

//BlogPostArray -
type BlogPostArray map[int]BlogPost

//Retrieve -
func Retrieve(db *gorm.DB) (*BlogPostArray, error) {
	var postSlice []BlogPost
	if err := db.Where("1 = 1").Find(&postSlice).Error; err != nil {
		return nil, err
	}

	posts := make(BlogPostArray, 10)
	for _, v := range postSlice {
		v.Labels = strings.Split(v.LabelsJ, ",")
		posts[v.ID] = v
	}
	return &posts, nil
}

//NewBlogPost --
func (posts *BlogPostArray) NewBlogPost(db *gorm.DB, about string, text template.HTML, labels []string) (int, error) {
	tmp := BlogPost{
		About:      about,
		Text:       text,
		Labels:     labels,
		LabelsJ:    strings.Join(labels, ","),
		PublicDate: time.Now(),
	}
	if err := db.Save(&tmp).Error; err != nil {
		return 0, err
	}
	(*posts)[tmp.ID] = tmp
	return tmp.ID, nil
}

//UpdateBlogPost --
func (posts *BlogPostArray) UpdateBlogPost(db *gorm.DB, id int, about string, text template.HTML, labels []string) error {
	tmp := BlogPost{
		ID:         id,
		About:      about,
		Text:       text,
		Labels:     labels,
		LabelsJ:    strings.Join(labels, ","),
		PublicDate: time.Now(),
	}

	if err := db.Save(tmp).Error; err != nil {
		return err
	}
	(*posts)[id] = tmp
	//fmt.Println(*posts)
	return nil
}

//DeleteBlogPost --
func (posts *BlogPostArray) DeleteBlogPost(db *gorm.DB, id int) error {
	tmp := BlogPost{ID: id}
	if err := db.Delete(&tmp).Error; err != nil {
		return err
	}
	delete((*posts), id)
	return nil
}
