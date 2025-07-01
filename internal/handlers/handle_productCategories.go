package handlers

import (
	"e-store-backend/internal/repository"
	"e-store-backend/internal/utils"
	"encoding/json"
	"net/http"
)

func (h *Handlers) CategoryAdd(w http.ResponseWriter, r *http.Request) {
	var newProductCategory repository.AddProductCategoryParams

	err := json.NewDecoder(r.Body).Decode(&newProductCategory)
	if err != nil {
		h.clientError(w, http.StatusBadRequest, "Invalid or missing data")
		return
	}

	result, err := h.queries.AddProductCategory(r.Context(), newProductCategory)
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

	insertedProductCategory, err := h.queries.GetProductCategory(r.Context(), int32(insertedID))
	if err != nil {
		h.serverError(w, r, err)
		return
	}

	err = utils.WriteJSON(w, http.StatusCreated, utils.Envelope{"category": insertedProductCategory})
	if err != nil {
		h.serverError(w, r, err)
		return
	}
}

func (h *Handlers) CategoryAll(w http.ResponseWriter, r *http.Request) {
	productCategories, err := h.queries.GetAllProductCategories(r.Context())
	if err != nil {
		h.serverError(w, r, err)
		return
	}

	err = utils.WriteJSON(w, http.StatusOK, utils.Envelope{"categories": productCategories})
	if err != nil {
		h.serverError(w, r, err)
		return
	}
}
