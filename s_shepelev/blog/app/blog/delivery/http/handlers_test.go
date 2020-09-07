package http

import (
	"html/template"
	"net/http"
	"net/http/httptest"
	"path"
	"testing"

	"github.com/Toringol/group-golang-1/tree/master/s_shepelev/blog/app/blog"
	"github.com/Toringol/group-golang-1/tree/master/s_shepelev/blog/app/model"
	"github.com/Toringol/group-golang-1/tree/master/s_shepelev/blog/views"
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestGetPostInfoHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	blogCRUD := blog.NewMockRepository(ctrl)

	id := "sdgsdgsdgsdgsdgsg"
	posts := []*model.Post{
		{
			ID:          id,
			Title:       "New post",
			Description: "hello",
			Author:      "Toringol",
		},
	}

	blogCRUD.EXPECT().ListPosts().Return(posts, nil)

	blogHandlers := &blogHandlers{
		usecase: blogCRUD,
	}

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	templates := make(map[string]*template.Template)

	templates["posts"] = template.Must(template.ParseFiles(path.Join("../../../../views", "layout.html"),
		path.Join("../../../../views", "posts.html")))
	templates["post"] = template.Must(template.ParseFiles(path.Join("../../../../views", "layout.html"),
		path.Join("../../../../views", "post.html")))
	templates["newPost"] = template.Must(template.ParseFiles(path.Join("../../../../views", "layout.html"),
		path.Join("../../../../views", "newPost.html")))

	e.Renderer = &views.TemplateRegistry{
		Templates: templates,
	}

	if assert.NoError(t, blogHandlers.getAllPostsHandlers(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}
