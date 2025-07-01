package utils

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
)

type Envelope map[string]any

func WriteJSON(w http.ResponseWriter, status int, data Envelope) error {
	js, err := json.MarshalIndent(data, "", "	")
	if err != nil {
		return err
	}

	js = append(js, '\n')

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_, err = w.Write(js)
	if err != nil {
		return err
	}
	return nil
}

func ReadIDParam(r *http.Request) (int32, error) {
	idParam := r.PathValue("id")
	if idParam == "" {
		return 0, errors.New("Invalid id parameter")
	}

	id, err := strconv.ParseInt(idParam, 10, 32)
	if err != nil {
		return 0, errors.New("Invalid id parameter type")
	}

	if id < 1 {
		return 0, errors.New("Invalid id parameter")
	}

	return int32(id), nil
}
