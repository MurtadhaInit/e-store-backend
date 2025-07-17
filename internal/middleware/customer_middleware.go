package middleware

import (
	"context"
	"crypto/sha256"
	"database/sql"
	"e-store-backend/internal/repository"
	"e-store-backend/internal/utils"
	"errors"
	"net/http"
	"strings"
	"time"
)

type contextKey string

const customerContextKey = contextKey("customer")

var anonymousCustomer = &repository.Customer{}

func setCustomer(r *http.Request, customer *repository.Customer) *http.Request {
	ctx := context.WithValue(r.Context(), customerContextKey, customer)
	return r.WithContext(ctx)
}

func GetCustomer(r *http.Request) *repository.Customer {
	customer, ok := r.Context().Value(customerContextKey).(*repository.Customer)
	if !ok {
		panic("missing customer in request") // possible bad actor call
	}
	return customer
}

func (m *Middleware) Authenticate(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Vary", "Authorization")
		authHeader := r.Header.Get("Authorization")

		if authHeader == "" {
			r = setCustomer(r, anonymousCustomer)
			next.ServeHTTP(w, r)
			return
		}

		headerParts := strings.Split(authHeader, " ")
		if len(headerParts) != 2 || headerParts[0] != "Bearer" {
			utils.WriteJSON(w, http.StatusUnauthorized, utils.Envelope{"error": "invalid authorization header"})
			return
		}

		token := headerParts[1]
		tokenHash := sha256.Sum256([]byte(token))

		customer, err := m.queries.GetCustomerByToken(r.Context(), repository.GetCustomerByTokenParams{
			TokenHash:  tokenHash[:],
			TokenScope: utils.Scopes.ScopeAuth,
			Date:       time.Now(),
		})
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				utils.WriteJSON(w, http.StatusUnauthorized, utils.Envelope{"error": "token expired or invalid"})
				return
			}
			utils.WriteJSON(w, http.StatusUnauthorized, utils.Envelope{"error": "invalid token"})
			return
		}

		r = setCustomer(r, &customer)
		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}

func (m *Middleware) RequireCustomer(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		customer := GetCustomer(r)
		if customer == anonymousCustomer {
			utils.WriteJSON(w, http.StatusUnauthorized, utils.Envelope{"error": "you must be logged in to access this route"})
			return
		}

		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}
