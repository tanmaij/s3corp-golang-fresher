package service

import (
	"log"
	"net/http"
	"s3corp-golang-fresher/internal/errors"
	"s3corp-golang-fresher/internal/models"
	"s3corp-golang-fresher/internal/repository"
	"s3corp-golang-fresher/pkg"
	"s3corp-golang-fresher/utils"
	"strings"
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

	// 4. Create a token string with username and role
	_, token, err := jwtAuth.Encode(map[string]any{"username": user.Username, "role": user.Role})
	if err != nil {
		return *user, "", errors.NewError(errors.InternalServerError, http.StatusInternalServerError)
	}
	// 5. return user information, token string and not error
	return *user, token, err
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
func (userServiceImpl UserServiceImpl) UsersStatsCSVFile(year int) ([]byte, error) {

	//Get data from repo (sorted)
	users, err := userServiceImpl.UserRepo.GetUsersByYear(year)
	if err != nil {
		return nil, errors.NewError(errors.InternalServerError, http.StatusInternalServerError)
	}
	if len(users) == 0 {
		return nil, errors.NewError(errors.NoDataAvailable, http.StatusNotFound)
	}
	// Define All record
	var records = make([]string, len(users))

	type record struct {
		index int // Index sorted of record
		desc  string
	}
	// Define worker to handle some user data (users slice, record channel, start index)
	var worker = func(users models.UserSlice, r chan<- record, index int) {
		for i, v := range users {
			// One row
			desc := v.Username + "," + v.Name + "," + v.Email + "," + v.CreatedAt.Format("02-01-2006") + "\n"
			r <- record{i + index, desc}
		}
	}

	numUsers := len(users)                           // Number of data
	numWorker := 5                                   // Number of worker
	rowChan := make(chan record, numUsers/numWorker) // Per row channel

	for i := 0; i < numWorker; i++ {
		// index and last index to slice
		start := i * (len(users) / numWorker)
		end := (i + 1) * (len(users) / numWorker)
		// start routine
		go worker(users[start:end], rowChan, start)
	}

	for i := 0; i < numUsers; i++ {
		// get one row to the records, use index to make sure the record is sorted
		oneRow := <-rowChan
		records[oneRow.index] = oneRow.desc
		if oneRow.index == numUsers-1 {
			records[oneRow.index] = strings.ReplaceAll(oneRow.desc, "\n", "")
		}
	}

	// Join the records to the result(text of file) and add first rows as a header
	result := "USERNAME,NAME,EMAIL,CREATED AT\n" + strings.Join(records, "")
	return []byte(result), nil
}
