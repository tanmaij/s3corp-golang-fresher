package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"s3corp-golang-fresher/internal/errors"
	"s3corp-golang-fresher/internal/models"
	"s3corp-golang-fresher/internal/roles"
	"s3corp-golang-fresher/internal/service"
	"s3corp-golang-fresher/pkg"
	"strconv"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth"
)

type UserHandler struct {
	UserService service.UserService
}

// NewUserHandler Return new user_handler with user_service parameter
func NewUserHandler(userService service.UserService) UserHandler {
	return UserHandler{UserService: userService}
}

var jwtAuth = pkg.GetJWTAuth()

func (userHandler UserHandler) UserHandler(r chi.Router) {

	// Protected routes
	r.Group(func(r chi.Router) {
		// Seek, verify and validate JWT tokens
		r.Use(jwtauth.Verifier(jwtAuth))

		// Handle valid / invalid tokens. In this example, we use
		// the provided authenticator middleware, but you can write your
		// own very easily, look at the Authenticator method in jwtauth.go
		// and tweak it, its not scary.
		r.Use(jwtauth.Authenticator)

		r.Put("/{id}", userHandler.UpdateUser)

		r.Delete("/{id}", userHandler.DeleteUser)

	})
	r.Group(func(r chi.Router) {

		r.Get("/", userHandler.GetUsers)

		r.Post("/login", userHandler.Login)

		r.Post("/", userHandler.CreateUser)
	})
}

func (userHandler UserHandler) Login(w http.ResponseWriter, r *http.Request) {

	// 1. Define a variable name requestBody
	// Decode data from r.Body to the variable
	requestBody := make(map[string]interface{})
	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, errors.InvalidData)
		return
	}

	// 2. Checking username and password variable from request
	// If not true, response status 400, and message
	username, ok := requestBody["username"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, errors.UsernameIsNotFound)
		return
	}
	password, ok := requestBody["password"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, errors.PasswordIsNotFound)
		return
	}

	// 3. Call Login method from UserService, get user information, token and error
	// if error, response error which returned by service
	user, token, err := userHandler.UserService.Login(username.(string), password.(string))
	if err != nil {
		err.(errors.Error).Response(w)
		return
	}

	// 4. If not error, response user information and token
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(map[string]any{"user": user, "token": token}); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, errors.InternalServerError)
		return
	}
}

func (userHandler UserHandler) GetUsers(w http.ResponseWriter, r *http.Request) {

	// 1. Get limit and page variable from url Query
	limit := r.URL.Query().Get("limit")
	page := r.URL.Query().Get("page")

	//Check limit and page is unsigned int
	_limit, err := strconv.Atoi(limit)
	_page, err := strconv.Atoi(page)
	if err != nil || _limit <= 0 || _page <= 0 {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, errors.InvalidData)
		return
	}

	// 2. Call get users method from user service with limit and page variable
	users, pagination, err := userHandler.UserService.GetUsers(map[string]int{"limit": _limit, "page": _page})

	// if any error, response to client
	if err != nil {
		err.(errors.Error).Response(w)
		return
	}

	// 3. If not error, response data that include user and pagination
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(map[string]any{"users": users, "pagination": pagination}); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, errors.InternalServerError)
		return
	}
}

func (userHandler UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {

	// 1. Define a variable name requestBody
	// Decode data from r.Body to the variable
	requestBody := make(map[string]interface{})
	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, errors.InternalServerError)
		return
	}

	// 2. Checking username,name,email and password variable from request
	// If not true, response status 400, and message
	username, ok := requestBody["username"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, errors.UsernameIsNotFound)
		return
	}
	password, ok := requestBody["password"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, errors.PasswordIsNotFound)
		return
	}
	name, ok := requestBody["name"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, errors.NameIsNotFound)
		return
	}
	email, ok := requestBody["email"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, errors.EmailIsNotFound)
		return
	}

	// 2. Checking email variable from request
	// If it doesn't contain '@' character, response status 400, and message
	// Or if it contains more than 2 '@' character, response status 400, and message
	if !strings.Contains(email.(string), "@") || strings.Count(email.(string), "@") > 1 {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, errors.InvalidEmail)
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
	if err != nil {
		err.(errors.Error).Response(w)
		return
	}

	// 4. If not error, response successfully message
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Create user successfully"))
}

func (userHandler UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {

	//Get username from url params
	username := chi.URLParam(r, "username")

	// Get body data
	body := make(map[string]interface{})
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, errors.InvalidData)
		return
	}

	// Authenticate
	// Get user data from context
	token, claims, err := jwtauth.FromContext(r.Context())
	if err != nil || token == nil || claims == nil {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, errors.InvalidToken)
		return
	}
	// Get role to authorization
	role, ok := claims["role"].(string)
	if !ok {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, errors.InvalidToken)
		return
	}
	// get author username of the request
	author, ok := claims["username"].(string)
	if !ok {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, errors.InvalidToken)
		return
	}

	// Authorization role or username
	if role != roles.Admin && author != username {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, errors.PermissionDenied)
		return
	}

	//Get user data from body
	name, ok := body["name"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Name is NOT FOUND")
		return
	}
	email, ok := body["email"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Email is NOT FOUND")
		return
	}
	if !strings.Contains(email.(string), "@") || strings.Count(email.(string), "@") > 1 {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, errors.InvalidEmail)
		return
	}
	user := models.User{
		Username: username,
		Email:    email.(string),
		Name:     name.(string)}

	err2 := userHandler.UserService.UpdateUser(user)

	if err2 != nil {
		err2.(errors.Error).Response(w)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Update user successfully"))
}

func (userHandler UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	// Get username from url
	username := chi.URLParam(r, "username")
	if username == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, errors.InvalidData)
		return
	}
	// Get user data from context
	token, claims, err := jwtauth.FromContext(r.Context())
	if err != nil || token == nil || claims == nil {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, errors.InvalidToken)
		return
	}
	// Get role to authorization
	role, ok := claims["role"].(string)
	if !ok {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, errors.InvalidToken)
		return
	}
	// get author username of the request
	author, ok := claims["username"].(string)
	if !ok {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, errors.InvalidToken)
		return
	}

	// Authorization role or username
	if role != roles.Admin && author != username {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, errors.PermissionDenied)
		return
	}
	// If there are not any error, delete user
	err2 := userHandler.UserService.DeleteUser(username)

	if err2 != nil {
		err2.(errors.Error).Response(w)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Delete user successfully"))
}
