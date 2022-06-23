package service

import (
	"net/http"
	"s3corp-golang-fresher/internal/models"
	"s3corp-golang-fresher/internal/repository"
	"s3corp-golang-fresher/utils"
)

type DocItemServiceImpl struct {
	DocItemRepo repository.DocItemRepo
}

func NewDocItemService(docItemRepo repository.DocItemRepo) DocItemService {
	return &DocItemServiceImpl{DocItemRepo: docItemRepo}
}

func (docItemServiceImpl DocItemServiceImpl) GetDocItemById(docItemId string) (*models.DocumentItem, utils.Error) {
	docItem, err := docItemServiceImpl.DocItemRepo.GetDocItemById(docItemId)
	if err != nil {
		return docItem, utils.NewError("Internal Server Error", http.StatusInternalServerError)
	}

	return docItem, utils.NewError("Successfully", http.StatusOK)
}

func (docItemServiceImpl DocItemServiceImpl) GetDocItems() (models.DocumentItemSlice, utils.Error) {
	docItems, err := docItemServiceImpl.DocItemRepo.GetDocItems()

	if err != nil {
		return docItems, utils.NewError("Internal Server Error", http.StatusInternalServerError)
	}
	if len(docItems) == 0 {
		return []*models.DocumentItem{}, utils.NewError("Successfully", http.StatusOK)
	}
	return docItems, utils.NewError("Successfully", http.StatusOK)
}

func (docItemServiceImpl DocItemServiceImpl) CreateDocItem(docItem models.DocumentItem) utils.Error {
	err := docItemServiceImpl.DocItemRepo.CreateDocItem(docItem)
	if err != nil {
		return utils.NewError("Internal Server Error", http.StatusInternalServerError)
	}

	return utils.NewError("Successfully", http.StatusOK)
}

func (docItemServiceImpl DocItemServiceImpl) UpdateDocItem(docItem models.DocumentItem) utils.Error {
	affectedRows, err := docItemServiceImpl.DocItemRepo.UpdateDocItem(docItem)
	if err != nil {
		return utils.NewError("Internal Server Error", http.StatusInternalServerError)
	}
	if affectedRows <= 0 {
		return utils.Error{Message: "The record is not exists", StatusCode: 404}
	}
	return utils.NewError("Successfully", http.StatusOK)
}

func (docItemServiceImpl DocItemServiceImpl) DeleteDocItemById(docItemId string) utils.Error {
	affectedRows, err := docItemServiceImpl.DocItemRepo.DeleteDocById(docItemId)
	if err != nil {
		return utils.NewError("Internal Server Error", http.StatusInternalServerError)
	}
	if affectedRows <= 0 {
		return utils.Error{Message: "The record is not exists", StatusCode: 404}
	}
	return utils.NewError("Successfully", http.StatusOK)
}
