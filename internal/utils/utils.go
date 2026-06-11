package utils

import (
	"encoding/json"
	"net/http"
)

// helper function to write JSON responses
func WriteJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(v)
}