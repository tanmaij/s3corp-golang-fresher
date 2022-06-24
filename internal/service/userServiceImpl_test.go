package service

import (
	"fmt"
	"github.com/volatiletech/null/v8"
	"net/http"
	"s3corp-golang-fresher/internal/models"
	"s3corp-golang-fresher/internal/repository/mocks"
	"s3corp-golang-fresher/utils"
	"testing"
)

var data = []models.User{
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

var userService UserService

func TestUserServiceImpl_UpdateUser(t *testing.T) {

	err1 := userService.UpdateUser(data[0])
	if err1.Message != utils.Successfully {
		t.Errorf("The error which is returned is inncorrect")
	}

	err2 := userService.UpdateUser(data[1])
	if err2.Message != utils.NotExist {
		t.Errorf("The error which is returned is inncorrect")
	}
}
func TestUserServiceImpl_DeleteUser(t *testing.T) {

	err1 := userService.DeleteUser("mai")
	if err1.Message != utils.Successfully {
		t.Errorf("The error which is returned is inncorrect")
	}

	err2 := userService.DeleteUser("loc")
	if err2.Message != utils.NotExist {
		t.Errorf("The error which is returned is inncorrect")
	}
}

func TestUserServiceImpl_CreateUser(t *testing.T) {

	err1 := userService.CreateUser(data[1])
	t.Log(err1.Message)
	if err1.StatusCode != 200 {
		t.Errorf("The error which is returned is inncorrect")
	}

	err2 := userService.CreateUser(data[0])

	if err2.Message != utils.UserAlreadyExist {
		t.Errorf("The error which is returned is inncorrect")
	}
}

func TestUserServiceImpl_GetUserByUsername(t *testing.T) {

	_, err1 := userService.GetUserByUsername("mai")
	if err1.StatusCode != 200 {
		t.Errorf("The error which is returned is inncorrect")
	}

	_, err2 := userService.GetUserByUsername("thu")
	if err2.StatusCode != 500 {
		t.Errorf("The error which is returned is inncorrect")
	}
}
func TestUserServiceImpl_Login(t *testing.T) {

	user, err := userService.Login("mai", "1")
	if err.StatusCode != 200 {
		t.Errorf("The error which is returned is inncorrect")
	}
	if user == nil {
		t.Errorf("User is nil")
	}

	_, err2 := userService.Login("mai", "2")
	if err2.Message != utils.PasswordIsIncorrect {
		t.Errorf("It may not")
	}

	_, err3 := userService.Login("long", "1")
	if err3.StatusCode != http.StatusInternalServerError {
		t.Errorf("It may not")
	}
}

func TestMain(m *testing.M) {
	userRepo := new(mocks.UserRepo)
	userService = NewUserService(userRepo)
	userRepo.On("GetUserByUsername", "mai").Return(&data[0], nil)
	userRepo.On("GetUserByUsername", "thu").Return(nil, fmt.Errorf("not Exist"))

	userRepo.On("Login", "mai").Return(&data[0], nil)
	userRepo.On("Login", "long").Return(nil, fmt.Errorf("not found"))

	userRepo.On("CreateUser", data[1]).Return(nil)
	userRepo.On("CreateUser", data[0]).Return(data[0], nil)
	userRepo.On("GetUserByUsername", "loc").Return(nil, fmt.Errorf("not exist"))

	userRepo.On("UpdateUser", data[0]).Return(int64(1), nil)
	userRepo.On("UpdateUser", data[1]).Return(int64(0), nil)

	userRepo.On("DeleteUser", "mai").Return(int64(1), nil)
	userRepo.On("DeleteUser", "loc").Return(int64(0), nil)
	m.Run()
}
func getUserByUserName(username string) (*models.User, error) {
	for i, v := range data {
		if v.Username == username {
			return &data[i], nil
		}
	}
	return nil, fmt.Errorf("NOT FOUND")
}
