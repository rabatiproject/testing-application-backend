package services

import (
	"encoding/json"
	"github.com/rabatiproject/testing-application-backend/managers"
	"github.com/rabatiproject/testing-application-backend/model/creator"
	"net/http"
)

func CreateMultipleChoiceQuestion(w http.ResponseWriter, r *http.Request) {
	mcQuestion := &creator.MultipleChoiceQuestionC{}
	json.NewDecoder(r.Body).Decode(mcQuestion)
	managers.CreateMultipleQuestion(mcQuestion)
}
func CreateOpenEndedQuestion(w http.ResponseWriter, r *http.Request) {
	oeQuestion := &creator.OpenEndedQuestionC{}
	json.NewDecoder(r.Body).Decode(oeQuestion)
	managers.CreateOpenEndedQuestion(oeQuestion)
}
func CreateProgrammingQuestion(w http.ResponseWriter, r *http.Request) {
	programmingQuestion := &creator.ProgrammingQuestionC{}
	json.NewDecoder(r.Body).Decode(programmingQuestion)
	managers.CreateProgrammingQuestion(programmingQuestion)
}
