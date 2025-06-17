package main

import "net/http"

func (app *application) routes() *http.ServeMux {
	mux := http.NewServeMux()

	// Product routes
	mux.HandleFunc("GET /products/{$}", app.productAll)
	mux.HandleFunc("GET /products/latest", app.productLatest)
	mux.HandleFunc("GET /products/{id}", app.productView)
	mux.HandleFunc("DELETE /products/{id}", app.productDelete)
	mux.HandleFunc("PATCH /products/{id}", app.productEdit)
	mux.HandleFunc("POST /products/add", app.productAdd)

	// Categories routes
	mux.HandleFunc("POST /categories/add", app.categoryAdd)

	return mux
}
