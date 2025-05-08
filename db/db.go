package db

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/uptrace/opentelemetry-go-extra/otelsql"
	"github.com/uptrace/opentelemetry-go-extra/otelsqlx"
)

var db *sqlx.DB

var (
	dbUser = getEnvOrPanic("DB_USER")
	dbPass = getEnvOrPanic("DB_PASS")
	dbHost = getEnvOrPanic("DB_HOST")
	dbPort = getEnvOrPanic("DB_PORT")
)

func GetConnection(ctx context.Context) (*sqlx.DB, error) {
	if db != nil {
		return db, nil
	}

	d, err := connect(ctx)
	if err != nil {
		return nil, err
	}

	db = d
	return db, nil
}

func connect(ctx context.Context) (*sqlx.DB, error) {
	db, err := otelsqlx.ConnectContext(
		ctx,
		"postgres",
		fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=postgres sslmode=disable",
			dbUser,
			dbPass,
			dbHost,
			dbPort,
		),
		otelsql.WithDBName("dinheiros-postgres"),
		otelsql.WithDBSystem("postgres"))
	if err != nil {
		return nil, err
	}

	err = db.PingContext(ctx)
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
