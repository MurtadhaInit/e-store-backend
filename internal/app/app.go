package app

import (
	"database/sql"
	"e-store-backend/db/migrations"
	"e-store-backend/internal/database"
	"e-store-backend/internal/handlers"
	"e-store-backend/internal/middleware"
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
	Logger     *slog.Logger
	DB         *sql.DB
	Handlers   *handlers.Handlers
	Middleware *middleware.Middleware
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

	handlers := handlers.CreateHandlers(logger, db, queries)
	middleware := middleware.CreateMiddleware(logger, queries)

	app := &Application{
		Logger:     logger,
		DB:         db,
		Handlers:   handlers,
		Middleware: middleware,
	}

	return app, nil
}

func (app *Application) HealthCheck(w http.ResponseWriter, _ *http.Request) {
	// Health probes only read the status code; if the body write fails the
	// client has already disconnected, so there is nothing useful to recover.
	_, _ = fmt.Fprintln(w, "Service is healthy")
}
