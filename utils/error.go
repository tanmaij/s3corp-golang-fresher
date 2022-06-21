package utils

import (
	"encoding/json"
	"net/http"
)

type Error struct {
	Message    string
	StatusCode int
}

func (error Error) Error() string {
	return error.Message
}
func (error Error) Response(w http.ResponseWriter) {
	w.WriteHeader(error.StatusCode)
	json.NewEncoder(w).Encode(Response{false, error.Message})
}
