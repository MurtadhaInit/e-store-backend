package middleware

import (
	"e-store-backend/internal/repository"
)

type Middleware struct {
	queries *repository.Queries
}

func CreateMiddleware(queries *repository.Queries) *Middleware {
	return &Middleware{
		queries: queries,
	}
}
