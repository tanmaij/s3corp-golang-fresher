package service

import (
	"github.com/stretchr/testify/require"
	"net/http"
	"os"
	"s3corp-golang-fresher/internal/errors"
	"s3corp-golang-fresher/internal/models"
	"s3corp-golang-fresher/internal/repository/mocks"
	"s3corp-golang-fresher/internal/service/test_data/fake_data"
	"testing"
)

var userService UserService

func TestUserServiceImpl_GetUsers(t *testing.T) {

}
func TestUserServiceImpl_UpdateUser(t *testing.T) {

}
func TestUserServiceImpl_DeleteUser(t *testing.T) {

}
func TestUserServiceImpl_CreateUser(t *testing.T) {

}
func TestUserServiceImpl_GetUserByUsername(t *testing.T) {

}
func TestUserServiceImpl_UsersStatsCSVFile(t *testing.T) {
	userRepoMock := new(mocks.UserRepoMock)
	userService = NewUserService(userRepoMock)

	//Define type struct the same type as userRepo return
	type GivenData struct {
		Users models.UserSlice
		Error error
	}

	tcs := map[string]struct {
		input     int
		expResult string
		expErr    error
		givenData GivenData
	}{
		"success": {
			input:     2022,
			expResult: "test_data/user_service/output/users_stat_csv_file_success.csv",
			givenData: GivenData{fake_data.UserSliceByYear, nil},
		},
		"no_data_available": {
			input:     2021,
			givenData: GivenData{models.UserSlice{}, nil},
			expErr:    errors.NewError(errors.NoDataAvailable, http.StatusNotFound),
		},
	}

	for desc, tc := range tcs {
		t.Run(desc, func(t *testing.T) {

			// Given
			var expFile []byte // Define expect file
			if tc.expErr == nil {
				var err error
				expFile, err = os.ReadFile(tc.expResult)
				if err != nil {
					t.Error("Error on reading the result file")
				}
			}

			// Set up data will be return if method GetUsers is called(with some different arguments)
			userRepoMock.On("GetUsersByYear", tc.input).Return(tc.givenData.Users, tc.givenData.Error)

			//When
			res, err := userService.UsersStatsCSVFile(2022)

			//Then
			if tc.expErr != nil { // Must be error
				//Equal error
				require.EqualError(t, tc.expErr, err.Error())

			} else { // Must be success
				// Compare expect result and result
				require.Equal(t, expFile, res)
			}
		})
	}

}
func TestMain(m *testing.M) {
	m.Run()
}
