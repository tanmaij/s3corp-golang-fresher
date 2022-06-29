package service

import (
	"log"
	"net/http"
	"s3corp-golang-fresher/internal/errors"
	"s3corp-golang-fresher/internal/models"
	"s3corp-golang-fresher/internal/repository"
	"s3corp-golang-fresher/pkg"
	"s3corp-golang-fresher/utils"
)

type UserServiceImpl struct {
	UserRepo repository.UserRepo
}

func NewUserService(userRepo repository.UserRepo) UserService {
	return &UserServiceImpl{userRepo}
}

func (userServiceImpl UserServiceImpl) Login(username string, password string) (models.User, string, error) {

	// 1. Get instance jwtAuth from pkg package
	jwtAuth := pkg.GetJWTAuth()

	// 2. Get user record (username and password)
	// if error, return the error and nil data
	// if not error, continue to next step
	user, err := userServiceImpl.UserRepo.GetUserByUsername(username)
	if err != nil {
		return *user, "", errors.NewError(errors.NotFound, http.StatusNotFound)
	}

	// 3. Verify password parameter with user's password from user repository return
	// If they are different, return password incorrect error
	// Continue to next step if they are them same
	if user.Password != password {
		return *user, "", errors.NewError(errors.PasswordIsIncorrect, http.StatusUnauthorized)
	}

	// 4. Create a token string with username
	_, token, _ := jwtAuth.Encode(map[string]any{"username": user.Username})

	// 5. return user information, token string and not error
	return *user, token, nil
}

func (userServiceImpl UserServiceImpl) GetUserByUsername(username string) (models.User, error) {
	user, err := userServiceImpl.UserRepo.GetUserByUsername(username)

	if err != nil {
		return *user, errors.NewError(errors.InternalServerError, http.StatusInternalServerError)
	}
	return *user, errors.NewError(errors.Successfully, http.StatusOK)
}

func (userServiceImpl UserServiceImpl) GetUsers(queriesParams map[string]int) ([]models.User, utils.Pagination, error) {

	//limit := queriesParams["limit"]
	//page := queriesParams["page"]

	_, err := userServiceImpl.UserRepo.GetUsers()

	if err != nil {
		return []models.User{},
			utils.Pagination{}, errors.NewError(errors.InternalServerError, http.StatusInternalServerError)
	}

	//skip := (page - 1) * limit
	//get := (page-1)*limit + limit
	//totalRows := len(users)
	//totalPages := totalRows / limit
	//
	//if totalRows%limit == 0 {
	//	totalPages++
	//}
	//
	//users = users[skip:get]

	return []models.User{},
		utils.Pagination{}, nil
}

func (userServiceImpl UserServiceImpl) CreateUser(user models.User) error {

	// 1. Find one user with username
	// If user is exists, return UserAlreadyExist Error
	checkUser, err := userServiceImpl.UserRepo.GetByUsernameOrEmail(user.Username, user.Email)
	if checkUser != nil {
		return errors.NewError(errors.UserAlreadyExist, http.StatusBadRequest)
	}

	// 2. Call create method from userRepo
	// return errors if any
	err = userServiceImpl.UserRepo.CreateUser(user)
	if err != nil {
		log.Fatalln(err)
		return errors.NewError(errors.InternalServerError, http.StatusInternalServerError)
	}

	// return non error
	return nil
}

func (userServiceImpl UserServiceImpl) UpdateUser(user models.User) error {
	affectedRows, err := userServiceImpl.UserRepo.UpdateUser(user)

	if affectedRows <= 0 {
		return errors.NewError(errors.NotExist, http.StatusNotFound)
	}
	if err != nil {
		return errors.NewError(errors.Successfully, http.StatusInternalServerError)
	}
	return errors.NewError(errors.Successfully, http.StatusOK)
}

func (userServiceImpl UserServiceImpl) DeleteUser(username string) error {
	// number of Affected rows
	affectedRows, err := userServiceImpl.UserRepo.DeleteUser(username)
	//if Affected rows < 0 , return not exist error
	if affectedRows <= 0 {
		return errors.NewError(errors.NotExist, http.StatusNotFound)
	}
	if err != nil {
		return errors.NewError(errors.InternalServerError, http.StatusInternalServerError)
	}
	return nil
}
