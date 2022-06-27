package handler

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"net/http"
	"s3corp-golang-fresher/internal/models"
	"s3corp-golang-fresher/internal/service"
	"s3corp-golang-fresher/utils"
)

type DocHandler struct {
	DocService service.DocService
}

func NewDocHandler(docService service.DocService) DocHandler {
	return DocHandler{DocService: docService}
}

func (docHandler DocHandler) DocHandler(r chi.Router) {
	r.Get("/{id}", docHandler.GetDocById)
	r.Get("/", docHandler.GetDocs)
	r.Put("/{id}", docHandler.UpdateDoc)
	r.Delete("/", docHandler.DeleteDocById)
	r.Post("/", docHandler.CreateDoc)
}

func (docHandler DocHandler) GetDocById(w http.ResponseWriter, r *http.Request) {

	docId := chi.URLParam(r, "id")

	doc, err := docHandler.DocService.GetDocById(docId)

	if err.StatusCode != http.StatusOK {
		err.Response(w)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(utils.Response{Success: true, Data: doc})
}

func (docHandler DocHandler) GetDocs(w http.ResponseWriter, r *http.Request) {

	documents, err := docHandler.DocService.GetDocs()
	if err.StatusCode != http.StatusOK {
		err.Response(w)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(utils.Response{Success: true, Data: documents})

}
func (docHandler DocHandler) CreateDoc(w http.ResponseWriter, r *http.Request) {
	document := make(map[string]interface{})
	json.NewDecoder(r.Body).Decode(&document)

	subject, ok := document["subject"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(utils.Response{Data: "The field 'subject' is NOT FOUND"})
		return
	}

	err := docHandler.DocService.CreateDoc(models.Document{Subject: subject.(string)})

	if err.StatusCode != http.StatusOK {
		err.Response(w)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(utils.Response{Success: true, Data: "Create data successfully"})
}
func (docHandler DocHandler) UpdateDoc(w http.ResponseWriter, r *http.Request) {
	_ = chi.URLParam(r, "id")
	document := make(map[string]interface{})
	json.NewDecoder(r.Body).Decode(&document)

	_, ok := document["subject"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(utils.Response{Data: "The field 'subject' is NOT FOUND"})
		return
	}

	err := docHandler.DocService.UpdateDoc(models.Document{})

	if err.StatusCode != http.StatusOK {
		err.Response(w)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(utils.Response{Success: true, Data: "Create data successfully"})

}
func (docHandler DocHandler) DeleteDocById(w http.ResponseWriter, r *http.Request) {
	documentId := chi.URLParam(r, "id")
	err := docHandler.DocService.DeleteDocById(documentId)
	if err.StatusCode != http.StatusOK {
		err.Response(w)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(utils.Response{Success: true, Data: "Delete data successfully"})
}
