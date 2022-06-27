package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	errors "s3corp-golang-fresher/internal/errors"
	"s3corp-golang-fresher/internal/models"
	"s3corp-golang-fresher/internal/service/mocks"
	"s3corp-golang-fresher/utils"
	"testing"
)

// define url to send request for test
const url = "/api/user/"

var userHandler UserHandler

// define type struct as input for each test case
type userTest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Name     string `json:"name"`
}

func TestUserHandler_Login(t *testing.T) {

	// 1. Define the data will be response if request successfully
	responseSuccess := map[string]any{"user": map[string]any{"username": "mai", "password": "1", "name": "Mãi", "email": "mai@gmail.com"}, "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6Im1haSJ9.qYNphS_Xycc7-XY9MD9o_kTHocUjV6kCH0hD1EzTDk4"}

	// 2. Create new user service mock for test
	// Set up data will be return if method Login is called(with some different arguments)
	// Create new user handler with user service mock
	userServiceMock := new(mocks.UserService)
	userServiceMock.On("Login", "mai", "1").Return(&models.User{Username: "mai", Password: "1", Name: "Mãi", Email: "mai@gmail.com"}, "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6Im1haSJ9.qYNphS_Xycc7-XY9MD9o_kTHocUjV6kCH0hD1EzTDk4", errors.NewError(errors.Successfully, http.StatusOK))
	userServiceMock.On("Login", "mai2", "1").Return(nil, "", errors.NewError(errors.NotFound, http.StatusInternalServerError))
	userServiceMock.On("Login", "mai", "2").Return(nil, "", errors.NewError(errors.PasswordIsIncorrect, http.StatusUnauthorized))
	userHandler = NewUserHandler(userServiceMock)

	tcs := map[string]struct {
		input     userTest
		expResult utils.Response
		expStatus int
		expErr    error
	}{
		"success": {
			input:     userTest{Username: "mai", Password: "1"},
			expResult: utils.Response{Success: true, Data: responseSuccess},
			expStatus: http.StatusOK,
			expErr:    nil},

		"password_is_incorrect": {
			input:     userTest{Username: "mai", Password: "2"},
			expResult: utils.Response{},
			expStatus: http.StatusUnauthorized,
			expErr:    fmt.Errorf(errors.PasswordIsIncorrect)},

		"not_found": {
			input:     userTest{Username: "mai2", Password: "1"},
			expResult: utils.Response{},
			expStatus: http.StatusNotFound,
			expErr:    fmt.Errorf(errors.InternalServerError)},
	}

	for desc, tc := range tcs {
		t.Run(desc, func(t *testing.T) {

			// 1. Define body data to send request
			b, _ := json.Marshal(tc.input)

			// Define http test request with query params for offset and limit
			r := httptest.NewRequest(http.MethodPost, url+"login", bytes.NewBuffer(b))

			// 2. Define http test response
			w := httptest.NewRecorder()

			// Call handler Login with body
			userHandler.Login(w, r)

			// 3. Check response code and body
			if tc.expErr != nil { // Must be error
				//Equal status code
				require.Equal(t, tc.expStatus, w.Code)
				//Equal error string
				require.EqualError(t, tc.expErr, w.Body.String())

			} else {
				// Must be success
				require.Equal(t, http.StatusOK, w.Code)

				// Define response data
				var res utils.Response

				// Parse body data to response data
				err := json.Unmarshal(w.Body.Bytes(), &res)
				if err != nil {
					t.Fatal(err)
				}
				// Compare the result
				require.Equal(t, tc.expResult, res)
			}
		})
	}

}

func TestUserHandler_CreateUser(t *testing.T) {

	// 2. Create new user service mock for test
	// Set up data will be return if method Login is called(with some different arguments)
	userServiceMock := new(mocks.UserService)
	userServiceMock.On("CreateUser", models.User{Username: "mai", Password: "1", Email: "mai@gmail.com", Name: "Mãi"}).Return(errors.NewError(errors.Successfully, http.StatusOK))
	userServiceMock.On("CreateUser", models.User{Username: "mai2", Password: "1", Email: "mai@gmail.com", Name: "Mãi"}).Return(errors.NewError(errors.UserAlreadyExist, http.StatusBadRequest))

	// Create new user handler with user service mock
	userHandler = NewUserHandler(userServiceMock)

	tcs := map[string]struct {
		input     userTest
		expResult utils.Response
		expStatus int
		expErr    error
	}{
		"successful_completion": {
			input:     userTest{"mai", "1", "mai@gmail.com", "Mãi"},
			expResult: utils.Response{Success: true, Data: "Create user successfully"},
			expStatus: http.StatusOK,
		},
		"user_already_exist": {
			input:     userTest{"mai2", "1", "mai@gmail.com", "Mãi"},
			expResult: utils.Response{},
			expStatus: http.StatusBadRequest,
			expErr:    fmt.Errorf(errors.UserAlreadyExist),
		},
	}

	for desc, tc := range tcs {
		t.Run(desc, func(t *testing.T) {

			// 1. Define body data to send request
			b, _ := json.Marshal(tc.input)

			// Define http test request with query params for offset and limit
			r := httptest.NewRequest(http.MethodPost, url, bytes.NewBuffer(b))

			// 2. Define http test response
			w := httptest.NewRecorder()

			// 3. Call handler createUser
			userHandler.CreateUser(w, r)

			// 4. Check response code and body
			if tc.expErr != nil {
				// Must be error
				require.Equal(t, tc.expStatus, w.Code)
				require.EqualError(t, tc.expErr, w.Body.String())

			} else {
				// Must be success
				require.Equal(t, http.StatusOK, w.Code)
				var res utils.Response
				err := json.Unmarshal(w.Body.Bytes(), &res)
				if err != nil {
					t.Fatal(err)
				}
				require.Equal(t, tc.expResult, res)
			}
		})

	}
}

func TestUserHandler_UpdateUser(t *testing.T) {
}

func TestUserHandler_DeleteUser(t *testing.T) {
}

func TestMain(m *testing.M) {

	m.Run()
}
