package test

import (
	"bytes"
	"encoding/json"
	"github.com/SwarzChen/url-shortener/constants"
	"github.com/SwarzChen/url-shortener/model"
	"github.com/stretchr/testify/assert"
	"net/http"
	"os"
	"testing"
	"time"
)

func TestCreateUrlWithInvalidExpireTime(t *testing.T) {
	// create invalid expire time
	testUrlEntity := model.UrlEntityInDB{
		Url:      "test",
		ExpireAt: time.Now().In(constants.TIME_ZONE).Add(-time.Hour).Format(time.RFC3339),
	}

	body, err := json.Marshal(testUrlEntity)
	if err != nil {
		panic(err.Error())
	}

	// get server host
	serverHost, isExisted := os.LookupEnv("SERVER_HOST")
	if !isExisted {
		panic("MISSING_SERVER_HOST_ENV_VAR")
	}

	req, err := http.NewRequest("POST", serverHost+"/"+constants.CURRENT_VERSION+"/urls/", bytes.NewBuffer(body))
	if err != nil {
		panic(err.Error())
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err.Error())
	}
	defer resp.Body.Close()

	var rspEntity constants.ErrorEntity
	err = json.NewDecoder(resp.Body).Decode(&rspEntity)
	if err != nil {
		panic(err.Error())
	}

	// checking error handling
	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
	assert.Equal(t, constants.INPUT_EXPIRE_TIME_IS_INVALID.ErrorCode, rspEntity.ErrorCode)
}
