package main

import (
	"e-store-backend/internal/repository"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func (app *application) productView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 1 {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	product, err := app.queries.GetProduct(r.Context(), int32(id))
	if err != nil {
		app.clientError(w, http.StatusNotFound, "No matching record found")
		return
	}

	// TODO
	// category, err := app.queries.GetProductCategory(r.Context(), product.Category.Int32)
	// if err != nil {
	// }

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(product)
	if err != nil {
		app.serverError(w, r, err)
		return
	}
}

func (app *application) productAdd(w http.ResponseWriter, r *http.Request) {
	var newProduct repository.AddProductParams

	err := json.NewDecoder(r.Body).Decode(&newProduct)
	if err != nil {
		app.clientError(w, http.StatusBadRequest, "Invalid or missing data")
		return
	}

	result, err := app.queries.AddProduct(r.Context(), newProduct)
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

	insertedProduct, err := app.queries.GetProduct(r.Context(), int32(insertedID))
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(insertedProduct)
	if err != nil {
		app.serverError(w, r, err)
		return
	}
}

func (app *application) productDelete(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 1 {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	result, err := app.queries.DeleteProduct(r.Context(), int32(id))
	if err != nil {
		app.clientError(w, http.StatusBadRequest, "No matching record found")
		return
	}

	rows, err := result.RowsAffected()
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	if rows > 0 {
		_, err := fmt.Fprintf(w, "Successfully removed %d product(s)", int(rows))
		if err != nil {
			app.serverError(w, r, err)
		}
	} else {
		app.clientError(w, http.StatusBadRequest, "No matching products removed")
	}
}

func (app *application) productAll(w http.ResponseWriter, r *http.Request) {
	products, err := app.queries.GetAllProducts(r.Context())
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(products)
	if err != nil {
		app.serverError(w, r, err)
		return
	}
}

func (app *application) productLatest(w http.ResponseWriter, r *http.Request) {
	limitPar := r.URL.Query().Get("limit")
	var limit int32
	if limitPar == "" {
		limit = 10 // default value
	} else {
		limitInt, err := strconv.Atoi(limitPar)
		if err != nil || limitInt < 1 {
			app.clientError(w, http.StatusBadRequest, "invalid limit provided")
			return
		}
		limit = int32(limitInt)
	}

	products, err := app.queries.GetLatestProducts(r.Context(), limit)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(products)
	if err != nil {
		app.serverError(w, r, err)
		return
	}
}

func (app *application) productEdit(w http.ResponseWriter, r *http.Request) {
	// result, err := app.queries.EditProduct(r.Context(), repository.EditProductParams{
	// 	ProductID: ,
	// })
}
