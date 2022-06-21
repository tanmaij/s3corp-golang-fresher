package handler

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/volatiletech/null/v8"
	"net/http"
	"s3corp-golang-fresher/internal/models"
	"s3corp-golang-fresher/internal/service"
	"s3corp-golang-fresher/utils"
)

type DocumentHandler struct {
	DocumentService service.DocumentService
}

func (documentHandler DocumentHandler) DocumentHandler(r chi.Router) {
	r.Get("/{id}", documentHandler.getOneById)
	r.Get("/", documentHandler.getAll)
	r.Put("/{id}", documentHandler.updateOneById)
	r.Delete("/", documentHandler.deleteOneById)
	r.Post("/", documentHandler.createOne)
}

func (documentHandler DocumentHandler) getOneById(w http.ResponseWriter, r *http.Request) {

	documentId := chi.URLParam(r, "id")

	document, err := documentHandler.DocumentService.GetOneById(documentId)

	if err != nil {
		err.Response(w)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(utils.Response{Success: true, Data: document})
}

func (documentHandler DocumentHandler) getAll(w http.ResponseWriter, r *http.Request) {

	documents, err := documentHandler.DocumentService.GetAll()
	if err != nil {
		err.Response(w)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(utils.Response{Success: true, Data: documents})

}
func (documentHandler DocumentHandler) createOne(w http.ResponseWriter, r *http.Request) {
	document := make(map[string]interface{})
	json.NewDecoder(r.Body).Decode(&document)

	subject, ok := document["subject"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(utils.Response{Data: "The field 'subject' is NOT FOUND"})
		return
	}

	err := documentHandler.DocumentService.CreateOne(models.Document{Subject: null.String{String: subject.(string), Valid: true}})

	if err != nil {
		err.Response(w)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(utils.Response{Success: true, Data: "Create data successfully"})
}
func (documentHandler DocumentHandler) updateOneById(w http.ResponseWriter, r *http.Request) {

	documentId := chi.URLParam(r, "id")

	document := make(map[string]interface{})
	json.NewDecoder(r.Body).Decode(&document)

	subject, ok := document["subject"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(utils.Response{Data: "The field 'subject' is NOT FOUND"})
		return
	}

	err := documentHandler.DocumentService.UpdateOne(models.Document{DocumentId: documentId, Subject: null.String{String: subject.(string), Valid: true}})
	if err != nil {
		err.Response(w)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(utils.Response{Success: true, Data: "Update data successfully"})

}
func (documentHandler DocumentHandler) deleteOneById(w http.ResponseWriter, r *http.Request) {
	documentId := chi.URLParam(r, "id")
	err := documentHandler.DocumentService.DeleteOneById(documentId)
	if err != nil {
		err.Response(w)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(utils.Response{Success: true, Data: "Delete data successfully"})
}
