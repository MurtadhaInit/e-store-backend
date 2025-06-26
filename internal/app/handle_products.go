package app

import (
	"database/sql"
	"e-store-backend/internal/repository"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func (app *Application) ProductView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 1 {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	product, err := app.Queries.GetProduct(r.Context(), int32(id))
	if err != nil {
		app.clientError(w, http.StatusNotFound, "No matching record found")
		return
	}

	// TODO
	// category, err := app.Queries.GetProductCategory(r.Context(), product.Category.Int32)
	// if err != nil {
	// }

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(product)
	if err != nil {
		app.serverError(w, r, err)
		return
	}
}

func (app *Application) ProductAdd(w http.ResponseWriter, r *http.Request) {
	var newProduct repository.AddProductParams

	err := json.NewDecoder(r.Body).Decode(&newProduct)
	if err != nil {
		app.clientError(w, http.StatusBadRequest, "Invalid or missing data")
		return
	}

	result, err := app.Queries.AddProduct(r.Context(), newProduct)
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

	insertedProduct, err := app.Queries.GetProduct(r.Context(), int32(insertedID))
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

func (app *Application) ProductDelete(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 1 {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	result, err := app.Queries.DeleteProduct(r.Context(), int32(id))
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

func (app *Application) ProductAll(w http.ResponseWriter, r *http.Request) {
	products, err := app.Queries.GetAllProducts(r.Context())
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

func (app *Application) ProductLatest(w http.ResponseWriter, r *http.Request) {
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

	products, err := app.Queries.GetLatestProducts(r.Context(), limit)
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

func (app *Application) ProductEdit(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 1 {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	var updateData map[string]interface{}
	err = json.NewDecoder(r.Body).Decode(&updateData)
	if err != nil {
		app.clientError(w, http.StatusBadRequest, "Invalid JSON data")
		return
	}

	if len(updateData) == 0 {
		app.clientError(w, http.StatusBadRequest, "No fields to update")
		return
	}

	params := repository.EditProductParams{
		ProductID: int32(id),
	}

	if title, ok := updateData["title"].(string); ok {
		params.Title = sql.NullString{String: title, Valid: true}
	}
	if desc, ok := updateData["product_description"].(string); ok {
		params.ProductDescription = sql.NullString{String: desc, Valid: true}
	}
	if price, ok := updateData["price"].(string); ok {
		params.Price = sql.NullString{String: price, Valid: true}
	}
	if imageURL, ok := updateData["image_url"].(string); ok {
		params.ImageUrl = sql.NullString{String: imageURL, Valid: true}
	}
	if category, ok := updateData["category"]; ok {
		if catFloat, isFloat := category.(float64); isFloat {
			params.Category = sql.NullInt32{Int32: int32(catFloat), Valid: true}
		}
	}
	if stock, ok := updateData["stock_quantity"]; ok {
		if stockFloat, isFloat := stock.(float64); isFloat {
			params.StockQuantity = sql.NullInt32{Int32: int32(stockFloat), Valid: true}
		}
	}

	result, err := app.Queries.EditProduct(r.Context(), params)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	rows, err := result.RowsAffected()
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	if rows > 0 {
		updatedProduct, err := app.Queries.GetProduct(r.Context(), int32(id))
		if err != nil {
			app.serverError(w, r, err)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		err = json.NewEncoder(w).Encode(updatedProduct)
		if err != nil {
			app.serverError(w, r, err)
			return
		}
	} else {
		app.clientError(w, http.StatusBadRequest, "No rows were updated")
	}
}
