package middleware

import (
	"e-store-backend/internal/repository"
	"log/slog"
)

type Middleware struct {
	logger  *slog.Logger
	queries *repository.Queries
}

func CreateMiddleware(logger *slog.Logger, queries *repository.Queries) *Middleware {
	return &Middleware{
		logger:  logger,
		queries: queries,
	}
}
