package service

import (
	"s3corp-golang-fresher/internal/models"
	"s3corp-golang-fresher/utils"
)

type UserService interface {
	// Login with username and password
	Login(username string, password string) (*models.User, utils.Error)
	// GetUserByUsername Return one user with id parameter
	GetUserByUsername(username string) (*models.User, utils.Error)

	// GetUsers Return a user Slice
	GetUsers() (models.UserSlice, utils.Error)

	// CreateUser Insert data by user parameter
	CreateUser(user models.User) utils.Error

	// UpdateUser Update one record by user parameter
	UpdateUser(user models.User) utils.Error

	// DeleteUser Delete one record by username parameter
	DeleteUser(username string) utils.Error
}

type DocItemService interface {

	// GetDocItemById Return one documentItem with id parameter
	GetDocItemById(docItemId string) (*models.DocumentItem, utils.Error)

	// GetDocItems Return a subDocument Slice
	GetDocItems() (models.DocumentItemSlice, utils.Error)

	// CreateDocItem Insert data by documentItem parameter
	CreateDocItem(docItem models.DocumentItem) utils.Error

	// UpdateDocItem Update one record by documentItem parameter
	UpdateDocItem(docItem models.DocumentItem) utils.Error

	// DeleteDocItemById DeleteDocById Delete one record by documentItemId parameter
	DeleteDocItemById(docItemId string) utils.Error
}

type DocService interface {

	// GetDocById Return one document with id parameter
	GetDocById(docId string) (*models.Document, utils.Error)

	// GetDocs Return a Document Slice
	GetDocs() (models.DocumentSlice, utils.Error)

	// GetDocsByUsername Return all doc of user with id parameter
	GetDocsByUsername(username string) (models.DocumentSlice, utils.Error)

	// CreateDoc Insert data by document parameter
	CreateDoc(document models.Document) utils.Error

	// UpdateDoc Update one record by document parameter
	UpdateDoc(document models.Document) utils.Error

	// DeleteDocById Delete one record by id parameter
	DeleteDocById(documentId string) utils.Error
}
