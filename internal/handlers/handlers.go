package handlers

import (
	"database/sql"
	"e-store-backend/internal/repository"
	"log/slog"
)

type Handlers struct {
	logger  *slog.Logger
	db      *sql.DB
	queries *repository.Queries
}

func CreateHandlers(logger *slog.Logger, db *sql.DB, queries *repository.Queries) *Handlers {
	handlers := &Handlers{
		logger:  logger,
		db:      db,
		queries: queries,
	}

	return handlers
}
