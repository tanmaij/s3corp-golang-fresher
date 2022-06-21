package service

import (
	"s3corp-golang-fresher/internal/models"
	"s3corp-golang-fresher/internal/repository"
	"s3corp-golang-fresher/utils"
)

type SubDocumentServiceImpl struct {
	SubDocumentReponsitory repository.SubdocumentRepository
}

func (subDocumentServiceImpl SubDocumentServiceImpl) GetByDocumentId(documentId string) (models.SubDocumentSlice, *utils.Error) {
	subDocuments, err := subDocumentServiceImpl.SubDocumentReponsitory.GetByDocumentId(documentId)
	if err != nil {
		return subDocuments, &utils.Error{Message: "Internal Server Error", StatusCode: 500}
	}

	return subDocuments, nil

}

func (subDocumentServiceImpl SubDocumentServiceImpl) GetOneById(id int) (*models.SubDocument, *utils.Error) {
	subDocument, err := subDocumentServiceImpl.SubDocumentReponsitory.GetOneById(id)
	if err != nil {
		return subDocument, &utils.Error{Message: "Internal Server Error", StatusCode: 500}
	}

	return subDocument, nil
}

func (subDocumentServiceImpl SubDocumentServiceImpl) GetAll() (models.SubDocumentSlice, *utils.Error) {
	subDocuments, err := subDocumentServiceImpl.SubDocumentReponsitory.GetAll()

	if err != nil {
		return subDocuments, &utils.Error{Message: "Internal Server Error", StatusCode: 500}
	}
	if len(subDocuments) == 0 {
		return []*models.SubDocument{}, nil
	}
	return subDocuments, nil
}

func (subDocumentServiceImpl SubDocumentServiceImpl) CreateOne(subDocument models.SubDocument) *utils.Error {
	err := subDocumentServiceImpl.SubDocumentReponsitory.CreateOne(subDocument)
	if err != nil {
		return &utils.Error{Message: "Internal Server Error", StatusCode: 500}
	}

	return nil
}

func (subDocumentServiceImpl SubDocumentServiceImpl) UpdateOne(subDocument models.SubDocument) *utils.Error {
	affectedRows, err := subDocumentServiceImpl.SubDocumentReponsitory.UpdateOne(subDocument)
	if err != nil {
		return &utils.Error{Message: "Internal Server Error", StatusCode: 500}
	}
	if affectedRows <= 0 {
		return &utils.Error{Message: "The record is not exists", StatusCode: 404}
	}
	return nil
}

func (subDocumentServiceImpl SubDocumentServiceImpl) DeleteOneById(subDocumentId int) *utils.Error {
	affectedRows, err := subDocumentServiceImpl.SubDocumentReponsitory.DeleteOneById(subDocumentId)
	if err != nil {
		return &utils.Error{Message: "Internal Server Error", StatusCode: 500}
	}
	if affectedRows <= 0 {
		return &utils.Error{Message: "The record is not exists", StatusCode: 404}
	}
	return nil
}
