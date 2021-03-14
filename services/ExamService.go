package services

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/rabatiproject/testing-application-backend/managers"
	"github.com/rabatiproject/testing-application-backend/model/base"
	"net/http"
	"strconv"
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

func DeleteExam(write http.ResponseWriter, request *http.Request) {
	i, err := strconv.ParseInt(
		mux.Vars(request)["id"],
		10,
		64,
	)

	if err != nil {
		write.WriteHeader(http.StatusBadRequest)
		return
	}
	notFoundError := managers.DeleteExam(i)

	if notFoundError != nil {
		write.WriteHeader(http.StatusNoContent)
		return
	}
}
