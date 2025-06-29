package database

import (
	"database/sql"
	"embed"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/pressly/goose/v3"
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

	fmt.Println("Successfully connected to database...")
	return db, nil
}

func Migrate(db *sql.DB, migrationsFS embed.FS, migrationsDir string) error {
	goose.SetBaseFS(migrationsFS)
	defer goose.SetBaseFS(nil)

	if err := goose.SetDialect("mysql"); err != nil {
		return err
	}

	if err := goose.Up(db, migrationsDir); err != nil {
		return err
	}

	return nil
}
