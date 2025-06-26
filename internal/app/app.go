package app

import (
	"database/sql"
	"e-store-backend/internal/repository"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/lmittmann/tint"
)

type Application struct {
	Logger  *slog.Logger
	DB      *sql.DB
	Queries *repository.Queries
}

func NewApplication(dsn string) (*Application, error) {
	logger := slog.New(tint.NewHandler(os.Stdout, &tint.Options{
		Level:      slog.LevelDebug,
		AddSource:  true,
		TimeFormat: time.Kitchen,
	}))

	db, err := openDB(dsn)
	if err != nil {
		// logger.Error(err.Error())
		// os.Exit(1)
		return nil, err
	}
	// defer db.Close()

	queries := repository.New(db)

	app := &Application{
		Logger:  logger,
		DB:      db,
		Queries: queries,
	}

	return app, nil
}

func (app *Application) HealthCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Service is healthy")
}

func openDB(dsn string) (*sql.DB, error) {
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
