package routes

import (
	"e-store-backend/internal/app"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func SetupRoutes(app *app.Application) *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.StripSlashes)
	// TODO: configure properly for production
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"http://localhost:5173"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
	}))

	r.Group(func(r chi.Router) {
		r.Use(app.Middleware.Authenticate)
		r.Use(app.Middleware.RequireCustomer)

		r.Delete("/products/{id}", app.Handlers.ProductDelete)
		r.Patch("/products/{id}", app.Handlers.ProductEdit)
		r.Post("/products", app.Handlers.ProductAdd)

		r.Post("/categories/add", app.Handlers.CategoryAdd)
	})

	r.Get("/health", app.HealthCheck)

	// Product routes
	r.Get("/products", app.Handlers.ProductAll)
	r.Get("/products/latest", app.Handlers.ProductLatest)
	r.Get("/products/{id}", app.Handlers.ProductView)

	// Categories routes
	r.Get("/categories", app.Handlers.CategoryAll)

	// Customers routes
	r.Post("/customers", app.Handlers.AddCustomer)

	// Tokens routes
	r.Post("/tokens/authentication", app.Handlers.CreateToken)

	return r
}
