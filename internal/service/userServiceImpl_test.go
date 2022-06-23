package service

import (
	"fmt"
	"github.com/volatiletech/null/v8"
	"s3corp-golang-fresher/internal/models"
	"s3corp-golang-fresher/internal/repository/mocks"
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

func TestUserServiceImpl_GetUserByUsername(t *testing.T) {
	userRepo := new(mocks.UserRepo)
	userService = NewUserService(userRepo)
	userRepo.On("GetUserByUsername", "mai").Return(&data[0], nil)
	_, err1 := userService.GetUserByUsername("mai")
	if err1.StatusCode != 200 {
		t.Errorf("The error which is returned is inncorrect")
	}

	_, err2 := userService.GetUserByUsername("t")
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
	if err2.Message != "Password is incorrect" {
		t.Errorf("It may not")
	}
}

func TestMain(m *testing.M) {

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
