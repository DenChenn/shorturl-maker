package test

import (
	"encoding/json"
	"github.com/SwarzChen/url-shortener/constants"
	"github.com/segmentio/ksuid"
	"github.com/stretchr/testify/assert"
	"net/http"
	"os"
	"testing"
)

func TestCacheHandlingInvalidUrlQuery(t *testing.T) {
	// get server host
	serverHost, isExisted := os.LookupEnv("SERVER_HOST")
	if !isExisted {
		panic("MISSING_SERVER_HOST_ENV_VAR")
	}

	// random generate not exist url

	notExistedId := ksuid.New().String()

	// first query
	req, err := http.NewRequest("GET", serverHost+"/"+constants.CURRENT_VERSION+"/"+notExistedId, nil)
	if err != nil {
		panic(err.Error())
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err.Error())
	}

	var rspEntity constants.ErrorEntity
	err = json.NewDecoder(resp.Body).Decode(&rspEntity)
	if err != nil {
		panic(err.Error())
	}

	// first query is not caught by the cache
	assert.Equal(t, http.StatusNotFound, rspEntity.StatusCode)
	assert.Equal(t, constants.URL_NOT_EXIST_OR_EXPIRE.ErrorCode, rspEntity.ErrorCode)

	// second query
	resp, err = client.Do(req)
	if err != nil {
		panic(err.Error())
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&rspEntity)
	if err != nil {
		panic(err.Error())
	}

	// second query is caught by cache
	assert.Equal(t, http.StatusNotFound, rspEntity.StatusCode)
	assert.Equal(t, constants.URL_ID_NOT_EXIST_ERROR.ErrorCode, rspEntity.ErrorCode)
}
