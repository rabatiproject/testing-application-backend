package creator

import (
	"github.com/rabatiproject/testing-application-backend/model/base"
	"os/user"
)

type OpenEndedQuestionC struct {
	base.OpenEndedQuestion
	Answer string
}

func (receiver *OpenEndedQuestionC) GetCreator() *user.User {
	return nil
}
