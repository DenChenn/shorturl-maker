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
	loc, _ := time.LoadLocation(constants.TIME_ZONE)

	testUrlEntity := model.UrlEntityInDB{
		Url:      "test",
		ExpireAt: time.Now().In(loc).Add(time.Hour).Format(time.RFC3339),
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

	req, err := http.NewRequest("POST", serverHost+"/v1/urls/", bytes.NewBuffer(body))
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
	assert.Equal(t, "v1", partition[3])
	assert.Equal(t, 27, len(partition[4]))
}
