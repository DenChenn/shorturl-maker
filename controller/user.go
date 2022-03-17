package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetUrlHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"Message": "Success",
	})
	return
}

func CreateUrlHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"Message": "Success",
	})
	return
}
