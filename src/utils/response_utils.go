package utils

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/aivot-digital/preppy-backend/src/models"
)

func WriteError(w http.ResponseWriter, status int, message string, args ...any) error {
	err := models.Error{
		Status:  status,
		Message: fmt.Sprintf(message, args...),
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(err)
}

func WriteBadRequest(w http.ResponseWriter) {
	WriteError(w, http.StatusBadRequest, "Bad request - Submitted object did not match required schema")
}

func WriteUnauthorized(w http.ResponseWriter) {
	WriteError(w, http.StatusUnauthorized, "Unauthorized - No valid access token or expired access token submitted")
}

func WriteNotFound(w http.ResponseWriter) {
	WriteError(w, http.StatusNotFound, "Not found - The requested resource could not be found")
}
