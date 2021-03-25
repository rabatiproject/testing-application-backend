package implementation

import (
	"github.com/google/uuid"
	"github.com/rabatiproject/testing-application-backend/model/base"
)

type questionRepo struct {
	TableName string
}

func NewQuestionRepo() *questionRepo {
	return &questionRepo{
		TableName: "QUESTIONS",
	}
}

func (q *questionRepo) CreateMCQuestion(question *base.MultipleChoiceQuestion) error {
	question.Id = uuid.New().String()
	question.Type = base.QuestionMultipleChoice
	return insert(question, NewQuestionRepo().TableName)
}

func (q *questionRepo) CreateOEQuestion(question *base.OpenEndedQuestion) error {
	question.Id = uuid.New().String()
	question.Type = base.QuestionOpenEnded
	return insert(question, NewQuestionRepo().TableName)
}
func (q *questionRepo) CreatePQuestion(question *base.ProgrammingQuestion) error {
	question.Id = uuid.New().String()
	question.Type = base.QuestionProgramming
	return insert(question, NewQuestionRepo().TableName)
}
