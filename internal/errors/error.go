package errors

import (
	"fmt"
	"net/http"
)

type Error struct {
	Message    string `json:"message"`
	StatusCode int    `json:"statusCode"`
}

func (error Error) Response(w http.ResponseWriter) {
	w.WriteHeader(error.StatusCode)
	fmt.Fprint(w, error.Message)
}
func NewError(message string, statusCode int) Error {
	return Error{message, statusCode}
}

const (
	UserAlreadyExist    = "user already exist"
	PasswordIsIncorrect = "password is incorrect"
	Successfully        = "Successfully"
	InternalServerError = "internal Server Error"
	NotExist            = "Not exist"
	QueryDataIncorrect  = "Query data is incorrect"
	NotFound            = "not Found"
)
