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
func (error Error) Error() string {
	return error.Message
}

const (
	InvalidData         = "invalid Data"
	UserAlreadyExist    = "user already exist"
	PasswordIsIncorrect = "password is incorrect"
	Successfully        = "Successfully"
	InternalServerError = "internal Server Error"
	NotExist            = "not exist"
	QueryDataIncorrect  = "Query data is incorrect"
	NotFound            = "not found"
	InvalidEmail        = "invalid Email"
	UsernameIsNotFound  = "username is not found"
	PasswordIsNotFound  = "username is not found"
	EmailIsNotFound     = "email is not found"
	NameIsNotFound      = "name is not found"
)
