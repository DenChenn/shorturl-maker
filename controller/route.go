package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func MissingRouteHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"Message": "Success",
	})
	return
}
