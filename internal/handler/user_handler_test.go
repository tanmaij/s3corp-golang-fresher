package handler

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"s3corp-golang-fresher/internal/errors"
	"s3corp-golang-fresher/internal/handler/test_data/fake_data"
	"s3corp-golang-fresher/internal/models"
	"s3corp-golang-fresher/internal/service/mocks"
	"s3corp-golang-fresher/utils"
	"strconv"
	"testing"
)

// define url to send request for test
const url = "/api/user/"

var userHandler UserHandler

func readJsonFile(path string) ([]byte, error) {
	// Open our jsonFile
	// if we os.Open returns error then handle it
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	b, _ := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}
	// defer the closing of our jsonFile so that we can parse it later on
	defer f.Close()

	return b, nil
}

func TestUserHandler_Login(t *testing.T) {
	// 1. Create new user service mock for test
	// Create new user handler with user service mock
	userServiceMock := new(mocks.UserServiceMock)
	userHandler = NewUserHandler(userServiceMock)

	// Define struct is the same type login handler's response
	type LoginResponse struct {
		User  models.User `json:"user"`
		Token string      `json:"token"`
	}

	// Define struct  which is return by user service login
	type GivenData struct {
		User  models.User
		Token string
		Error error
	}

	tcs := map[string]struct {
		input     string
		expResult string
		expStatus int
		expErr    error
		givenData GivenData
	}{
		"success": {
			input:     "test_data/user_handler/request/login_success.json",
			expResult: "test_data/user_handler/response/login_success.json",
			expStatus: http.StatusOK,
			expErr:    nil,
			givenData: GivenData{User: fake_data.UserLogin, Token: fake_data.TokenLogin, Error: nil},
		},
		"password_is_incorrect": {
			input:     "test_data/user_handler/request/login_password_is_incorrect.json",
			expStatus: http.StatusUnauthorized,
			expErr:    errors.NewError(errors.PasswordIsIncorrect, http.StatusUnauthorized),
			givenData: GivenData{Error: errors.NewError(errors.PasswordIsIncorrect, http.StatusUnauthorized)},
		},
		"not_found": {
			input:     "test_data/user_handler/request/login_not_found.json",
			expStatus: http.StatusNotFound,
			expErr:    errors.NewError(errors.NotFound, http.StatusNotFound),
			givenData: GivenData{Error: errors.NewError(errors.NotFound, http.StatusNotFound)},
		},
	}

	for desc, tc := range tcs {
		t.Run(desc, func(t *testing.T) {

			// Given
			input, err := readJsonFile(tc.input) // Get input data from json file (byte)
			if err != nil {
				t.Error("Error on reading the input file")
			}

			// Define user input for test
			var testUser models.User
			if err := json.Unmarshal(input, &testUser); err != nil {
				t.Error("Error on reading the input file")
			}
			// Define response for test (if test case is not error case)
			var response LoginResponse
			if tc.expErr == nil {
				// Get result data from json file (if test case is not error case)
				output, err := readJsonFile(tc.expResult)
				if err != nil {
					t.Error("Error on reading the result file")
				}
				if err := json.Unmarshal(output, &response); err != nil {
					t.Log("This is error case")
				}
			}

			// Set up data will be return if method Login is called
			userServiceMock.On("Login", testUser.Username, testUser.Password).Return(tc.givenData.User, tc.givenData.Token, tc.givenData.Error)

			// Define http test request with post method and body (from input)
			r := httptest.NewRequest(http.MethodPost, url+"login", bytes.NewBuffer(input))
			// Define http test response
			w := httptest.NewRecorder()

			//When
			userHandler.Login(w, r)

			//Then
			if tc.expErr != nil { // Must be error
				//Equal status code
				require.Equal(t, tc.expStatus, w.Code)
				//Equal error string
				require.EqualError(t, tc.expErr, w.Body.String())

			} else {

				// Must be success
				require.Equal(t, http.StatusOK, w.Code)

				// Define response data
				var res LoginResponse

				// Parse body data to response data
				err := json.Unmarshal(w.Body.Bytes(), &res)
				if err != nil {
					t.Fatal(err)
				}

				// Compare res and response which is defined
				require.Equal(t, response, res)
			}
		})
	}

}

func TestUserHandler_CreateUser(t *testing.T) {

	// 1. Create new user service mock for test
	// Create new user handler with user service mock
	userServiceMock := new(mocks.UserServiceMock)
	userHandler = NewUserHandler(userServiceMock)

	tcs := map[string]struct {
		input     string
		expResult string
		expStatus int
		expErr    error
	}{
		"success": {
			input:     "test_data/user_handler/request/create_user_success.json",
			expResult: "Create user successfully",
			expStatus: http.StatusOK,
			expErr:    nil,
		},
		"user_is_already_exist": {
			input:     "test_data/user_handler/request/create_user_user_already_exist.json",
			expResult: errors.UserAlreadyExist,
			expStatus: http.StatusBadRequest,
			expErr:    errors.NewError(errors.UserAlreadyExist, http.StatusBadRequest),
		},
		"email_is_invalid": {
			input:     "test_data/user_handler/request/create_user_invalid_email.json",
			expResult: errors.InvalidEmail,
			expStatus: http.StatusBadRequest,
			expErr:    errors.NewError(errors.InvalidEmail, http.StatusBadRequest),
		},
	}

	for desc, tc := range tcs {
		t.Run(desc, func(t *testing.T) {

			// Given
			input, err := readJsonFile(tc.input) // Get input data from json file
			if err != nil {
				t.Error("Error on reading the input file")
			}

			// Define user testUser for test (from input)
			var testUser models.User
			if err := json.Unmarshal(input, &testUser); err != nil {
				t.Error("Error on reading the input file")
			}

			// Set up data will be return if method CreateUser is called
			userServiceMock.On("CreateUser", testUser).Return(tc.expErr)

			// Define http test request with body(from input)
			r := httptest.NewRequest(http.MethodPost, url, bytes.NewBuffer(input))
			// Define http test response
			w := httptest.NewRecorder()

			//When
			userHandler.CreateUser(w, r)

			//Then
			if tc.expErr != nil { // Must be error
				require.Equal(t, tc.expStatus, w.Code)
				require.EqualError(t, tc.expErr, w.Body.String())

			} else {
				// Must be success
				require.Equal(t, http.StatusOK, w.Code)
				// Compare expect result and response body (string)
				require.Equal(t, tc.expResult, string(w.Body.Bytes()))
			}
		})
	}
}

func TestUserHandler_GetUsers(t *testing.T) {

	// Create new user service mock for test
	// Create new user handler with user service mock
	userServiceMock := new(mocks.UserServiceMock)
	userHandler = NewUserHandler(userServiceMock)

	// define struct for get user response
	type GetUserResponse struct {
		Pagination utils.Pagination `json:"pagination"`
		Users      []models.User    `json:"users"`
	}
	// define struct for user service mock return
	type GivenData struct {
		Pagination utils.Pagination
		Users      []models.User
		Error      error
	}

	tcs := map[string]struct {
		input     utils.Pagination // using Page and Limit field
		expResult string
		expStatus int
		expErr    error
		givenData GivenData
	}{
		"success": {
			input:     utils.Pagination{Page: 2, Limit: 2},
			expResult: "test_data/user_handler/response/get_users_success.json",
			expStatus: http.StatusOK,
			expErr:    nil,
			givenData: GivenData{Users: fake_data.UserSlice, Pagination: fake_data.Pagination, Error: nil},
		},
		"invalid_data": {
			input:     utils.Pagination{Page: 2, Limit: -2},
			expStatus: http.StatusBadRequest,
			expErr:    fmt.Errorf(errors.InvalidData),
		},
	}

	for desc, tc := range tcs {
		t.Run(desc, func(t *testing.T) {
			// Given
			var expResponse GetUserResponse // Define response for test (if test case is not error case)
			if tc.expErr == nil {
				output, err := readJsonFile(tc.expResult)
				if err != nil {
					t.Error("Error on reading the result file")
				}
				if err := json.Unmarshal(output, &expResponse); err != nil { // Write data to response variable
					t.Log("This is error case")
				}
			}

			// Set up data will be return if method GetUsers is called(with some different arguments)
			userServiceMock.On("GetUsers", map[string]int{"limit": tc.input.Limit, "page": tc.input.Page}).Return(tc.givenData.Users, tc.givenData.Pagination, tc.givenData.Error)

			// define request for test
			r := httptest.NewRequest(http.MethodGet, url, nil)

			// Add limit and page variable to request url params
			q := r.URL.Query()
			q.Add("limit", strconv.Itoa(tc.input.Limit))
			q.Add("page", strconv.Itoa(tc.input.Page))
			r.URL.RawQuery = q.Encode()

			w := httptest.NewRecorder()

			//When
			userHandler.GetUsers(w, r)

			//Then
			if tc.expErr != nil { // Must be error
				//Equal status code
				require.Equal(t, tc.expStatus, w.Code)
				//Equal error string
				require.EqualError(t, tc.expErr, w.Body.String())

			} else {
				// Must be success
				require.Equal(t, http.StatusOK, w.Code)

				var res GetUserResponse

				if err := json.Unmarshal(w.Body.Bytes(), &res); err != nil {
					t.Fatal(err)
				}

				// Compare expect result and response body
				require.Equal(t, expResponse, res)
			}
		})
	}
}

func TestUserHandler_DeleteUser(t *testing.T) {
	// 1. Create new user service mock for test
	// Create new user handler with user service mock
	userServiceMock := new(mocks.UserServiceMock)
	userHandler = NewUserHandler(userServiceMock)

	tcs := map[string]struct {
		input     string
		expResult string
		expStatus int
		expErr    error
	}{
		"success": {
			input:     "mai",
			expResult: "Delete user successfully",
			expStatus: http.StatusOK,
		},
		"not_exist": {
			input:     "loc",
			expResult: "Delete user successfully",
			expStatus: http.StatusNotFound,
			expErr:    errors.NewError(errors.NotExist, http.StatusNotFound),
		},
	}

	for desc, tc := range tcs {
		t.Run(desc, func(t *testing.T) {

			// Given

			// Set up data will be return if method DeleteUser is called
			userServiceMock.On("DeleteUser", tc.input).Return(tc.expErr)

			// Define http test request
			r := httptest.NewRequest(http.MethodDelete, url+tc.input, nil)
			// Define http test response
			w := httptest.NewRecorder()
			// Init chi route context
			rctx := chi.NewRouteContext()
			// Set username to chi route context
			rctx.URLParams.Add("username", tc.input)
			// Add chi route context to request
			r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, rctx))

			//When
			userHandler.DeleteUser(w, r)

			//Then
			if tc.expErr != nil { // Must be error
				require.Equal(t, tc.expStatus, w.Code)
				require.EqualError(t, tc.expErr, w.Body.String())

			} else {
				// Must be success
				require.Equal(t, http.StatusOK, w.Code)
				// Compare expect result and response body (string)
				require.Equal(t, tc.expResult, string(w.Body.Bytes()))
			}
		})
	}
}

func TestUserHandler_UpdateUser(t *testing.T) {
	// 1. Create new user service mock for test
	// Create new user handler with user service mock
	userServiceMock := new(mocks.UserServiceMock)
	userHandler = NewUserHandler(userServiceMock)

	type InputType struct {
		Username string
		Body     string
	}

	tcs := map[string]struct {
		input     InputType
		expResult string
		expStatus int
		expErr    error
	}{
		"success": {
			input:     InputType{"mai", "test_data/user_handler/request/update_user_success.json"},
			expResult: "Update user successfully",
			expStatus: http.StatusOK,
		},
		"not_found": {
			input:     InputType{"mai2", "test_data/user_handler/request/update_user_not_found.json"},
			expResult: "",
			expStatus: http.StatusNotFound,
			expErr:    errors.NewError(errors.NotFound, http.StatusNotFound),
		},
	}

	for desc, tc := range tcs {
		t.Run(desc, func(t *testing.T) {

			// Given
			input, err := readJsonFile(tc.input.Body) // Get input data from json file
			if err != nil {
				t.Error("Error on reading the input file")
			}

			// Define user testUser for test (from input)
			var testUser models.User
			testUser.Username = tc.input.Username                    // Set username for user service argument
			if err := json.Unmarshal(input, &testUser); err != nil { // Set data for user service argument
				t.Error("Error on reading the input file")
			}

			// Set up data will be return if method DeleteUser is called
			userServiceMock.On("UpdateUser", testUser).Return(tc.expErr)

			// Define http test request
			r := httptest.NewRequest(http.MethodPut, url+tc.input.Username, bytes.NewBuffer(input))
			// Define http test response
			w := httptest.NewRecorder()
			// Init chi route context
			rctx := chi.NewRouteContext()
			// Set username to chi route context
			rctx.URLParams.Add("username", tc.input.Username)
			// Add chi route context to request
			r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, rctx))

			//When
			userHandler.UpdateUser(w, r)

			//Then
			if tc.expErr != nil { // Must be error
				require.Equal(t, tc.expStatus, w.Code)
				require.EqualError(t, tc.expErr, w.Body.String())

			} else {
				// Must be success
				require.Equal(t, http.StatusOK, w.Code)
				// Compare expect result and response body (string)
				require.Equal(t, tc.expResult, string(w.Body.Bytes()))
			}
		})
	}
}
