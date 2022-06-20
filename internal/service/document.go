package service

import (
	"s3corp-golang-fresher/internal/models"
	"s3corp-golang-fresher/internal/reponsitory"
)

type DocumentService struct {
	DocumentReponsitory *reponsitory.DocumentReponsitory
}

func (documentService DocumentService) GetOneById(documentId string) (*models.Document, error) {
	return documentService.DocumentReponsitory.GetOneById(documentId)
}

// GetAll Return a Document Slice
func (documentService DocumentService) GetAll() (models.DocumentSlice, error) {
	return documentService.DocumentReponsitory.GetAll()
}

// CreateOne Insert data by document parameter
func (documentService DocumentService) CreateOne(document models.Document) error {
	return documentService.DocumentReponsitory.CreateOne(document)
}

// UpdateOne Update one record by document parameter
func (documentService DocumentService) UpdateOne(document models.Document) (int64, error) {
	return documentService.DocumentReponsitory.UpdateOne(document)
}

// Delete one record by id parameter
func (documentService DocumentService) deleteOneById(documentId int) (int64, error) {
	return documentService.DocumentReponsitory.DeleteOneById(documentId)
}
