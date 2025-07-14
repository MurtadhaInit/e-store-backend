package handlers

import (
	"e-store-backend/internal/repository"
	"e-store-backend/internal/utils"
	"encoding/json"
	"errors"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

type newCustomer struct {
	repository.AddCustomerParams
	Password string `json:"password"`
}

func (newCustomer *newCustomer) validate() error {
	switch {
	case len(newCustomer.Password) < 10:
		return errors.New("password should be 10 characters or more")
	case len(newCustomer.Password) > 72:
		return errors.New("password should not be more than 72 characters")
	case newCustomer.Username == "":
		return errors.New("username is required")
	}

	return nil
}

func (newCustomer *newCustomer) setPassword() error {
	hash, err := bcrypt.GenerateFromPassword([]byte(newCustomer.Password), bcrypt.DefaultCost+2)
	if err != nil {
		return err
	}
	newCustomer.PasswordHash = hash
	return nil
}

func passwordMatch(clearTextPass string, hash []byte) (bool, error) {
	if err := bcrypt.CompareHashAndPassword(hash, []byte(clearTextPass)); err != nil {
		switch {
		case errors.Is(err, bcrypt.ErrMismatchedHashAndPassword):
			return false, nil
		default:
			return false, err // internal server error
		}
	}

	return true, nil
}

func (h *Handlers) AddCustomer(w http.ResponseWriter, r *http.Request) {
	var newCustomer newCustomer

	err := json.NewDecoder(r.Body).Decode(&newCustomer)
	if err != nil {
		utils.WriteJSON(w, http.StatusBadRequest, utils.Envelope{"error": "invalid or missing data"})
		return
	}

	if err := newCustomer.validate(); err != nil {
		utils.WriteJSON(w, http.StatusBadRequest, utils.Envelope{"error": err.Error()})
		return
	}

	if err := newCustomer.setPassword(); err != nil {
		h.serverError(w, r, err)
		return
	}

	params := repository.AddCustomerParams{
		Username:     newCustomer.Username,
		PasswordHash: newCustomer.PasswordHash,
		FirstName:    newCustomer.FirstName,
		LastName:     newCustomer.LastName,
		Email:        newCustomer.Email,
		BirthDay:     newCustomer.BirthDay,
		PhoneNumber:  newCustomer.PhoneNumber,
		Address:      newCustomer.Address,
	}

	result, err := h.queries.AddCustomer(r.Context(), params)
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

	insertedCustomer, err := h.queries.GetCustomerByID(r.Context(), int32(insertedID))
	if err != nil {
		h.serverError(w, r, err)
		return
	}

	err = utils.WriteJSON(w, http.StatusCreated, utils.Envelope{"customer": insertedCustomer})
	if err != nil {
		h.serverError(w, r, err)
		return
	}
}

func (h *Handlers) GetCustomerByUsername(w http.ResponseWriter, r *http.Request) {
	username := r.PathValue("username")
	if username == "" {
		h.clientError(w, http.StatusBadRequest, "missing or invalid username")
		return
	}

	customer, err := h.queries.GetCustomerByUsername(r.Context(), username)
	if err != nil {
		h.clientError(w, http.StatusNotFound, "No matching record found")
		return
	}

	err = utils.WriteJSON(w, http.StatusFound, utils.Envelope{"customer": customer})
	if err != nil {
		h.serverError(w, r, err)
		return
	}
}

func (h *Handlers) UpdateCustomer(w http.ResponseWriter, r *http.Request) {

}
