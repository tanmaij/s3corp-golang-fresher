package handler

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/volatiletech/null/v8"
	"net/http"
	"s3corp-golang-fresher/internal/models"
	"s3corp-golang-fresher/internal/service"
	"s3corp-golang-fresher/utils"
	"strings"
)

type UserHandler struct {
	UserService service.UserService
}

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

	user := make(map[string]interface{})
	json.NewDecoder(r.Body).Decode(&user)
	username, ok := user["username"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(utils.Response{Data: "Username is NOT FOUND"})
		return
	}
	password, ok := user["password"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(utils.Response{Data: "Password is NOT FOUND"})
		return
	}

	verify, err := userHandler.UserService.Login(username.(string), password.(string))
	if err.StatusCode != http.StatusOK {
		err.Response(w)
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(utils.Response{Success: true, Data: verify})
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
	users, err := userHandler.UserService.GetUsers()

	if err.StatusCode != http.StatusOK {
		err.Response(w)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(utils.Response{Success: true, Data: users})
}

func (userHandler UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	user := make(map[string]interface{})
	json.NewDecoder(r.Body).Decode(&user)

	username, ok := user["username"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(utils.Response{Data: "Username is NOT FOUND"})
		return
	}
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
	newUser := models.User{Password: null.String{String: password.(string), Valid: true},
		Username: username.(string),
		Email:    null.String{String: email.(string), Valid: true},
		Name:     null.String{String: name.(string), Valid: true}}
	err := userHandler.UserService.CreateUser(newUser)

	if err.StatusCode != http.StatusOK {
		err.Response(w)
		return
	}

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
		Password: null.String{String: password.(string), Valid: true},
		Username: username,
		Email:    null.String{String: email.(string), Valid: true},
		Name:     null.String{String: name.(string), Valid: true}}
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
