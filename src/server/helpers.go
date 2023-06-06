package server

import (
	"encoding/json"
	"net/http"
)

func WriteJson(w http.ResponseWriter, status int, v any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}

type serveFunc func(http.ResponseWriter, *http.Request) error

type ServeError struct {
	Error string `json:"error"`
}

type ReponseMap map[string]any

func handleFuncWrapper(fn serveFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := fn(w, r); err != nil {
			WriteJson(w, http.StatusBadRequest, ServeError{err.Error()})
		}
	}
}
