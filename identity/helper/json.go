package helper

import (
	"encoding/json"
	"net/http"
)

func WriteJSON(w http.ResponseWriter, code int, v any) error {
	w.Header().Add("content-type", "application/json")
	w.WriteHeader(code)
	return json.NewEncoder(w).Encode(v)
}

type Result struct {
	Value     any   `json:"value,omitempty"`
	IsSuccess bool  `json:"is_success"`
	Error     error `json:"error,omitempty"`
}
