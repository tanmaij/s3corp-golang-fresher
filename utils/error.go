package utils

import (
	"encoding/json"
	"net/http"
)

type Error struct {
	Message    string
	StatusCode int
}

func (error Error) Response(w http.ResponseWriter) {
	(w).WriteHeader(error.StatusCode)
	json.NewEncoder(w).Encode(Response{false, error.Message})
}
func NewError(message string, statusCode int) Error {
	return Error{message, statusCode}
}
