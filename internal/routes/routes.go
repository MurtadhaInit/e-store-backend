package routes

import (
	"e-store-backend/internal/app"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func SetupRoutes(app *app.Application) *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.StripSlashes)

	r.Get("/health", app.HealthCheck)

	// Product routes
	r.Get("/products", app.Handlers.ProductAll)
	r.Get("/products/latest", app.Handlers.ProductLatest)
	r.Get("/products/{id}", app.Handlers.ProductView)
	r.Delete("/products/{id}", app.Handlers.ProductDelete)
	r.Patch("/products/{id}", app.Handlers.ProductEdit)
	r.Post("/products/add", app.Handlers.ProductAdd)

	// Categories routes
	r.Get("/categories", app.Handlers.CategoryAll)
	r.Post("/categories/add", app.Handlers.CategoryAdd)

	return r
}
