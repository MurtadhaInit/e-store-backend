package handlers

import (
	"database/sql"
	"e-store-backend/internal/repository"
	"e-store-backend/internal/utils"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func (h *Handlers) ProductView(w http.ResponseWriter, r *http.Request) {
	id, err := utils.ReadIDParam(r)
	if err != nil {
		h.clientError(w, http.StatusBadRequest, err.Error())
		return
	}

	product, err := h.queries.GetProduct(r.Context(), id)
	if err != nil {
		h.clientError(w, http.StatusNotFound, "No matching record found")
		return
	}

	// TODO
	// category, err := h.queries.GetProductCategory(r.Context(), product.Category.Int32)
	// if err != nil {
	// }

	err = utils.WriteJSON(w, http.StatusFound, utils.Envelope{"product": product})
	if err != nil {
		h.serverError(w, r, err)
		return
	}
}

func (h *Handlers) ProductAdd(w http.ResponseWriter, r *http.Request) {
	var newProduct repository.AddProductParams

	err := json.NewDecoder(r.Body).Decode(&newProduct)
	if err != nil {
		h.clientError(w, http.StatusBadRequest, "Invalid or missing data")
		return
	}

	result, err := h.queries.AddProduct(r.Context(), newProduct)
	if err != nil {
		// TODO: add a check for duplicate errors and return a client error with a custom message
		h.serverError(w, r, err)
		return
	}

	insertedID, err := result.LastInsertId()
	if err != nil {
		h.serverError(w, r, err)
		return
	}

	insertedProduct, err := h.queries.GetProduct(r.Context(), int32(insertedID))
	if err != nil {
		h.serverError(w, r, err)
		return
	}

	err = utils.WriteJSON(w, http.StatusCreated, utils.Envelope{"product": insertedProduct})
	if err != nil {
		h.serverError(w, r, err)
		return
	}
}

func (h *Handlers) ProductDelete(w http.ResponseWriter, r *http.Request) {
	id, err := utils.ReadIDParam(r)
	if err != nil {
		h.clientError(w, http.StatusBadRequest)
		return
	}

	result, err := h.queries.DeleteProduct(r.Context(), id)
	if err != nil {
		h.clientError(w, http.StatusBadRequest, "No matching record found")
		return
	}

	rows, err := result.RowsAffected()
	if err != nil {
		h.serverError(w, r, err)
		return
	}

	if rows > 0 {
		w.WriteHeader(http.StatusNoContent)
		_, err := fmt.Fprintf(w, "Successfully removed %d product(s)", int(rows))
		if err != nil {
			h.serverError(w, r, err)
		}
	} else {
		h.clientError(w, http.StatusBadRequest, "No matching products removed")
	}
}

func (h *Handlers) ProductAll(w http.ResponseWriter, r *http.Request) {
	products, err := h.queries.GetAllProducts(r.Context())
	if err != nil {
		h.serverError(w, r, err)
		return
	}

	err = utils.WriteJSON(w, http.StatusOK, utils.Envelope{"products": products})
	if err != nil {
		h.serverError(w, r, err)
		return
	}
}

func (h *Handlers) ProductLatest(w http.ResponseWriter, r *http.Request) {
	limitPar := r.URL.Query().Get("limit")
	var limit int32
	if limitPar == "" {
		limit = 3 // default value
	} else {
		limitInt, err := strconv.Atoi(limitPar)
		if err != nil || limitInt < 1 {
			h.clientError(w, http.StatusBadRequest, "invalid limit provided")
			return
		}
		limit = int32(limitInt)
	}

	products, err := h.queries.GetLatestProducts(r.Context(), limit)
	if err != nil {
		h.serverError(w, r, err)
		return
	}

	err = utils.WriteJSON(w, http.StatusOK, utils.Envelope{"products": products})
	if err != nil {
		h.serverError(w, r, err)
		return
	}
}

func (h *Handlers) ProductEdit(w http.ResponseWriter, r *http.Request) {
	id, err := utils.ReadIDParam(r)
	if err != nil {
		h.clientError(w, http.StatusBadRequest, err.Error())
		return
	}

	var updateData map[string]any
	err = json.NewDecoder(r.Body).Decode(&updateData)
	if err != nil {
		h.clientError(w, http.StatusBadRequest, "Invalid JSON data")
		return
	}

	if len(updateData) == 0 {
		h.clientError(w, http.StatusBadRequest, "No fields to update")
		return
	}

	params := repository.EditProductParams{
		ProductID: id,
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

	result, err := h.queries.EditProduct(r.Context(), params)
	if err != nil {
		h.serverError(w, r, err)
		return
	}

	rows, err := result.RowsAffected()
	if err != nil {
		h.serverError(w, r, err)
		return
	}

	if rows > 0 {
		updatedProduct, err := h.queries.GetProduct(r.Context(), id)
		if err != nil {
			h.serverError(w, r, err)
			return
		}

		err = utils.WriteJSON(w, http.StatusOK, utils.Envelope{"product": updatedProduct})
		if err != nil {
			h.serverError(w, r, err)
			return
		}
	} else {
		h.clientError(w, http.StatusBadRequest, "No rows were updated")
	}
}
