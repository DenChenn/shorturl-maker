package constants

import (
	"fmt"
)

type ErrorEntity struct {
	StatusCode int
	ErrorCode  string
	Message    string
}

func (err *ErrorEntity) Error() string {
	return fmt.Sprintf("StatusCode: %d, ErrorCode: %s, Message: %s", err.StatusCode, err.ErrorCode, err.Message)
}

var SOME_ENV_VARIABLE_MISSING_ERROR = &ErrorEntity{
	StatusCode: 500,
	ErrorCode:  "SOME_ENV_VARIABLES_MISSING_ERROR",
	Message:    "Some environment variables are missing.",
}

var DATABASE_CONNETION_ERROR = &ErrorEntity{
	StatusCode: 500,
	ErrorCode:  "DATABASE_CONNETION_ERROR",
}

var SQL_WRAPPING_ERROR = &ErrorEntity{
	StatusCode: 500,
	ErrorCode:  "SQL_WRAPPING_ERROR",
}

var SCHEMA_PREPARATION_FAIL_ERROR = &ErrorEntity{
	StatusCode: 500,
	ErrorCode:  "SCHEMA_PREPARATION_FAIL_ERROR",
}

var TABLE_CREATION_ERROR = &ErrorEntity{
	StatusCode: 500,
	ErrorCode:  "TABLE_CREATION_ERROR",
}
