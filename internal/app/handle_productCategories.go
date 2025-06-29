package app

import (
	"e-store-backend/internal/repository"
	"encoding/json"
	"net/http"
)

func (app *Application) CategoryAdd(w http.ResponseWriter, r *http.Request) {
	var newProductCategory repository.AddProductCategoryParams

	err := json.NewDecoder(r.Body).Decode(&newProductCategory)
	if err != nil {
		app.clientError(w, http.StatusBadRequest, "Invalid or missing data")
		return
	}

	result, err := app.Queries.AddProductCategory(r.Context(), newProductCategory)
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

	insertedProductCategory, err := app.Queries.GetProductCategory(r.Context(), int32(insertedID))
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

func (app *Application) CategoryAll(w http.ResponseWriter, r *http.Request) {
	productCategories, err := app.Queries.GetAllProductCategories(r.Context())
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(productCategories)
	if err != nil {
		app.serverError(w, r, err)
		return
	}
}
