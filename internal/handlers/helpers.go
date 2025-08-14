package handlers

import (
	"e-store-backend/internal/utils"
	"net/http"
	"runtime/debug"
)

// TODO: write function docs
func (h *Handlers) serverError(w http.ResponseWriter, r *http.Request, err error) {
	var (
		method = r.Method
		uri    = r.URL.RequestURI()
		trace  = string(debug.Stack())
	)

	h.logger.Error(err.Error(), "method", method, "uri", uri, "track", trace)
	err = utils.WriteJSON(w, http.StatusInternalServerError, utils.Envelope{"error": http.StatusText(http.StatusInternalServerError)})
	if err != nil {
		h.logger.Error(err.Error())
	}
}

// TODO: might remove. It can be replaced by just using writeJSON alone.
func (h *Handlers) clientError(w http.ResponseWriter, status int, message ...string) {
	msg := http.StatusText(status)
	if len(message) > 0 && message[0] != "" {
		msg = message[0]
	}
	err := utils.WriteJSON(w, status, utils.Envelope{"error": msg})
	if err != nil {
		h.logger.Error(err.Error())
	}
}
