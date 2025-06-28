package app

import (
	"database/sql"
	"e-store-backend/internal/db"
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

func NewApplication() (*Application, error) {
	logger := slog.New(tint.NewHandler(os.Stdout, &tint.Options{
		Level:      slog.LevelDebug,
		AddSource:  true,
		TimeFormat: time.Kitchen,
	}))

	db, err := db.OpenDB()
	if err != nil {
		return nil, err
	}

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
