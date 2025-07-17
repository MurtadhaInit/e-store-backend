package handlers

import (
	"crypto/rand"
	"crypto/sha256"
	"e-store-backend/internal/repository"
	"e-store-backend/internal/utils"
	"encoding/base32"
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type token struct {
	repository.AddTokenParams
	TokenPlaintext string `json:"token"`
}

func generateToken(customerID int32, ttl time.Duration, scope string) (*token, error) {
	emptyBytes := make([]byte, 32)
	_, err := rand.Read(emptyBytes)
	if err != nil {
		return nil, err
	}
	plaintext := base32.StdEncoding.WithPadding(base32.NoPadding).EncodeToString(emptyBytes)
	hash := sha256.Sum256([]byte(plaintext))

	newToken := &token{
		AddTokenParams: repository.AddTokenParams{
			CustomerID: customerID,
			Expiry:     time.Now().Add(ttl),
			TokenScope: scope,
			TokenHash:  hash[:],
		},
		TokenPlaintext: plaintext,
	}

	return newToken, nil
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

func (h *Handlers) CreateToken(w http.ResponseWriter, r *http.Request) {
	type newTokenRequest struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	var newTokenReq newTokenRequest

	err := json.NewDecoder(r.Body).Decode(&newTokenReq)
	if err != nil {
		utils.WriteJSON(w, http.StatusBadRequest, utils.Envelope{"error": "invalid or missing data"})
		return
	}

	customer, err := h.queries.GetCustomerByUsername(r.Context(), newTokenReq.Username)
	if err != nil {
		h.serverError(w, r, err)
		return
	}

	passwordMatch, err := passwordMatch(newTokenReq.Password, customer.PasswordHash)
	if err != nil {
		h.serverError(w, r, err)
		return
	}

	if !passwordMatch {
		h.clientError(w, http.StatusUnauthorized, "invalid credentials")
	}

	token, err := generateToken(customer.CustomerID, 24*time.Hour, utils.Scopes.ScopeAuth)
	if err != nil {
		h.serverError(w, r, err)
		return
	}

	result, err := h.queries.AddToken(r.Context(), token.AddTokenParams)
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
		err = utils.WriteJSON(w, http.StatusCreated, utils.Envelope{"auth_token": token})
		if err != nil {
			h.serverError(w, r, err)
			return
		}
	} else {
		h.clientError(w, http.StatusBadRequest, "Not token has been generated")
	}
}
