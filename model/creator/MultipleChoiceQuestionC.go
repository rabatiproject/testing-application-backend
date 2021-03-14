package creator

import (
	"github.com/rabatiproject/testing-application-backend/model/base"
	"os/user"
)

type MultipleChoiceQuestionC struct {
	base.MultipleChoiceQuestion
	CorrectChoices []int64 `json:"correct_choices"`
}

func (receiver *MultipleChoiceQuestionC) GetCreator() *user.User {
	return nil
}
