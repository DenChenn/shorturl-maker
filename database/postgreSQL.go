package database

import (
	"database/sql"
	"fmt"
	"github.com/SwarzChen/url-shortener/constants"
	_ "github.com/lib/pq"
	"log"
)

var PostgreSQLClient = ConstructDB()

type DB struct {
	Client *sql.DB
}

func ConstructDB() *DB {
	db := new(DB)
	err := db.Connect()
	if err != nil {
		log.Fatalln(err)
	}
	return db
}

func (db *DB) Connect() error {
	// fetch env variables
	err := constants.FetchEnvVariables()
	if err != nil {
		return err
	}

	dbURI := fmt.Sprintf("host=%s user=%s password=%s port=%s database=%s sslmode=disable", "127.0.0.1",
		constants.SQL_USER,
		constants.SQL_PASSWORD, "5432", constants.DATABASE_NAME)

	fmt.Println(dbURI)
	db.Client, err = sql.Open("postgres", dbURI)
	if err != nil {
		errorCopy := constants.DATABASE_CONNETION_ERROR
		errorCopy.Message = err.Error()
		return errorCopy
	}

	return nil
}
