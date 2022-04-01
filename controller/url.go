package controller

import (
	"database/sql"
	"encoding/json"
	"github.com/SwarzChen/url-shortener/constants"
	"github.com/SwarzChen/url-shortener/database"
	"github.com/SwarzChen/url-shortener/model"
	"github.com/gin-gonic/gin"
	"github.com/teris-io/shortid"
	"io/ioutil"
	"net/http"
	"time"
)

func GetUrlHandler(c *gin.Context) {
	urlId := c.Param("url_id")

	postgreSQLClient := database.DatabaseClient.PostgreSQLClient
	redisClient := database.DatabaseClient.RedisClient
	ctx := c.Request.Context()

	var urlEntity model.UrlEntityInDB
	// get url
	err := postgreSQLClient.QueryRowContext(ctx, "SELECT * FROM url WHERE id=$1", urlId).Scan(&urlEntity.Id,
		&urlEntity.Url, &urlEntity.ExpireAt, &urlEntity.CreateAt)

	switch {
	// not exist in database
	case err == sql.ErrNoRows:
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

	// other kinds database error
	case err != nil:
		errorCopy := constants.POSTGRESQL_QUERY_ERROR
		errorCopy.Message = err.Error()
		c.JSON(errorCopy.StatusCode, errorCopy)
		return
	}

	expireUTCTime, _ := time.Parse(time.RFC3339, urlEntity.ExpireAt)
	currentTime := time.Now().In(constants.TIME_ZONE)

	// This url expired, we will delete it
	if currentTime.After(expireUTCTime) {
		// remove expired url from database
		_, err = postgreSQLClient.ExecContext(ctx, "DELETE FROM url WHERE id=$1", urlId)
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

	currentTime := time.Now().In(constants.TIME_ZONE)
	currentTimeString := time.Now().In(constants.TIME_ZONE).Format(time.RFC3339)

	// check expire time validation
	if expireUTCTime.Before(currentTime) {
		c.JSON(constants.INPUT_EXPIRE_TIME_IS_INVALID.StatusCode, constants.INPUT_EXPIRE_TIME_IS_INVALID)
		return
	}

	// generate short unique id
	generator, err := shortid.New(1, shortid.DefaultABC, 2342)
	if err != nil {
		errorCopy := constants.SETTING_GENERATOR_SEED_ERROR
		errorCopy.Message = err.Error()
		c.JSON(errorCopy.StatusCode, errorCopy)
		return
	}

	sid, err := generator.Generate()
	if err != nil {
		errorCopy := constants.ID_GENERATION_ERROR
		errorCopy.Message = err.Error()
		c.JSON(errorCopy.StatusCode, errorCopy)
		return
	}

	// insert into postgresql
	ctx := c.Request.Context()
	postgreSQLClient := database.DatabaseClient.PostgreSQLClient

	_, err = postgreSQLClient.ExecContext(ctx, "INSERT INTO url (id, url, expire_at, create_at) VALUES ($1, $2, $3, "+
		"$4);", sid, urlEntity.Url, urlEntity.ExpireAt, currentTimeString)
	if err != nil {
		errorCopy := constants.POSTGRESQL_INSERTION_ERROR
		errorCopy.Message = err.Error()
		c.JSON(errorCopy.StatusCode, errorCopy)
		return
	}

	urlResponseEntity := model.UrlEntity{
		Id:       sid,
		ShortUrl: constants.SERVER_HOST + "/" + constants.CURRENT_VERSION + "/" + sid,
	}

	c.JSON(http.StatusOK, urlResponseEntity)
	return
}
