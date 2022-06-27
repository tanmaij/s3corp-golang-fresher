package service

import (
	"net/http"
	"s3corp-golang-fresher/internal/errors"
	"s3corp-golang-fresher/internal/models"
	"s3corp-golang-fresher/internal/repository"
)

type DocItemServiceImpl struct {
	DocItemRepo repository.DocItemRepo
}

func NewDocItemService(docItemRepo repository.DocItemRepo) DocItemService {
	return &DocItemServiceImpl{DocItemRepo: docItemRepo}
}

func (docItemServiceImpl DocItemServiceImpl) GetDocItemById(docItemId string) (*models.DocumentItem, errors.Error) {
	docItem, err := docItemServiceImpl.DocItemRepo.GetDocItemById(docItemId)
	if err != nil {
		return docItem, errors.NewError("Internal Server Error", http.StatusInternalServerError)
	}

	return docItem, errors.NewError("Successfully", http.StatusOK)
}

func (docItemServiceImpl DocItemServiceImpl) GetDocItems() (models.DocumentItemSlice, errors.Error) {
	docItems, err := docItemServiceImpl.DocItemRepo.GetDocItems()

	if err != nil {
		return docItems, errors.NewError("Internal Server Error", http.StatusInternalServerError)
	}
	if len(docItems) == 0 {
		return []*models.DocumentItem{}, errors.NewError("Successfully", http.StatusOK)
	}
	return docItems, errors.NewError("Successfully", http.StatusOK)
}

func (docItemServiceImpl DocItemServiceImpl) CreateDocItem(docItem models.DocumentItem) errors.Error {
	err := docItemServiceImpl.DocItemRepo.CreateDocItem(docItem)
	if err != nil {
		return errors.NewError("Internal Server Error", http.StatusInternalServerError)
	}

	return errors.NewError("Successfully", http.StatusOK)
}

func (docItemServiceImpl DocItemServiceImpl) UpdateDocItem(docItem models.DocumentItem) errors.Error {
	affectedRows, err := docItemServiceImpl.DocItemRepo.UpdateDocItem(docItem)
	if err != nil {
		return errors.NewError("Internal Server Error", http.StatusInternalServerError)
	}
	if affectedRows <= 0 {
		return errors.Error{Message: "The record is not exists", StatusCode: 404}
	}
	return errors.NewError("Successfully", http.StatusOK)
}

func (docItemServiceImpl DocItemServiceImpl) DeleteDocItemById(docItemId string) errors.Error {
	affectedRows, err := docItemServiceImpl.DocItemRepo.DeleteDocById(docItemId)
	if err != nil {
		return errors.NewError("Internal Server Error", http.StatusInternalServerError)
	}
	if affectedRows <= 0 {
		return errors.Error{Message: "The record is not exists", StatusCode: 404}
	}
	return errors.NewError("Successfully", http.StatusOK)
}
