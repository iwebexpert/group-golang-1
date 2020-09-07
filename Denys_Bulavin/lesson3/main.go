package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"text/template"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

type PageArticles struct {
	Title    string
	Articles Articles
}

type Articles []Article

type Article struct {
	ArticleID    int
	TitleArticle string
	TextArticle  string
}

func (a *PageArticles) getListArticles(w http.ResponseWriter, r *http.Request) {

	f, _ := os.Open("./www/articles.html")
	data, _ := ioutil.ReadAll(f)

	tempArticles := template.Must(template.New("Articles").Parse(string(data)))
	tempArticles.ExecuteTemplate(w, "Articles", a)

}

func (a *PageArticles) getArticle(w http.ResponseWriter, r *http.Request) {

	id, _ := strconv.Atoi(chi.URLParam(r, "articleID"))

	for _, value := range a.Articles {
		if id == value.ArticleID {
			f, _ := os.Open("./www/article.html")
			data, _ := ioutil.ReadAll(f)
			tempArticle := template.Must(template.New("Article").Parse(string(data)))
			tempArticle.ExecuteTemplate(w, "Article", value)
		}
	}

}

func (a *PageArticles) createArticle(w http.ResponseWriter, r *http.Request) {

	f, _ := os.Open("./www/create.html")
	data, _ := ioutil.ReadAll(f)

	tempArticles := template.Must(template.New("createArticle").Parse(string(data)))
	tempArticles.ExecuteTemplate(w, "createArticle", a)
}

func (a *PageArticles) addArticle(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}
	a.Articles = append(a.Articles, Article{ArticleID: 4, TitleArticle: r.FormValue("TitleArticle"), TextArticle: r.FormValue("TextArticle")})

	fmt.Fprintln(w, "PageAdd")

}

func main() {
	r := chi.NewRouter()

	r.Use(middleware.Logger)

	articles := PageArticles{
		Title: "Статьи",
		Articles: Articles{
			{ArticleID: 1, TitleArticle: "Статья 1", TextArticle: "Текст к статьй 1"},
			{ArticleID: 2, TitleArticle: "Статья 2", TextArticle: "Текст к статьй 2"},
			{ArticleID: 3, TitleArticle: "Статья 3", TextArticle: "Текст к статьй 3"},
		},
	}

	r.Route("/articles", func(r chi.Router) {
		r.Get("/", articles.getListArticles)
		r.Get("/{articleID}", articles.getArticle)
		r.Get("/create", articles.createArticle)
		r.Post("/add", articles.addArticle)
	})

	http.ListenAndServe(":8080", r)
}
