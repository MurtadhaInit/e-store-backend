package app

import (
	"database/sql"
	"e-store-backend/db/migrations"
	"e-store-backend/internal/database"
	"e-store-backend/internal/handlers"
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
	Logger   *slog.Logger
	DB       *sql.DB
	Handlers *handlers.Handlers
}

func NewApplication() (*Application, error) {
	logger := slog.New(tint.NewHandler(os.Stdout, &tint.Options{
		Level:      slog.LevelDebug,
		AddSource:  true,
		TimeFormat: time.Kitchen,
	}))

	db, err := database.OpenDB()
	if err != nil {
		return nil, err
	}

	if err := database.Migrate(db, migrations.EmbeddedMigrations, "."); err != nil {
		return nil, err
	}

	queries := repository.New(db)

	if err := database.SeedDatabase(db, queries); err != nil {
		return nil, err
	}

	handlers := handlers.NewHandlers(logger, db, queries)

	app := &Application{
		Logger:   logger,
		DB:       db,
		Handlers: handlers,
	}

	return app, nil
}

func (app *Application) HealthCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Service is healthy")
}
