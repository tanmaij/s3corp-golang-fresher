package service

import (
	"s3corp-golang-fresher/internal/models"
	"s3corp-golang-fresher/utils"
)

type DocumentService interface {

	// GetOneById Return one Document with id parameter
	GetOneById(documentId string) (*models.Document, *utils.Error)

	// GetAll Return a Document Slice
	GetAll() (models.DocumentSlice, *utils.Error)

	// CreateOne Insert data by document parameter
	CreateOne(document models.Document) *utils.Error

	// UpdateOne Update one record by document parameter
	UpdateOne(document models.Document) *utils.Error

	// DeleteOneById Delete one record by id parameter
	DeleteOneById(documentId string) *utils.Error
}

type SubDocumentService interface {
	// GetByDocumentId Return a slice subdocument by Document paramater
	GetByDocumentId(documentId string) (models.SubDocumentSlice, *utils.Error)

	// GetOneById Return one subdocument with id parameter
	GetOneById(id int) (*models.SubDocument, *utils.Error)

	// GetAll Return a subDocument Slice
	GetAll() (models.SubDocumentSlice, *utils.Error)

	// CreateOne Insert data by subdocument parameter
	CreateOne(subDocument models.SubDocument) *utils.Error

	// UpdateOne Update one record by subdocument parameter
	UpdateOne(subDocument models.SubDocument) *utils.Error

	// DeleteOneById Delete one record by subdocument parameter
	DeleteOneById(subDocumentId int) *utils.Error
}
