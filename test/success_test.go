package test

import (
	"bytes"
	"encoding/json"
	"github.com/SwarzChen/url-shortener/constants"
	"github.com/SwarzChen/url-shortener/model"
	"github.com/stretchr/testify/assert"
	"net/http"
	"os"
	"strings"
	"testing"
	"time"
)

func TestCreateUrlEntitySuccess(t *testing.T) {
	testUrlEntity := model.UrlEntityInDB{
		Url:      "test",
		ExpireAt: time.Now().In(constants.TIME_ZONE).Add(time.Hour).Format(time.RFC3339),
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

	var rspEntity model.UrlEntity
	err = json.NewDecoder(resp.Body).Decode(&rspEntity)
	if err != nil {
		panic(err.Error())
	}

	// split short url to check form
	partition := strings.Split(rspEntity.ShortUrl, "/")

	// checking
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Equal(t, 5, len(partition))
	assert.Equal(t, serverHost, "http://"+partition[2])
	assert.Equal(t, constants.CURRENT_VERSION, partition[3])
}

func TestRedirectToUrlSuccess(t *testing.T) {
	testUrl := "test.com"

	// Create url
	testUrlEntity := model.UrlEntityInDB{
		Url:      testUrl,
		ExpireAt: time.Now().In(constants.TIME_ZONE).Add(time.Hour).Format(time.RFC3339),
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

	var rspEntity model.UrlEntity
	err = json.NewDecoder(resp.Body).Decode(&rspEntity)
	if err != nil {
		panic(err.Error())
	}

	// test redirecting
	req, err = http.NewRequest("GET", serverHost+"/"+constants.CURRENT_VERSION+"/"+rspEntity.Id, nil)
	if err != nil {
		panic(err.Error())
	}

	client = &http.Client{
		// do not follow the redirect path, return 302 response directly
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			// skip first 301 redirect
			if len(via) <= 1 {
				return nil
			}
			return http.ErrUseLastResponse
		},
	}

	resp, err = client.Do(req)
	if err != nil {
		panic(err.Error())
	}
	defer resp.Body.Close()

	// check redirect location
	rsp302Header := resp.Header.Get("Location")
	partition := strings.Split(rsp302Header, "/")
	rspTargetLocationUrl := partition[len(partition)-1]

	assert.Equal(t, testUrl, rspTargetLocationUrl)
}
