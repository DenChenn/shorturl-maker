package constants

import (
	"os"
)

var SQL_USER string
var SQL_PASSWORD string
var INSTANCE_CONNECTION_NAME string
var DATABASE_NAME string
var DATABASE_SOCKET_DIR string

func FetchEnvVariables() error {
	var isExisted bool
	SQL_USER, isExisted = os.LookupEnv("SQL_USER")
	if !isExisted {
		return SOME_ENV_VARIABLE_MISSING_ERROR
	}

	SQL_PASSWORD, isExisted = os.LookupEnv("SQL_PASSWORD")
	if !isExisted {
		return SOME_ENV_VARIABLE_MISSING_ERROR
	}

	INSTANCE_CONNECTION_NAME, isExisted = os.LookupEnv("INSTANCE_CONNECTION_NAME")
	if !isExisted {
		return SOME_ENV_VARIABLE_MISSING_ERROR
	}

	DATABASE_NAME, isExisted = os.LookupEnv("DATABASE_NAME")
	if !isExisted {
		return SOME_ENV_VARIABLE_MISSING_ERROR
	}

	DATABASE_SOCKET_DIR, isExisted = os.LookupEnv("DB_SOCKET_DIR")
	if !isExisted {
		DATABASE_SOCKET_DIR = "/cloudsql"
	}

	return nil
}
