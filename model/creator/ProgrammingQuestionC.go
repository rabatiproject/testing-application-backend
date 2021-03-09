package creator

import (
	"github.com/rabatiproject/testing-application-backend/model/base"
)

type ProgrammingQuestionC struct {
	base.ProgrammingQuestion
	QuestionC
	Answer string
}
