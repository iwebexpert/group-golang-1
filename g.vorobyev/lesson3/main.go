package main

import (
	"blog"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.GET("/", blog.ShowPosts)
	b := r.Group("/blog")
	{
		b.GET("/post/:id", blog.ShowPost)
		b.POST("/edit/:id", blog.EditPost)
		b.POST("/add", blog.AddPost)
		b.GET("/edit/:id", blog.ShowEditPost)
		b.GET("/add", blog.ShowAddPost)
	}
	_ = r.Run(":8080")
}
