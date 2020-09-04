package server

import (
	"blog/models"
	"bytes"
	"html/template"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

//Пришлось templates скопировать в папку server, чтобы ее было видно во время выполнения теста
func TestGetPage(t *testing.T) {
	testCases := []string{
		"index",
		"blog",
	}
	posts := make(models.BlogPostArray, 1)
	posts[1] = models.BlogPost{ID: 888, About: "zzz", Text: template.HTML("asdfasdfasdf"), PublicDate: time.Now()}
	srv := &BlogServer{Title: "Test", Posts: posts}

	for _, testCase := range testCases {
		reader := bytes.NewReader([]byte(""))
		req, _ := http.NewRequest("GET", "/", reader)

		r := httptest.NewRecorder()
		handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { srv.GetPage(w, r, testCase) })
		handler.ServeHTTP(r, req)
		//fmt.Print(r.Body.String())

		if strings.Contains(r.Body.String(), "Error") {
			t.Errorf(r.Body.String())
			continue
		}
	}
}
