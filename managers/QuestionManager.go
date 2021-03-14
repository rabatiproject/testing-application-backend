package managers

import (
	"github.com/rabatiproject/testing-application-backend/model/base"
	"github.com/rabatiproject/testing-application-backend/model/creator"
)

var (
	questions []base.CreatedQuestion
)

func CreateMultipleQuestion(c *creator.MultipleChoiceQuestionC) {
	questions = append(questions, c)
}

func CreateOpenEndedQuestion(c *creator.OpenEndedQuestionC) {
	questions = append(questions, c)
}

func CreateProgrammingQuestion(c *creator.ProgrammingQuestionC) {
	questions = append(questions, c)
}

func AddQuestionToExam(id int64) {

}
