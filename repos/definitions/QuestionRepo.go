package definitions

import "github.com/rabatiproject/testing-application-backend/model/base"

type QuestionRepo interface {
	CreateMCQuestion(question *base.MultipleChoiceQuestion) error
	CreateOEQuestion(question *base.OpenEndedQuestion) error
	CreatePQuestion(question *base.ProgrammingQuestion) error
	QuestionExists(questionId string) bool
}
