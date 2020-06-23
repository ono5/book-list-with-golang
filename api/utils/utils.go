package utils

import (
	"encoding/json"
	"net/http"

	"github.com/ono5/book-list/api/models"
)

// SendError - send error to user
func SendError(w http.ResponseWriter, status int, err models.Error) {
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(err)
}

// SendSuccess - send success to user
func SendSuccess(w http.ResponseWriter, data interface{}) {
	json.NewEncoder(w).Encode(data)
}
