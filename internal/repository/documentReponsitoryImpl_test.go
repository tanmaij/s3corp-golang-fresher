package repository

//
//import (
//	"database/sql"
//	"fmt"
//	"github.com/volatiletech/null/v8"
//	"log"
//	"s3corp-golang-fresher/data"
//	"s3corp-golang-fresher/internal/models"
//	"testing"
//	"time"
//)
//
//var documentRepositoryImpl DocumentRepositoryImpl
//var documentTime, _ = time.Parse(time.RFC3339, "2022-06-21T15:43:48Z")
//var documentTest = models.Document{DocumentId: `29ad5ecc-c8bf-45ae-bea9-dcd0c9b184f1`, Subject: null.String{String: "Test", Valid: true}, CreatedAt: null.Time{Time: documentTime, Valid: true}}
//
//func TestDocumentRepositoryImpl_GetOneById(t *testing.T) {
//	got, _ := documentRepositoryImpl.GetOneById(documentTest.DocumentId)
//	if got.DocumentId != documentTest.DocumentId || got.Subject.String != documentTest.Subject.String {
//		t.Errorf("It is fake!!!")
//	}
//	newDocumentTime, _ := time.Parse(time.RFC3339, got.CreatedAt.Time.String())
//	if newDocumentTime == documentTest.CreatedAt.Time {
//		t.Errorf("It is fake!!!")
//	}
//}
//
//func TestDocumentRepositoryImpl_UpdateOneById(t *testing.T) {
//	old, _ := documentRepositoryImpl.GetOneById(documentTest.DocumentId)
//
//	dataUpdate := models.Document{DocumentId: documentTest.DocumentId, Subject: null.String{String: "Test update", Valid: true}, CreatedAt: old.CreatedAt}
//
//	_, _ = documentRepositoryImpl.UpdateOne(dataUpdate)
//
//	new, _ := documentRepositoryImpl.GetOneById(documentTest.DocumentId)
//
//	if old.Subject.String == new.Subject.String {
//		t.Errorf("Subject field was not changed!!!")
//	}
//	if new.Subject.String != dataUpdate.Subject.String {
//		t.Errorf("Subject field was changed into data which not true!!!")
//	}
//
//	//Reset the data
//	_, _ = documentRepositoryImpl.UpdateOne(*old)
//}
//
//func TestDocumentRepositoryImpl_DeleteOneById(t *testing.T) {
//
//	_, err := documentRepositoryImpl.DeleteOneById(documentTest.DocumentId)
//
//	new, _ := documentRepositoryImpl.GetOneById(documentTest.DocumentId)
//
//	if err != nil {
//		t.Errorf("Delete failed")
//		return
//	}
//	if new != nil {
//		t.Errorf("Delete failed")
//		return
//	}
//
//	//Reset the data
//	_ = documentRepositoryImpl.CreateOne(documentTest)
//}
//
//func TestDocumentRepositoryImpl_InsertOne(t *testing.T) {
//	var documentTestInsert = models.Document{Subject: null.String{String: "Test insert", Valid: true}}
//	err := documentRepositoryImpl.CreateOne(documentTestInsert)
//	_, err = documentRepositoryImpl.DeleteOneById(documentTest.DocumentId)
//	if err != nil {
//		t.Errorf("Insert failed")
//		return
//	}
//}
//func TestMain(m *testing.M) {
//	fmt.Println("Connecting Database")
//
//	data := data.Data{}
//
//	dsn := fmt.Sprintf("host=localhost user=mai password=1 dbname=researchdocument port=5432 sslmode=disable TimeZone=ASIA/HO_CHI_MINH")
//
//	db, err := sql.Open("postgres", dsn)
//
//	if err != nil {
//		log.Fatalln("Can't connect to database", err)
//	} else {
//		data.DB = db
//	}
//
//	documentRepositoryImpl = DocumentRepositoryImpl{&data}
//
//	m.Run()
//}
