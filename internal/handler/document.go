package handler

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/volatiletech/null/v8"
	"net/http"
	"s3corp-golang-fresher/internal/models"
	"s3corp-golang-fresher/internal/service"
)

type DocumentHandler struct {
	DocumentService *service.DocumentService
}

func (documentHandler DocumentHandler) DocumentHandler(r chi.Router) {
	r.Get("/{id}", documentHandler.getOneById)
	r.Get("/", documentHandler.getAll)
	r.Put("/", documentHandler.updateOneById)
	r.Delete("/", documentHandler.deleteOneById)
	r.Post("/", documentHandler.createOne)
}

func (documentHandler DocumentHandler) getOneById(w http.ResponseWriter, r *http.Request) {

	documentId := (chi.URLParam(r, "id"))

	document, err := documentHandler.DocumentService.GetOneById(documentId)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(Response{false, "Internal Server Error"})
		return
	}
	if document == nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(Response{false, "The record is not Exíst"})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(Response{true, document})
}

func (documentHandler DocumentHandler) getAll(w http.ResponseWriter, r *http.Request) {

	documents, err := documentHandler.DocumentService.GetAll()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(Response{false, "Internal Server Error"})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(Response{true, documents})

}
func (documentHandler DocumentHandler) createOne(w http.ResponseWriter, r *http.Request) {
	document := make(map[string]interface{})
	json.NewDecoder(r.Body).Decode(&document)

	subject, ok := document["subject"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Response{false, "The field 'subject' is NOT FOUND"})
		return
	}

	err := documentHandler.DocumentService.CreateOne(models.Document{Subject: null.String{(subject.(string)), true}})

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(Response{false, "Internal Server Error"})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(Response{true, "Create data successfully"})
}
func (documentHandler DocumentHandler) updateOneById(w http.ResponseWriter, r *http.Request) {

}
func (documentHandler DocumentHandler) deleteOneById(w http.ResponseWriter, r *http.Request) {

}
