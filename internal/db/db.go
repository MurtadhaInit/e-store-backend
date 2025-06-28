package db

import (
	"database/sql"
	"errors"
	"os"
	"time"
)

func OpenDB() (*sql.DB, error) {
	dsn, dsn_set := os.LookupEnv("DSN")
	if !dsn_set || dsn == "" {
		return nil, errors.New("DSN environment variable not set. Could not connect to the database.")
	}

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		db.Close()
		return nil, err
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	// db.SetMaxOpenConns(10)
	// db.SetMaxIdleConns(10)

	return db, nil
}
