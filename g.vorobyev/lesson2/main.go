package main

import (
	"github.com/gin-gonic/gin"
	"service"
)

func main() {
	r := gin.Default()
	r.POST("/request", service.RequestHandler)
	r.GET("/writeCookie", service.SetCookie)
	r.GET("/readCookie", service.ReadCookie)
	_ = r.Run(":8080")
}
