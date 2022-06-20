package service

import (
	"s3corp-golang-fresher/internal/models"
	"s3corp-golang-fresher/internal/reponsitory"
)

type SubDocumentService struct {
	subDocumentReponsitory *reponsitory.SubDocumentReponsitory
}

// GetByDocumentId Return a slice subdocument by Document paramater
func (subDocumentService SubDocumentService) GetByDocumentId(documentId string) (models.SubDocumentSlice, error) {
	return subDocumentService.subDocumentReponsitory.GetByDocumentId(documentId)
}

// GetOneById Return one subdocument with id parameter
func (subDocumentService SubDocumentService) GetOneById(id int) (*models.SubDocument, error) {
	return subDocumentService.subDocumentReponsitory.GetOneById(id)
}

// GetAll Return a subDocument Slice
func (subDocumentService SubDocumentService) GetAll() (models.SubDocumentSlice, error) {
	return subDocumentService.subDocumentReponsitory.GetAll()
}

// CreateOne Insert data by subdocument parameter
func (subDocumentService SubDocumentService) CreateOne(subDocument models.Document) error {
	return subDocumentService.subDocumentReponsitory.CreateOne(subDocument)
}

// UpdateOne Update one record by subdocument parameter
func (subDocumentService SubDocumentService) UpdateOne(subDocument models.SubDocument) (int64, error) {
	return subDocumentService.subDocumentReponsitory.UpdateOne(subDocument)
}

// DeleteOneById Delete one record by subdocument parameter
func (subDocumentService SubDocumentService) DeleteOneById(subDocumentId int) (int64, error) {
	return subDocumentService.subDocumentReponsitory.DeleteOneById(subDocumentId)
}
