package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/rabatiproject/testing-application-backend/managers"
	"github.com/rabatiproject/testing-application-backend/model/base"
	"net/http"
)

func CreateExam(writer http.ResponseWriter, request *http.Request) {
	exam := &base.Exam{}
	json.NewDecoder(request.Body).Decode(exam)
	creationError := managers.CreateNewExam(exam)

	if creationError != nil {
		writer.WriteHeader(http.StatusNoContent)
		return
	}
}

func AttachQuestionToExam(writer http.ResponseWriter, request *http.Request) {
	fmt.Println(mux.Vars(request))
}
