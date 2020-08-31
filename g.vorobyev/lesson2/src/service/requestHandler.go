package service

import (
	"github.com/gin-gonic/gin"
	"log"
	"model"
	"net/http"
)

func RequestHandler(rc *gin.Context) {
	var req model.RequestType

	err := rc.ShouldBindJSON(&req)
	if err != nil {
		rc.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		log.Panicf("ERROR while processing request %s", err.Error())
	}

	v, err := searchV1(req.Search, req.Sites)

	if err != nil {
		rc.JSON(http.StatusNotFound, gin.H{"status": err.Error()})
		return
	}

	rc.JSON(http.StatusOK, gin.H{"Found": v})
}
