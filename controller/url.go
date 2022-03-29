package controller

import (
	"context"
	"encoding/json"
	"github.com/SwarzChen/url-shortener/constants"
	"github.com/SwarzChen/url-shortener/database"
	"github.com/SwarzChen/url-shortener/model"
	"github.com/gin-gonic/gin"
	"github.com/segmentio/ksuid"
	"io/ioutil"
	"net/http"
	"time"
)

func GetUrlHandler(c *gin.Context) {
	urlId := c.Param("url_id")

	postgreSQLClient := database.DatabaseClient.PostgreSQLClient
	redisClient := database.DatabaseClient.RedisClient
	ctx := context.Background()

	// get url
	rows, err := postgreSQLClient.Query("SELECT * FROM url WHERE id=$1", urlId)
	// id query cause error
	if err != nil {
		// record to redis
		err = redisClient.Set(ctx, urlId, constants.ID_NOT_EXIST, time.Hour).Err()
		if err != nil {
			errorCopy := constants.REDIS_INSERTION_ERROR
			errorCopy.Message = err.Error()
			c.JSON(errorCopy.StatusCode, errorCopy)
			return
		}

		c.JSON(constants.URL_NOT_EXIST_OR_EXPIRE.StatusCode, constants.URL_NOT_EXIST_OR_EXPIRE)
		return
	}

	var urlEntity model.UrlEntityInDB
	count := 0
	for rows.Next() {
		count++
		err := rows.Scan(&urlEntity.Id, &urlEntity.Url, &urlEntity.ExpireAt, &urlEntity.CreateAt)
		if err != nil {
			errorCopy := constants.POSTGRESQL_ROW_SCANNING_ERROR
			errorCopy.Message = err.Error()
			c.JSON(errorCopy.StatusCode, errorCopy)
			return
		}
	}
	// Not found any row
	if count == 0 {
		// record to redis
		err = redisClient.Set(ctx, urlId, constants.ID_NOT_EXIST, time.Hour).Err()
		if err != nil {
			errorCopy := constants.REDIS_INSERTION_ERROR
			errorCopy.Message = err.Error()
			c.JSON(errorCopy.StatusCode, errorCopy)
			return
		}

		c.JSON(constants.URL_NOT_EXIST_OR_EXPIRE.StatusCode, constants.URL_NOT_EXIST_OR_EXPIRE)
		return
	}

	expireUTCTime, _ := time.Parse(time.RFC3339, urlEntity.ExpireAt)

	// specify time zone
	loc, _ := time.LoadLocation(constants.TIME_ZONE)
	currentTime := time.Now().In(loc)

	// This url expired, we will delete it
	if currentTime.After(expireUTCTime) {
		stmt, err := postgreSQLClient.Prepare("DELETE FROM url WHERE id=$1")
		if err != nil {
			errorCopy := constants.SCHEMA_PREPARATION_FAIL_ERROR
			errorCopy.Message = err.Error()
			c.JSON(errorCopy.StatusCode, errorCopy)
			return
		}

		_, err = stmt.Exec(urlId)
		if err != nil {
			errorCopy := constants.POSTGRESQL_DELETION_ERROR
			errorCopy.Message = err.Error()
			c.JSON(errorCopy.StatusCode, errorCopy)
			return
		}

		// record to cache
		err = redisClient.Set(ctx, urlEntity.Id, constants.ID_NOT_EXIST, time.Hour).Err()
		if err != nil {
			errorCopy := constants.REDIS_INSERTION_ERROR
			errorCopy.Message = err.Error()
			c.JSON(errorCopy.StatusCode, errorCopy)
			return
		}

		c.JSON(constants.URL_NOT_EXIST_OR_EXPIRE.StatusCode, constants.URL_NOT_EXIST_OR_EXPIRE)
		return
	}

	// calculate cache time
	cacheExpireTime := currentTime.Add(1 * time.Hour)
	if expireUTCTime.Before(cacheExpireTime) {
		cacheExpireTime = expireUTCTime
	}

	// put it to the cache
	err = redisClient.Set(ctx, urlEntity.Id, urlEntity.Url, cacheExpireTime.Sub(currentTime)).Err()
	if err != nil {
		errorCopy := constants.REDIS_INSERTION_ERROR
		errorCopy.Message = err.Error()
		c.JSON(errorCopy.StatusCode, errorCopy)
		return
	}

	// redirect to url
	c.Redirect(http.StatusFound, urlEntity.Url)
	return
}

func CreateUrlHandler(c *gin.Context) {
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		errorCopy := constants.REQUEST_BODY_MALFORMED
		errorCopy.Message = err.Error()
		c.JSON(errorCopy.StatusCode, errorCopy)
		return
	}

	var urlEntity model.UrlEntityInDB
	err = json.Unmarshal(body, &urlEntity)
	if err != nil {
		errorCopy := constants.REQUEST_BODY_PARSING_ERROR
		errorCopy.Message = err.Error()
		c.JSON(errorCopy.StatusCode, errorCopy)
		return
	}

	expireUTCTime, err := time.Parse(time.RFC3339, urlEntity.ExpireAt)
	if err != nil {
		errorCopy := constants.INVALID_TIME_FORMAT
		errorCopy.Message = err.Error()
		c.JSON(errorCopy.StatusCode, errorCopy)
		return
	}

	// specify time zone
	loc, _ := time.LoadLocation(constants.TIME_ZONE)
	currentTime := time.Now().In(loc)
	currentTimeString := time.Now().In(loc).Format(time.RFC3339)

	// check expire time validation
	if expireUTCTime.Before(currentTime) {
		c.JSON(constants.INPUT_EXPIRE_TIME_IS_INVALID.StatusCode, constants.INPUT_EXPIRE_TIME_IS_INVALID)
		return
	}

	// generate primitive id
	uuidString := ksuid.New().String()

	// insert into postgresql
	postgreSQLClient := database.DatabaseClient.PostgreSQLClient
	stmt, err := postgreSQLClient.Prepare("INSERT INTO url (id, url, expire_at, create_at) VALUES ($1, $2, $3, $4);")
	if err != nil {
		errorCopy := constants.SCHEMA_PREPARATION_FAIL_ERROR
		errorCopy.Message = err.Error()
		c.JSON(errorCopy.StatusCode, errorCopy)
		return
	}

	_, err = stmt.Exec(uuidString, urlEntity.Url, urlEntity.ExpireAt, currentTimeString)
	if err != nil {
		errorCopy := constants.POSTGRESQL_INSERTION_ERROR
		errorCopy.Message = err.Error()
		c.JSON(errorCopy.StatusCode, errorCopy)
		return
	}

	// take back database data to check and response
	rows, err := postgreSQLClient.Query("SELECT * FROM url WHERE id=$1", uuidString)
	if err != nil {
		errorCopy := constants.POSTGRESQL_INSERTION_ERROR
		errorCopy.Message = err.Error()
		c.JSON(errorCopy.StatusCode, errorCopy)
		return
	}

	var urlResponseEntity model.UrlEntity
	var ignore string
	for rows.Next() {
		err := rows.Scan(&urlResponseEntity.Id, &ignore, &ignore, &ignore)
		if err != nil {
			errorCopy := constants.POSTGRESQL_ROW_SCANNING_ERROR
			errorCopy.Message = err.Error()
			c.JSON(errorCopy.StatusCode, errorCopy)
			return
		}
	}

	urlResponseEntity.ShortUrl = constants.SERVER_HOST + "/" + constants.CURRENT_VERSION + "/" + uuidString

	c.JSON(http.StatusOK, urlResponseEntity)
	return
}
