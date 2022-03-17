package utils

import (
	"github.com/SwarzChen/url-shortener/constants"
	"github.com/SwarzChen/url-shortener/database"
	"github.com/SwarzChen/url-shortener/schema"
)

func CreateTable() error {
	createTableSchema, err := schema.SQLWrapper("create_url_table.sql")
	if err != nil {
		return err
	}

	pqClient := database.PostgreSQLClient.Client
	stmt, err := pqClient.Prepare(*createTableSchema)
	if err != nil {
		errorCopy := constants.SCHEMA_PREPARATION_FAIL_ERROR
		errorCopy.Message = err.Error()
		return errorCopy
	}

	_, err = stmt.Exec()
	if err != nil {
		errorCopy := constants.TABLE_CREATION_ERROR
		errorCopy.Message = err.Error()
		return errorCopy
	}

	return nil
}
