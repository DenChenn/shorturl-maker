package database

import (
	"database/sql"
	"fmt"
	"github.com/SwarzChen/url-shortener/constants"
	"github.com/go-redis/redis/v8"
	_ "github.com/lib/pq"
	"log"
)

var DatabaseClient = ConstructDB()

type DB struct {
	PostgreSQLClient *sql.DB
	RedisClient      *redis.Client
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

	// connect to postgreSQL
	dbURI := fmt.Sprintf("host=%s user=%s password=%s port=%s database=%s sslmode=disable", constants.POSTGRESQL_IP,
		constants.SQL_USER,
		constants.SQL_PASSWORD, constants.POSTGRESQL_PORT_NUMBER, constants.DATABASE_NAME)

	db.PostgreSQLClient, err = sql.Open("postgres", dbURI)
	if err != nil {
		errorCopy := constants.DATABASE_CONNETION_ERROR
		errorCopy.Message = err.Error()
		return errorCopy
	}

	// connect to redis
	db.RedisClient = redis.NewClient(&redis.Options{
		Addr:     constants.REDIS_IP + ":" + constants.REDIS_PORT_NUMBER,
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	return nil
}
