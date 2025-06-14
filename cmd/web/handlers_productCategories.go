package main

import (
	"e-store-backend/internal/repository"
	"encoding/json"
	"net/http"
)

func (app *application) categoryAdd(w http.ResponseWriter, r *http.Request) {
	var newProductCategory repository.AddProductCategoryParams

	err := json.NewDecoder(r.Body).Decode(&newProductCategory)
	if err != nil {
		app.clientError(w, http.StatusBadRequest, "Invalid or missing data")
		return
	}

	result, err := app.queries.AddProductCategory(r.Context(), newProductCategory)
	if err != nil {
		// TODO: add a check for duplicate errors and return a client error with a custom message
		app.serverError(w, r, err)
		return
	}

	insertedID, err := result.LastInsertId()
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	insertedProductCategory, err := app.queries.GetProductCategory(r.Context(), int32(insertedID))
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(insertedProductCategory)
	if err != nil {
		app.serverError(w, r, err)
		return
	}
}
