package handlers

import (
	"net/http"
	"runtime/debug"
)

func (h *Handlers) serverError(w http.ResponseWriter, r *http.Request, err error) {
	var (
		method = r.Method
		uri    = r.URL.RequestURI()
		trace  = string(debug.Stack())
	)

	h.logger.Error(err.Error(), "method", method, "uri", uri, "track", trace)
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func (h *Handlers) clientError(w http.ResponseWriter, status int, message ...string) {
	msg := http.StatusText(status)
	if len(message) > 0 && message[0] != "" {
		msg = message[0]
	}
	http.Error(w, msg, status)
}
