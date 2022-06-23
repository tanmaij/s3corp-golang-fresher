package repository

import "s3corp-golang-fresher/internal/models"

type UserRepo interface {
	Login(username string) (*models.User, error)
	// GetUserByUsername Return one user with id parameter
	GetUserByUsername(username string) (*models.User, error)

	// GetUsers Return a user Slice
	GetUsers() (models.UserSlice, error)

	// CreateUser Insert data by user parameter
	CreateUser(user models.User) error

	// UpdateUser Update one record by user parameter
	UpdateUser(user models.User) (int64, error)

	// DeleteUser Delete one record by username parameter
	DeleteUser(username string) (int64, error)
}

type DocItemRepo interface {

	// GetDocItemById Return one documentItem with id parameter
	GetDocItemById(docItemId string) (*models.DocumentItem, error)

	// GetDocItems Return a subDocument Slice
	GetDocItems() (models.DocumentItemSlice, error)

	// CreateDocItem Insert data by documentItem parameter
	CreateDocItem(subDocument models.DocumentItem) error

	// UpdateDocItem Update one record by documentItem parameter
	UpdateDocItem(subDocument models.DocumentItem) (int64, error)

	// DeleteDocById Delete one record by documentItem parameter
	DeleteDocById(subDocumentId string) (int64, error)
}

type DocRepo interface {

	// GetDocById Return one document with id parameter
	GetDocById(id string) (*models.Document, error)

	// GetDocs Return a Document Slice
	GetDocs() (models.DocumentSlice, error)

	// GetDocsByUsername Return all doc of user with id parameter
	GetDocsByUsername(username string) (models.DocumentSlice, error)

	// CreateDoc Insert data by document parameter
	CreateDoc(document models.Document) error

	// UpdateDoc Update one record by document parameter
	UpdateDoc(document models.Document) (int64, error)

	// DeleteDocById Delete one record by id parameter
	DeleteDocById(documentId string) (int64, error)
}
