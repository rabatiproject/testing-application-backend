package creator

import "github.com/rabatiproject/testing-application-backend/model/base"

type OpenEndedQuestionC struct {
	base.OpenEndedQuestion
	QuestionC
	Answer string
}
