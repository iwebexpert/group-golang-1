package models

import (
	"blog/config"
	"database/sql"
	"fmt"
	"html/template"
	"strings"
	"testing"
	"time"

	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/spf13/viper"
)

var gDb *sql.DB

func TestDbConnect(t *testing.T) {
	err := config.Parse()
	if err != nil {
		t.Errorf("Ошибка чтения файла конфигурации %s", err)
		return
	}

	gDb, err = sql.Open("pgx", fmt.Sprintf("user=%s password=%s host=%s port=%s database=%s",
		viper.GetString("user"), viper.GetString("password"), viper.GetString("host"),
		viper.GetString("port"), viper.GetString("database")))
	if err != nil {
		t.Errorf("Не удалось соединиться с БД: %s", err)
		return
	}
}

func TestNewBlogPost(t *testing.T) {
	TestDbConnect(t)
	defer gDb.Close()

	etAbout := "zzzzzzzzzzzzz"
	etText := "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"
	etLabels := []string{"l1", "l2", "l3"}
	etPost := BlogPost{0, etAbout, template.HTML(etText), etLabels, time.Now()}

	posts := make(BlogPostArray, 1)
	id, err := posts.NewBlogPost(gDb, etAbout, template.HTML(etText), etLabels)

	allPosts, err := Retrieve(gDb)
	if err != nil {
		t.Errorf("Ошибка чтения записей из БД: %s", err)
		return
	}
	if posts[id].About != etPost.About || posts[id].Text != etPost.Text ||
		strings.Join(posts[id].Labels, ",") != strings.Join(etPost.Labels, ",") {
		t.Errorf("Закэшировано значение:\r\n%v\r\nэталон:\r\n%v", posts[id], etPost)
	}
	if etPost.About != (*allPosts)[id].About || etPost.Text != (*allPosts)[id].Text ||
		strings.Join(etPost.Labels, ",") != strings.Join((*allPosts)[id].Labels, ",") {
		t.Errorf("В БД записано значение:\r\n%v\r\nэталон:\r\n%v", (*allPosts)[id], etPost)
	}

	err = allPosts.DeleteBlogPost(gDb, id)
	if err != nil {
		t.Errorf("Ошибка удаления тестовой записи из БД: %s", err)
	}
}

func BenchmarkRetrieve(b *testing.B) {
	var t *testing.T
	TestDbConnect(t)
	defer gDb.Close()

	for i := 0; i < b.N; i++ {
		_, _ = Retrieve(gDb)
	}
}
