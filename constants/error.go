package constants

import (
	"fmt"
	"net/http"
)

type ErrorEntity struct {
	StatusCode int
	ErrorCode  string
	Message    string
}

func (err *ErrorEntity) Error() string {
	return fmt.Sprintf("StatusCode: %d, ErrorCode: %s, Message: %s", err.StatusCode, err.ErrorCode, err.Message)
}

// STATUS_INTERNAL_SERVER_ERROR
var SOME_ENV_VARIABLE_MISSING_ERROR = &ErrorEntity{
	StatusCode: http.StatusInternalServerError,
	ErrorCode:  "SOME_ENV_VARIABLES_MISSING_ERROR",
	Message:    "Some environment variables are missing.",
}

var DATABASE_CONNETION_ERROR = &ErrorEntity{
	StatusCode: http.StatusInternalServerError,
	ErrorCode:  "DATABASE_CONNETION_ERROR",
}

var REDIS_INSERTION_ERROR = &ErrorEntity{
	StatusCode: http.StatusInternalServerError,
	ErrorCode:  "REDIS_INSERTION_ERROR",
}
var POSTGRESQL_DELETION_ERROR = &ErrorEntity{
	StatusCode: http.StatusInternalServerError,
	ErrorCode:  "POSTGRESQL_DELETION_ERROR",
}
var REQUEST_BODY_PARSING_ERROR = &ErrorEntity{
	StatusCode: http.StatusInternalServerError,
	ErrorCode:  "REQUEST_BODY_PARSING_ERROR",
}
var POSTGRESQL_INSERTION_ERROR = &ErrorEntity{
	StatusCode: http.StatusInternalServerError,
	ErrorCode:  "POSTGRESQL_INSERTION_ERROR",
}
var POSTGRESQL_QUERY_ERROR = &ErrorEntity{
	StatusCode: http.StatusInternalServerError,
	ErrorCode:  "POSTGRESQL_QUERY_ERROR",
}
var SETTING_GENERATOR_SEED_ERROR = &ErrorEntity{
	StatusCode: http.StatusInternalServerError,
	ErrorCode:  "SETTING_GENERATOR_SEED_ERROR",
}
var ID_GENERATION_ERROR = &ErrorEntity{
	StatusCode: http.StatusInternalServerError,
	ErrorCode:  "ID_GENERATION_ERROR",
}

// STATUS_NOT_FOUND
var URL_ID_MISSING_ERROR = &ErrorEntity{
	StatusCode: http.StatusNotFound,
	ErrorCode:  "URL_ID_MISSING_ERROR",
	Message:    "There is no url_id in url.",
}
var URL_ID_NOT_EXIST_ERROR = &ErrorEntity{
	StatusCode: http.StatusNotFound,
	ErrorCode:  "URL_ID_NOT_EXIST_ERROR",
	Message:    "This url id does not exist",
}
var URL_NOT_EXIST_OR_EXPIRE = &ErrorEntity{
	StatusCode: http.StatusNotFound,
	ErrorCode:  "URL_NOT_EXIST_OR_EXPIRE",
	Message:    "This url was expired or it does not exist.",
}

// Bad Request
var REQUEST_BODY_MALFORMED = &ErrorEntity{
	StatusCode: http.StatusBadRequest,
	ErrorCode:  "REQUEST_BODY_MALFORMED",
}
var INVALID_TIME_FORMAT = &ErrorEntity{
	StatusCode: http.StatusBadRequest,
	ErrorCode:  "INVALID_TIME_FORMAT",
}
var INPUT_EXPIRE_TIME_IS_INVALID = &ErrorEntity{
	StatusCode: http.StatusBadRequest,
	ErrorCode:  "INPUT_EXPIRE_TIME_IS_INVALID",
	Message:    "Expiration time is behind current time, which is no sense.",
}
