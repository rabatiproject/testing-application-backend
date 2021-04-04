package controllers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/rabatiproject/testing-application-backend/managers"
	"github.com/rabatiproject/testing-application-backend/model/base"
	"github.com/rabatiproject/testing-application-backend/utils"
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

func GetExam(writer http.ResponseWriter, request *http.Request) {

}

func AttachQuestionToExam(writer http.ResponseWriter, request *http.Request) {

	pathVariables := mux.Vars(request)
	examId := pathVariables["exam-id"]
	questionId := pathVariables["question-id"]

	if utils.IsEmpty(examId) ||
		utils.IsEmpty(questionId) {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	err := managers.AttachQuestionToExam(examId, questionId)

	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
	}
}
