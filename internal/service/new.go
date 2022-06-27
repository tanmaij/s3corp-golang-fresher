package service

import (
	error2 "s3corp-golang-fresher/internal/errors"
	"s3corp-golang-fresher/internal/models"
)

type DocItemService interface {

	// GetDocItemById Return one documentItem with id parameter
	GetDocItemById(docItemId string) (*models.DocumentItem, error2.Error)

	// GetDocItems Return a subDocument Slice
	GetDocItems() (models.DocumentItemSlice, error2.Error)

	// CreateDocItem Insert data by documentItem parameter
	CreateDocItem(docItem models.DocumentItem) error2.Error

	// UpdateDocItem Update one record by documentItem parameter
	UpdateDocItem(docItem models.DocumentItem) error2.Error

	// DeleteDocItemById DeleteDocById Delete one record by documentItemId parameter
	DeleteDocItemById(docItemId string) error2.Error
}

type DocService interface {

	// GetDocById Return one document with id parameter
	GetDocById(docId string) (*models.Document, error2.Error)

	// GetDocs Return a Document Slice
	GetDocs() (models.DocumentSlice, error2.Error)

	// GetDocsByUsername Return all doc of user with id parameter
	GetDocsByUsername(username string) (models.DocumentSlice, error2.Error)

	// CreateDoc Insert data by document parameter
	CreateDoc(document models.Document) error2.Error

	// UpdateDoc Update one record by document parameter
	UpdateDoc(document models.Document) error2.Error

	// DeleteDocById Delete one record by id parameter
	DeleteDocById(documentId string) error2.Error
}
