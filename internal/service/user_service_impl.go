package service

import (
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"log"
	"net/http"
	"s3corp-golang-fresher/internal/errors"
	"s3corp-golang-fresher/internal/models"
	"s3corp-golang-fresher/internal/repository"
	"s3corp-golang-fresher/pkg"
	"s3corp-golang-fresher/utils"
	"strconv"
)

type UserServiceImpl struct {
	UserRepo repository.UserRepo
}

func NewUserService(userRepo repository.UserRepo) UserService {
	return &UserServiceImpl{userRepo}
}

func (userServiceImpl UserServiceImpl) Login(username string, password string) (*models.User, string, errors.Error) {

	// 1. Get instance jwtAuth from pkg package
	jwtAuth := pkg.GetJWTAuth()

	// 2. Get user record (username and password)
	// if error, return the error and nil data
	// if not error, continue to next step
	user, err := userServiceImpl.UserRepo.GetUserByUsername(username)
	if err != nil {
		return user, "", errors.NewError(errors.NotFound, http.StatusInternalServerError)
	}

	// 3. Verify password parameter with user's password from user repository return
	// If they are different, return password incorrect error
	// Continue to next step if they are them same
	if user.Password != password {
		return user, "", errors.NewError(errors.PasswordIsIncorrect, http.StatusUnauthorized)
	}

	// 4. Create a token string with username
	_, token, _ := jwtAuth.Encode(map[string]any{"username": user.Username})

	// 5. return user information, token string and not error
	return user, token, errors.NewError(errors.Successfully, http.StatusOK)
}

func (userServiceImpl UserServiceImpl) GetUserByUsername(username string) (*models.User, errors.Error) {
	user, err := userServiceImpl.UserRepo.GetUserByUsername(username)

	if err != nil {
		return user, errors.NewError(errors.InternalServerError, http.StatusInternalServerError)
	}
	return user, errors.NewError(errors.Successfully, http.StatusOK)
}

func (userServiceImpl UserServiceImpl) GetUsers(queriesParams map[string]string) (models.UserSlice, *utils.Pagination, errors.Error) {

	var queries []qm.QueryMod
	limit, ok := queriesParams["limit"]
	page, ok2 := queriesParams["page"]
	_limit, err1 := strconv.Atoi(limit)
	_page, err2 := strconv.Atoi(page)
	if ok2 && ok {
		if err1 == nil && err2 == nil {
			queries = append(queries, qm.Limit(_limit), qm.Offset((_page-1)*_limit))
		} else {
			return nil, nil, errors.NewError(errors.QueryDataIncorrect, http.StatusNotFound)
		}
	}
	name, ok := queriesParams["name"]
	if ok && name != "" {
		queries = append(queries, qm.Where("name LIKE %?% ", name))
	}
	users, err := userServiceImpl.UserRepo.GetUsers(queries...)
	if err != nil {
		return users, nil, errors.NewError(errors.InternalServerError, http.StatusInternalServerError)
	}
	allusers, err := userServiceImpl.UserRepo.GetUsers(nil)
	totalPages := len(allusers) / _limit
	if len(allusers)%_limit != 0 {
		totalPages++
	}
	return users,
		utils.NewPagination(totalPages, _limit, _page, len(allusers)),
		errors.NewError(errors.Successfully, http.StatusOK)
}

func (userServiceImpl UserServiceImpl) CreateUser(user models.User) errors.Error {

	// 1. Find one user with username
	// If user is exists, return UserAlreadyExist Error
	checkUser, err := userServiceImpl.UserRepo.GetUserByUsername(user.Username)
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

	// return error struct with status 200 (it means successfully)
	return errors.NewError(errors.Successfully, http.StatusOK)
}

func (userServiceImpl UserServiceImpl) UpdateUser(user models.User) errors.Error {
	affectedRows, err := userServiceImpl.UserRepo.UpdateUser(user)

	if affectedRows <= 0 {
		return errors.NewError(errors.NotExist, http.StatusNotFound)
	}
	if err != nil {
		return errors.NewError(errors.Successfully, http.StatusInternalServerError)
	}
	return errors.NewError(errors.Successfully, http.StatusOK)
}

func (userServiceImpl UserServiceImpl) DeleteUser(username string) errors.Error {
	affectedRows, err := userServiceImpl.UserRepo.DeleteUser(username)
	if affectedRows <= 0 {
		return errors.NewError(errors.NotExist, http.StatusNotFound)
	}
	if err != nil {
		return errors.NewError(errors.InternalServerError, http.StatusInternalServerError)
	}
	return errors.NewError(errors.Successfully, http.StatusOK)
}
