package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	errors "s3corp-golang-fresher/internal/errors"
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

type paginationTest struct {
	Page  string `json:"page"`
	Limit string `json:"limit"`
}

func TestUserHandler_Login(t *testing.T) {
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
			input:     "test_data/user_handler/request/login_success.json",
			expResult: "test_data/user_handler/response/login_success.json",
			expStatus: http.StatusOK,
			expErr:    nil},
		"password_is_incorrect": {
			input:     "test_data/user_handler/request/login_password_is_incorrect.json",
			expStatus: http.StatusUnauthorized,
			expErr:    errors.NewError(errors.PasswordIsIncorrect, http.StatusUnauthorized)},
		"not_found": {
			input:     "test_data/user_handler/request/login_not_found.json",
			expStatus: http.StatusNotFound,
			expErr:    errors.NewError(errors.NotFound, http.StatusNotFound)},
	}

	// Define struct is the same type login handler's response
	type LoginResponse struct {
		User  models.User `json:"user"`
		Token string      `json:"token"`
	}

	for desc, tc := range tcs {
		t.Run(desc, func(t *testing.T) {

			// Given
			// Get input data from json file
			given, err := readJsonFile(tc.input)
			if err != nil {
				t.Error("Error on reading the input file")
			}

			// Define user input for test
			var testUser models.User
			if err := json.Unmarshal(given, &testUser); err != nil {
				t.Error("Error on reading the input file")
			}
			// Define response for test (if test case is not error case)
			var response LoginResponse
			if tc.expErr == nil {
				// Get result data from json file (if test case is not error case)
				result, err := readJsonFile(tc.expResult)
				if err != nil {
					t.Error("Error on reading the result file")
				}
				if err := json.Unmarshal(result, &response); err != nil {
					t.Log("This is error case")
				}
			}

			// Set up data will be return if method Login is called(with some different arguments)
			userServiceMock.On("Login", testUser.Username, testUser.Password).Return(response.User, response.Token, tc.expErr)

			// 1. Define body data to send request
			// Define http test request with query params for offset and limit
			r := httptest.NewRequest(http.MethodPost, url+"login", bytes.NewBuffer(given))
			// 2. Define http test response
			w := httptest.NewRecorder()

			//When
			//Call handler Login with body
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
			expErr:    nil},
		"user_is_already_exist": {
			input:     "test_data/user_handler/request/create_user_user_already_exist.json",
			expResult: errors.UserAlreadyExist,
			expStatus: http.StatusBadRequest,
			expErr:    errors.NewError(errors.UserAlreadyExist, http.StatusBadRequest)},
		"email_is_invalid": {
			input:     "test_data/user_handler/request/create_user_invalid_email.json",
			expResult: errors.InvalidEmail,
			expStatus: http.StatusBadRequest,
			expErr:    errors.NewError(errors.InvalidEmail, http.StatusBadRequest)},
	}

	for desc, tc := range tcs {
		t.Run(desc, func(t *testing.T) {

			// Given
			// Get input data from json file
			given, err := readJsonFile(tc.input)
			if err != nil {
				t.Error("Error on reading the input file")
			}

			// Define user input for test
			var testUser models.User
			if err := json.Unmarshal(given, &testUser); err != nil {
				t.Error("Error on reading the input file")
			}

			// Set up data will be return if method Login is called(with some different arguments)
			userServiceMock.On("CreateUser", testUser).Return(tc.expErr)

			// 1. Define body data to send request
			// Define http test request with query params for offset and limit
			r := httptest.NewRequest(http.MethodPost, url, bytes.NewBuffer(given))
			// 2. Define http test response
			w := httptest.NewRecorder()

			//When
			//Call handler Login with body
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

	// 1. Create new user service mock for test
	// Create new user handler with user service mock
	userServiceMock := new(mocks.UserServiceMock)
	userHandler = NewUserHandler(userServiceMock)

	// define struct for get user response
	type GetUserResponse struct {
		pagination utils.Pagination `json:"pagination"`
		users      models.UserSlice `json:"users"`
	}

	tcs := map[string]struct {
		input     utils.Pagination // using Page and Limit field
		expResult string
		expStatus int
		expErr    error
	}{
		"success": {
			input:     utils.Pagination{Page: 2, Limit: 2},
			expResult: "test_data/user_handler/response/get_users_success.json",
			expStatus: http.StatusOK,
			expErr:    nil},
		"invalid_data": {
			input:     utils.Pagination{Page: 2, Limit: -2},
			expStatus: http.StatusBadRequest,
			expErr:    fmt.Errorf(errors.InvalidData)},
	}

	for desc, tc := range tcs {
		t.Run(desc, func(t *testing.T) {
			// Given
			var expResponse GetUserResponse // Define response for test (if test case is not error case)
			if tc.expErr == nil {
				result, err := readJsonFile(tc.expResult)
				if err != nil {
					t.Error("Error on reading the result file")
				}
				if err := json.Unmarshal(result, &expResponse); err != nil { // Write data to response variable
					t.Log("This is error case")
				}
			}

			// Set up data will be return if method GetUsers is called(with some different arguments)
			userServiceMock.On("GetUsers", map[string]int{"limit": tc.input.Limit, "page": tc.input.Page}).Return(expResponse.users, expResponse.pagination, nil)

			// query data to send request
			r := httptest.NewRequest(http.MethodGet, url, nil)

			// Add limit and page variable to request
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
}

func TestMain(m *testing.M) {

	m.Run()
}
