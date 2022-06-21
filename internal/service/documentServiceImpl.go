package service

import (
	"s3corp-golang-fresher/internal/models"
	"s3corp-golang-fresher/internal/repository"
	"s3corp-golang-fresher/utils"
)

type DocumentServiceImpl struct {
	DocumentReponsitory repository.DocumentRepository
}

func (documentServiceImpl DocumentServiceImpl) GetOneById(documentId string) (*models.Document, *utils.Error) {
	document, err := documentServiceImpl.DocumentReponsitory.GetOneById(documentId)

	if err != nil {
		return document, &utils.Error{Message: "Internal Server Error", StatusCode: 500}
	}

	return document, nil
}

// GetAll Return a Document Slice
func (documentServiceImpl DocumentServiceImpl) GetAll() (models.DocumentSlice, *utils.Error) {
	documents, err := documentServiceImpl.DocumentReponsitory.GetAll()
	if err != nil {
		return documents, &utils.Error{Message: "Internal Server Error", StatusCode: 500}
	}
	if len(documents) == 0 {
		return []*models.Document{}, nil
	}
	return documents, nil
}

// CreateOne Insert data by document parameter
func (documentServiceImpl DocumentServiceImpl) CreateOne(document models.Document) *utils.Error {
	err := documentServiceImpl.DocumentReponsitory.CreateOne(document)
	if err != nil {
		return &utils.Error{Message: "Internal Server Error", StatusCode: 500}
	}
	return nil
}

// UpdateOne Update one record by document parameter
func (documentServiceImpl DocumentServiceImpl) UpdateOne(document models.Document) *utils.Error {
	affectedRows, err := documentServiceImpl.DocumentReponsitory.UpdateOne(document)

	if err != nil {
		return &utils.Error{Message: "Internal Server Error", StatusCode: 500}
	}
	if affectedRows <= 0 {
		return &utils.Error{Message: "The record is not exists", StatusCode: 404}
	}
	return nil
}

// DeleteOneById Delete one record by id parameter
func (documentServiceImpl DocumentServiceImpl) DeleteOneById(documentId string) *utils.Error {
	affectedRows, err := documentServiceImpl.DocumentReponsitory.DeleteOneById(documentId)

	if err != nil {
		return &utils.Error{Message: "Internal Server Error", StatusCode: 500}
	}
	if affectedRows <= 0 {
		return &utils.Error{Message: "The record is not exists", StatusCode: 404}
	}
	return nil
}
