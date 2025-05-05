package db

import (
	"fmt"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var db *sqlx.DB

var (
	dbUser = getEnvOrPanic("DB_USER")
	dbPass = getEnvOrPanic("DB_PASS")
	dbHost = getEnvOrPanic("DB_HOST")
	dbPort = getEnvOrPanic("DB_PORT")
)

func GetConnection() (*sqlx.DB, error) {
	if db != nil {
		return db, nil
	}

	d, err := connect()
	if err != nil {
		return nil, err
	}

	db = d
	return db, nil
}

func connect() (*sqlx.DB, error) {
	db, err := sqlx.Connect(
		"postgres",
		fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=postgres sslmode=disable",
			os.Getenv("DB_USER"),
			os.Getenv("DB_PASS"),
			os.Getenv("DB_HOST"),
			os.Getenv("DB_PORT"),
		))
	if err != nil {
		return nil, err
	}

	return db, nil
}

func getEnvOrPanic(key string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}

	log.Fatalf("failed to get env: [%s]", key)
	return ""
}
