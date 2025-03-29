package main

import "net/http"

func (app *application) routes() *http.ServeMux {
	mux := http.NewServeMux()

	// Product routes
	mux.HandleFunc("GET /products/view/{id}", app.productView)
	mux.HandleFunc("POST /products/create", app.productCreate)
	mux.HandleFunc("POST /products/edit/{id}", app.productEdit)
	mux.HandleFunc("GET /products/delete/{id}", app.productDelete)
	mux.HandleFunc("GET /products/latest", app.productLatest)

	return mux
}
