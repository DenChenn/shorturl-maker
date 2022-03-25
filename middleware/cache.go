package middleware

import (
	"github.com/SwarzChen/url-shortener/constants"
	"github.com/SwarzChen/url-shortener/database"
	"github.com/gin-gonic/gin"
	"net/http"

	"context"
)

func Cache() gin.HandlerFunc {
	return func(c *gin.Context) {
		urlId := c.Param("url_id")
		if len(urlId) == 0 {
			c.JSON(constants.URL_ID_MISSING_ERROR.StatusCode, constants.URL_ID_MISSING_ERROR)
			c.Abort()
			return
		}

		redisClient := database.DatabaseClient.RedisClient
		// check whether this id has cache miss before
		ctx := context.Background()
		idContent, err := redisClient.Get(ctx, urlId).Result()
		if err != nil {
			// not in cache or expire
			c.Next()
			return
		}

		// this url id does not exist
		if idContent == constants.ID_NOT_EXIST {
			c.JSON(constants.URL_ID_NOT_EXIST_ERROR.StatusCode, constants.URL_ID_NOT_EXIST_ERROR)
			c.Abort()
			return
		}

		// this url content is in cache
		c.Redirect(http.StatusFound, idContent)
		c.Abort()
		return
	}
}
