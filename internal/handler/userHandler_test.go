package handler

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/volatiletech/null/v8"
	"net/http"
	"net/http/httptest"
	"s3corp-golang-fresher/internal/models"
	"s3corp-golang-fresher/internal/service/mocks"
	"s3corp-golang-fresher/utils"
	"testing"
)

var data = [...]models.User{
	models.User{
		Password: null.String{String: "1", Valid: true},
		Username: "mai",
		Email:    null.String{String: "mai@gmail.com", Valid: true},
		Name:     null.String{String: "Mãi", Valid: true}},
	models.User{
		Password: null.String{String: "1", Valid: true},
		Username: "loc",
		Email:    null.String{String: "loc@gmail.com", Valid: true},
		Name:     null.String{String: "Loc", Valid: true}},
	models.User{
		Password: null.String{String: "1", Valid: true},
		Username: "nguyen",
		Email:    null.String{String: "nguyen@gmail.com", Valid: true},
		Name:     null.String{String: "Nguyên", Valid: true}},
	models.User{
		Password: null.String{String: "1", Valid: true},
		Username: "trung",
		Email:    null.String{String: "trung@gmail.com", Valid: true},
		Name:     null.String{String: "Trung", Valid: true}},
	models.User{
		Password: null.String{String: "1", Valid: true},
		Username: "tai",
		Email:    null.String{String: "tai@gmail.com", Valid: true},
		Name:     null.String{String: "Tai", Valid: true}},
	models.User{
		Password: null.String{String: "1", Valid: true},
		Username: "duy",
		Email:    null.String{String: "duy@gmail.com", Valid: true},
		Name:     null.String{String: "Duy", Valid: true}}}

const url = "/api/user/"

var userHandler UserHandler

type urlParams struct {
	key   string
	value string
}

func sendRequest(body any, methodName string, url string, handlerParam http.HandlerFunc, urlParams ...urlParams) *httptest.ResponseRecorder {
	r, _ := json.Marshal(body)
	json.Marshal(body)
	req, err := http.NewRequest(methodName, url, bytes.NewBuffer(r))
	req.Header.Set("Content-type", "application/json")
	for _, v := range urlParams {
		rctx := chi.NewRouteContext()
		rctx.URLParams.Add(v.key, v.value)
		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))
	}
	if err != nil {
		fmt.Errorf("request can't be created")
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handlerParam)
	handler.ServeHTTP(rr, req)

	return rr
}
func toResponse(responseByte *bytes.Buffer) utils.Response {
	var response utils.Response

	json.NewDecoder(responseByte).Decode(&response)

	return response
}

func TestUserHandler_Login(t *testing.T) {

	rr := sendRequest(map[string]any{"username": "mai", "password": "1"}, "POST", url+"login", userHandler.Login)

	if rr.Code != http.StatusOK {
		t.Errorf("It must Successfully")
	}
	rr2 := sendRequest(map[string]any{"username": "mai", "password": "2"}, "POST", url+"login", userHandler.Login)
	if rr2.Code != http.StatusBadRequest {
		t.Errorf("That is impossible")
	}

	rr3 := sendRequest(map[string]any{"username": "mai"}, "POST", url+"login", userHandler.Login)

	if toResponse(rr3.Body).Success != false {
		t.Errorf("That is impossible")
	}

}

func TestUserHandler_CreateUser(t *testing.T) {
	rr := sendRequest(data[0], "POST", url, userHandler.CreateUser)
	if rr.Code != http.StatusOK {
		t.Errorf("It must Successfully")
	}
	rr2 := sendRequest(data[1], "POST", url, userHandler.CreateUser)
	if rr2.Code != http.StatusBadRequest && toResponse(rr2.Body).Data != utils.UserAlreadyExist {
		t.Errorf("It must 400 Bad Request")
	}
	rr3 := sendRequest(map[string]any{"username": "mai", "password": "2"}, "POST", url, userHandler.CreateUser)
	if rr3.Code != http.StatusBadRequest {
		t.Errorf("It must 400 Bad Request")
	}
	rr4 := sendRequest(map[string]any{"username": "mai", "password": "2", "email": "mai@gmail.com"}, "POST", url, userHandler.CreateUser)
	if rr4.Code != http.StatusBadRequest {
		t.Errorf("It must 400 Bad Request")
	}
	rr5 := sendRequest(map[string]any{"username": "mai", "password": "2", "email": "mai@@@gmail.com"}, "POST", url, userHandler.CreateUser)
	if rr5.Code != http.StatusBadRequest {
		t.Errorf("It must 400 Bad Request")
	}
}

func TestUserHandler_UpdateUser(t *testing.T) {
	rr := sendRequest(data[0], "PUT", url, userHandler.UpdateUser, urlParams{"username", "mai"})

	if rr.Code != http.StatusOK {
		t.Errorf("It must Successfully")
	}
	rr2 := sendRequest(data[1], "PUT", url, userHandler.UpdateUser, urlParams{"username", "loc"})
	if rr2.Code != http.StatusNotFound {
		t.Errorf("It is not exist")
	}
	rr3 := sendRequest(map[string]any{"username": "mai", "password": "2"}, "PUT", url, userHandler.UpdateUser, urlParams{"username", "mai"})
	if rr3.Code != http.StatusBadRequest {
		t.Errorf("It must 400 Bad Request")
	}
	rr4 := sendRequest(map[string]any{"username": "mai", "password": "2", "email": "mai@gmail.com"}, "PUT", url, userHandler.UpdateUser, urlParams{"username", "mai"})
	if rr4.Code != http.StatusBadRequest {
		t.Errorf("It must 400 Bad Request")
	}
	rr5 := sendRequest(map[string]any{"username": "mai", "password": "2", "email": "mai@@@gmail.com"}, "PUT", url, userHandler.UpdateUser, urlParams{"username", "mai"})
	if rr5.Code != http.StatusBadRequest {
		t.Errorf("It must 400 Bad Request")
	}
}

func TestUserHandler_DeleteUser(t *testing.T) {
	rr := sendRequest(nil, "GET", url, userHandler.DeleteUser, urlParams{"username", "mai"})

	if rr.Code != http.StatusOK {
		t.Errorf("It must Successfully")
	}
	rr2 := sendRequest(nil, "GET", url, userHandler.DeleteUser, urlParams{"username", "loc"})

	if rr2.Code == http.StatusOK {
		t.Errorf("It must NOT FOUND")
	}
}

func TestMain(m *testing.M) {
	userServiceMock := new(mocks.UserService)

	userServiceMock.On("Login", "mai", "1").Return(&data[0], utils.NewError(utils.Successfully, http.StatusOK))
	userServiceMock.On("Login", "mai", "2").Return(nil, utils.NewError(utils.PasswordIsIncorrect, http.StatusBadRequest))

	userServiceMock.On("CreateUser", data[0]).Return(utils.NewError(utils.Successfully, http.StatusOK))
	userServiceMock.On("CreateUser", data[1]).Return(utils.NewError(utils.UserAlreadyExist, http.StatusBadRequest))

	userServiceMock.On("UpdateUser", data[0]).Return(utils.NewError(utils.Successfully, http.StatusOK))
	userServiceMock.On("UpdateUser", data[1]).Return(utils.NewError(utils.NotExist, http.StatusNotFound))

	userServiceMock.On("DeleteUser", "mai").Return(utils.NewError(utils.Successfully, http.StatusOK))
	userServiceMock.On("DeleteUser", "loc").Return(utils.NewError(utils.NotExist, http.StatusNotFound))
	userHandler = NewUserHandler(userServiceMock)

	m.Run()
	//
	//r := chi.NewRouter()
	//r.Route("/api", func(r chi.Router) {
	//	r.Route("/user", userHandler.UserHandler)
	//})
	//http.ListenAndServe(":5000", r)
}
