package creator

import "github.com/rabatiproject/testing-application-backend/model/base"

type MultipleChoiceQuestionC struct {
	base.MultipleChoiceQuestion
	QuestionC
	CorrectChoices []base.Choice
}
