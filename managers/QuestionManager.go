package managers

import (
	"github.com/rabatiproject/testing-application-backend/beans"
	"github.com/rabatiproject/testing-application-backend/model/base"
)

func CreateMultipleQuestion(c *base.MultipleChoiceQuestion) {
	beans.QuestionRepository.CreateMCQuestion(c)
}

func CreateOpenEndedQuestion(c *base.OpenEndedQuestion) {
	beans.QuestionRepository.CreateOEQuestion(c)
}
func CreateProgrammingQuestion(c *base.ProgrammingQuestion) {
	beans.QuestionRepository.CreatePQuestion(c)
}
