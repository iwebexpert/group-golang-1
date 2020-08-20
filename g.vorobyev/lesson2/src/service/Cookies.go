package service

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func SetCookie(rc *gin.Context) {
	rc.SetCookie("CoolCookie", "CookieMonster", 3600, "/", "localhost",
		false, true)

	return
}

func ReadCookie(rc *gin.Context) {
	cookie, err := rc.Cookie("CoolCookie")

	if err != nil {
		rc.JSON(http.StatusNotFound, gin.H{"Cookie": "Not Found"})
		log.Panicf("ERROR while processing request %s", err.Error())
	}

	rc.JSON(http.StatusOK, gin.H{"Cookie": cookie})
}
