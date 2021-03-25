package controllers

import (
	"encoding/json"
	"github.com/rabatiproject/testing-application-backend/managers"
	"github.com/rabatiproject/testing-application-backend/model/base"
	"net/http"
)

func CreateMultipleChoiceQuestion(w http.ResponseWriter, r *http.Request) {
	mcQuestion := &base.MultipleChoiceQuestion{}
	json.NewDecoder(r.Body).Decode(mcQuestion)
	managers.CreateMultipleQuestion(mcQuestion)
}
func CreateOpenEndedQuestion(w http.ResponseWriter, r *http.Request) {
	oeQuestion := &base.OpenEndedQuestion{}
	json.NewDecoder(r.Body).Decode(oeQuestion)
	managers.CreateOpenEndedQuestion(oeQuestion)
}
func CreateProgrammingQuestion(w http.ResponseWriter, r *http.Request) {
	programmingQuestion := &base.ProgrammingQuestion{}
	json.NewDecoder(r.Body).Decode(programmingQuestion)
	managers.CreateProgrammingQuestion(programmingQuestion)
}
