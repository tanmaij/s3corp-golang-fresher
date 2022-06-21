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

type SubDocumentHandler struct {
	SubDocumentService service.SubDocumentService
}

func (subDocumentHandler SubDocumentHandler) SubDocumentHandler(r chi.Router) {
	r.Get("/{id}", subDocumentHandler.getOneById)
	r.Get("/", subDocumentHandler.getAll)
	r.Put("/", subDocumentHandler.updateOneById)
	r.Delete("/", subDocumentHandler.deleteOneById)
	r.Post("/", subDocumentHandler.createOne)
}
func (subDocumentHandler SubDocumentHandler) getOneById(w http.ResponseWriter, r *http.Request) {

}
func (subDocumentHandler SubDocumentHandler) getAll(w http.ResponseWriter, r *http.Request) {
	subDocuments, err := subDocumentHandler.SubDocumentService.GetAll()
	if err != nil {
		err.Response(w)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(utils.Response{Success: true, Data: subDocuments})
}
func (subDocumentHandler SubDocumentHandler) createOne(w http.ResponseWriter, r *http.Request) {
	document := make(map[string]interface{})
	json.NewDecoder(r.Body).Decode(&document)

	title, ok := document["title"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(utils.Response{Data: "The field 'title' is NOT FOUND"})
		return
	}
	content, ok := document["content"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(utils.Response{Data: "The field 'content' is NOT FOUND"})
		return
	}
	documentId, ok := document["documentId"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(utils.Response{Data: "The field 'documentId' is NOT FOUND"})
		return
	}
	newSubDocument := models.SubDocument{
		Title:      null.String{String: title.(string), Valid: true},
		Content:    null.String{String: content.(string), Valid: true},
		Documentid: null.String{String: documentId.(string), Valid: true}}
	err := subDocumentHandler.SubDocumentService.CreateOne(newSubDocument)

	if err != nil {
		err.Response(w)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(utils.Response{Success: true, Data: "Create data successfully"})
}
func (subDocumentHandler SubDocumentHandler) updateOneById(w http.ResponseWriter, r *http.Request) {

}
func (subDocumentHandler SubDocumentHandler) deleteOneById(w http.ResponseWriter, r *http.Request) {

}
