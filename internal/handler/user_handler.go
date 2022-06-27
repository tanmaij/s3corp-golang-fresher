package handler

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"net/http"
	"s3corp-golang-fresher/internal/models"
	"s3corp-golang-fresher/internal/service"
	"s3corp-golang-fresher/utils"
	"strings"
)

type UserHandler struct {
	UserService service.UserService
}

// NewUserHandler Return new user_handler with user_service parameter
func NewUserHandler(userService service.UserService) UserHandler {
	return UserHandler{UserService: userService}
}

func (userHandler UserHandler) UserHandler(r chi.Router) {

	r.Post("/login", userHandler.Login)

	r.Get("/", userHandler.GetUsers)

	r.Get("/{username}", userHandler.GetUserByUsername)

	r.Put("/{id}", userHandler.UpdateUser)

	r.Delete("/{id}", userHandler.DeleteUser)

	r.Post("/", userHandler.CreateUser)
}

func (userHandler UserHandler) Login(w http.ResponseWriter, r *http.Request) {

	// 1. Define a variable name requestBody
	// Decode data from r.Body to the variable
	requestBody := make(map[string]interface{})
	json.NewDecoder(r.Body).Decode(&requestBody)

	// 2. Checking username and password variable from request
	// If not true, response status 400, and message
	username, ok := requestBody["username"]
	if !ok {
		http.Error(w, "Username is NOT FOUND", http.StatusBadRequest)
		return
	}
	password, ok := requestBody["password"]
	if !ok {
		http.Error(w, "Password is NOT FOUND", http.StatusBadRequest)
		return
	}

	// 3. Call Login method from UserService, get user information, token and error
	// if error, response error which returned by service
	user, token, err := userHandler.UserService.Login(username.(string), password.(string))
	if err.StatusCode != http.StatusOK {
		err.Response(w)
		return
	}

	// 4. If not error, response user information and token
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(utils.Response{Success: true, Data: map[string]any{"user": user, "token": token}})
}

func (userHandler UserHandler) GetUserByUsername(w http.ResponseWriter, r *http.Request) {
	username := chi.URLParam(r, "username")

	user, err := userHandler.UserService.GetUserByUsername(username)

	if err.StatusCode != http.StatusOK {
		err.Response(w)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(utils.Response{Success: true, Data: user})
}

func (userHandler UserHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
	users, pagination, err := userHandler.UserService.GetUsers(map[string]string{})

	if err.StatusCode != http.StatusOK {
		err.Response(w)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(utils.Response{Success: true, Data: map[string]any{"users": users, "pagination": pagination}})
}

func (userHandler UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {

	// 1. Define a variable name requestBody
	// Decode data from r.Body to the variable
	requestBody := make(map[string]interface{})
	json.NewDecoder(r.Body).Decode(&requestBody)

	// 2. Checking username,name,email and password variable from request
	// If not true, response status 400, and message
	username, ok := requestBody["username"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Username is not found")
		return
	}
	password, ok := requestBody["password"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Password is not found")
		return
	}
	name, ok := requestBody["name"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Name is not found")
		return
	}
	email, ok := requestBody["email"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Email is not found")
		return
	}

	// 2. Checking email variable from request
	// If it doesn't contain '@' character, response status 400, and message
	// Or if it contains more than 2 '@' character, response status 400, and message
	if !strings.Contains(email.(string), "@") || strings.Count(email.(string), "@") > 1 {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Email is not correct")
		return
	}

	// 3. Define new user to save
	newUser := models.User{Password: password.(string),
		Username: username.(string),
		Email:    email.(string),
		Name:     name.(string)}

	// 4. Call createUser method from user service with user which is just defined
	// If service return any error , response the error to client
	err := userHandler.UserService.CreateUser(newUser)
	if err.StatusCode != http.StatusOK {
		err.Response(w)
		return
	}

	// 4. If not error, response successfully message
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(utils.Response{Success: true, Data: "Create user successfully"})
}

func (userHandler UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {

	username := chi.URLParam(r, "username")
	user := make(map[string]interface{})
	json.NewDecoder(r.Body).Decode(&user)

	password, ok := user["password"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(utils.Response{Data: "Password is NOT FOUND"})
		return
	}
	name, ok := user["name"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(utils.Response{Data: "Name is NOT FOUND"})
		return
	}
	email, ok := user["email"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(utils.Response{Data: "Email is NOT FOUND"})
		return
	}
	if !strings.Contains(email.(string), "@") || strings.Count(email.(string), "@") > 1 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(utils.Response{Data: "Email is NOT CORRECT"})
		return
	}
	newUser := models.User{
		Password: password.(string),
		Username: username,
		Email:    email.(string),
		Name:     name.(string)}
	err := userHandler.UserService.UpdateUser(newUser)

	if err.StatusCode != http.StatusOK {
		err.Response(w)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(utils.Response{Success: true, Data: "Update user successfully"})
}

func (userHandler UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {

	username := chi.URLParam(r, "username")

	err := userHandler.UserService.DeleteUser(username)

	if err.StatusCode != http.StatusOK {
		err.Response(w)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(utils.Response{Success: true, Data: "Update user successfully"})

}
