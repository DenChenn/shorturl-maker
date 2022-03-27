package constants

import (
	"os"
)

var SQL_USER string
var SQL_PASSWORD string
var DATABASE_NAME string
var POSTGRESQL_IP string
var REDIS_IP string
var POSTGRESQL_PORT_NUMBER string
var REDIS_PORT_NUMBER string
var SERVER_HOST string

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

	DATABASE_NAME, isExisted = os.LookupEnv("DATABASE_NAME")
	if !isExisted {
		return SOME_ENV_VARIABLE_MISSING_ERROR
	}

	POSTGRESQL_IP, isExisted = os.LookupEnv("POSTGRESQL_IP")
	if !isExisted {
		return SOME_ENV_VARIABLE_MISSING_ERROR
	}

	REDIS_IP, isExisted = os.LookupEnv("REDIS_IP")
	if !isExisted {
		return SOME_ENV_VARIABLE_MISSING_ERROR
	}

	POSTGRESQL_PORT_NUMBER, isExisted = os.LookupEnv("POSTGRESQL_PORT_NUMBER")
	if !isExisted {
		return SOME_ENV_VARIABLE_MISSING_ERROR
	}

	REDIS_PORT_NUMBER, isExisted = os.LookupEnv("REDIS_PORT_NUMBER")
	if !isExisted {
		return SOME_ENV_VARIABLE_MISSING_ERROR
	}

	SERVER_HOST, isExisted = os.LookupEnv("SERVER_HOST")
	if !isExisted {
		return SOME_ENV_VARIABLE_MISSING_ERROR
	}

	return nil
}
