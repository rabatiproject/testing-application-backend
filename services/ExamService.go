package services

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/rabatiproject/testing-application-backend/model/base"
	"log"
	"net/http"
)

func CreateExam(writer http.ResponseWriter, request *http.Request) {
	exam := &base.Exam{}
	json.NewDecoder(request.Body).Decode(exam)
	fmt.Println(exam.Title)
}

func DeleteExam(write http.ResponseWriter, request *http.Request) {
	log.Println("Deleting ", mux.Vars(request))
}
