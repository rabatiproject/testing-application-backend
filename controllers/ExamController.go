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
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
}

func GetExam(writer http.ResponseWriter, request *http.Request) {
	examId := mux.Vars(request)["exam-id"]
	if utils.IsEmpty(examId) {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	exam, err := managers.GetExam(examId)

	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
	} else {
		err2 := json.NewEncoder(writer).Encode(exam)

		if err2 != nil {
			writer.WriteHeader(http.StatusInternalServerError)
		}
	}
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
		writer.WriteHeader(http.StatusInternalServerError)
	}
}
