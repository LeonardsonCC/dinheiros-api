package db

import (
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

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
		fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=postgres",
			os.Getenv("DB_USER"),
			"#pX#%q!V$Uux97/",
			os.Getenv("DB_HOST"),
			os.Getenv("DB_PORT"),
		))
	if err != nil {
		return nil, err
	}

	return db, nil
}
