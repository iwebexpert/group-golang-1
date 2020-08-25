package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	_ "github.com/go-sql-driver/mysql"
)

// Post - описание структуры поста
type Post struct {
	ID       int
	Header   string
	Text     string
	Date     time.Time
	Comments []Comment
}

// Comment - комментарии к постам
type Comment struct {
	PostID int
	Text   string
}

var tmpl = template.Must(template.New("myBlog").ParseFiles("tmpl.html"))
var database *sql.DB

func allPosts(w http.ResponseWriter, r *http.Request) {
	posts := []Post{}

	rows, err := database.Query("select id, header, text from post")
	if err != nil {
		log.Fatal(err)
		w.WriteHeader(500)
	}

	defer rows.Close()

	for rows.Next() {
		post := Post{}

		err = rows.Scan(&post.ID, &post.Header, &post.Text)
		if err != nil {
			log.Println(err)
			continue
		}
		posts = append(posts, post)
	}

	if err := tmpl.ExecuteTemplate(w, "allPosts", posts); err != nil {
		log.Println(err)
	}
}

func selectPost(w http.ResponseWriter, r *http.Request) {
	post := Post{}
	PostID := r.URL.Query().Get("id")

	rows, err := database.Query(fmt.Sprintf("select id, header, text from post WHERE id = %v", PostID))
	if err != nil {
		log.Fatal(err)
		w.WriteHeader(500)
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&post.ID, &post.Header, &post.Text)
		if err != nil {
			log.Println(err)
			continue
		}
	}

	if err := tmpl.ExecuteTemplate(w, "selectPost", post); err != nil {
		log.Println(err)
	}
}

func newPost(w http.ResponseWriter, r *http.Request) {
	if err := tmpl.ExecuteTemplate(w, "newPost", nil); err != nil {
		log.Println(err)
	}
}

func createPost(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	log.Println("Creating new post")
	log.Println(r.Form)
	newHeader := r.Form["Header"][0]
	newText := r.Form["Text"][0]

	_, err := database.Exec("INSERT INTO post (Header, Text) VALUES (?, ?)", newHeader, newText)
	if err != nil {
		log.Fatal(err)
		w.WriteHeader(500)
	}

	http.Redirect(w, r, "/", 301)
}

func main() {
	db, err := sql.Open("mysql", "root:1@tcp(192.168.0.145)/my_blog")
	if err != nil {
		log.Fatal(err)
	}
	database = db

	defer database.Close()
	route := chi.NewRouter()

	route.Route("/", func(r chi.Router) {
		r.Get("/", allPosts)
		r.Get("/post", selectPost)
		r.Get("/newpost", newPost)
		r.Post("/newpost/new", createPost)
	})

	log.Fatal(http.ListenAndServe(":8080", route))
}
