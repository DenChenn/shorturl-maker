package schema

import (
	"github.com/SwarzChen/url-shortener/constants"
	"io/ioutil"
	"log"
	"os"
)

func SQLWrapper(filename string) (*string, error) {
	currentPath, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}

	c, ioErr := ioutil.ReadFile(currentPath + "/schema/" + filename)
	if ioErr != nil {
		errorCopy := constants.SQL_WRAPPING_ERROR
		errorCopy.Message = ioErr.Error()
		return nil, errorCopy
	}

	sql := string(c)
	return &sql, nil
}
