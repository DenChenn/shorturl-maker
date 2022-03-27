package constants

import (
	"fmt"
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
		fmt.Println("SQL_USER_MISSING")
		return SOME_ENV_VARIABLE_MISSING_ERROR
	}

	DATABASE_NAME, isExisted = os.LookupEnv("DATABASE_NAME")
	if !isExisted {
		fmt.Println("DATABASE_NAME_MISSING")
		return SOME_ENV_VARIABLE_MISSING_ERROR
	}

	POSTGRESQL_IP, isExisted = os.LookupEnv("POSTGRESQL_IP")
	if !isExisted {
		fmt.Println("POSTGRESQL_IP_MISSING")
		return SOME_ENV_VARIABLE_MISSING_ERROR
	}

	REDIS_IP, isExisted = os.LookupEnv("REDIS_IP")
	if !isExisted {
		fmt.Println("REDIS_IP_MISSING")
		return SOME_ENV_VARIABLE_MISSING_ERROR
	}

	POSTGRESQL_PORT_NUMBER, isExisted = os.LookupEnv("POSTGRESQL_PORT_NUMBER")
	if !isExisted {
		fmt.Println("POSTGRESQL_PORT_NUMBER_MISSING")
		return SOME_ENV_VARIABLE_MISSING_ERROR
	}

	REDIS_PORT_NUMBER, isExisted = os.LookupEnv("REDIS_PORT_NUMBER")
	if !isExisted {
		fmt.Println("REDIS_PORT_NUMBER_MISSING")
		return SOME_ENV_VARIABLE_MISSING_ERROR
	}

	SERVER_HOST, isExisted = os.LookupEnv("SERVER_HOST")
	if !isExisted {
		fmt.Println("SERVER_HOST_MISSING")
		return SOME_ENV_VARIABLE_MISSING_ERROR
	}

	return nil
}
