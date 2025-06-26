package routes

import (
	"e-store-backend/internal/app"

	"github.com/go-chi/chi/v5"
)

func SetupRoutes(app *app.Application) *chi.Mux {
	r := chi.NewRouter()

	r.Get("/health", app.HealthCheck)

	// Product routes
	r.Get("/products/", app.ProductAll)
	r.Get("/products/latest", app.ProductLatest)
	r.Get("/products/{id}", app.ProductView)
	r.Delete("/products/{id}", app.ProductDelete)
	r.Patch("/products/{id}", app.ProductEdit)
	r.Post("/products/add", app.ProductAdd)

	// Categories routes
	r.Post("/categories/add", app.CategoryAdd)

	return r
}
