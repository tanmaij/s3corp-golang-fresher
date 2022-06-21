package repository

import "s3corp-golang-fresher/internal/models"

type SubdocumentRepository interface {

	// GetByDocumentId Return a slice subdocument by Document paramater
	GetByDocumentId(documentId string) (models.SubDocumentSlice, error)

	// GetOneById Return one subdocument with id parameter
	GetOneById(id int) (*models.SubDocument, error)

	// GetAll Return a subDocument Slice
	GetAll() (models.SubDocumentSlice, error)

	// CreateOne Insert data by subdocument parameter
	CreateOne(subDocument models.SubDocument) error

	// UpdateOne Update one record by subdocument parameter
	UpdateOne(subDocument models.SubDocument) (int64, error)

	// DeleteOneById Delete one record by subdocument parameter
	DeleteOneById(subDocumentId int) (int64, error)
}

type DocumentRepository interface {

	// GetOneById Return one document with id parameter
	GetOneById(id string) (*models.Document, error)

	// GetAll Return a Document Slice
	GetAll() (models.DocumentSlice, error)

	// CreateOne Insert data by document parameter
	CreateOne(document models.Document) error

	// UpdateOne Update one record by document parameter
	UpdateOne(document models.Document) (int64, error)

	// DeleteOneById Delete one record by id parameter
	DeleteOneById(documentId string) (int64, error)
}
