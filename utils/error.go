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

const (
	UserAlreadyExist    = "User already exist"
	PasswordIsIncorrect = "Password is incorrect"
	Successfully        = "Successfully"
	InternalServerError = "Internal Server Error"
	NotExist            = "Not exist"
)
