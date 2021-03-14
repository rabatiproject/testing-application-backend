package creator

import (
	"github.com/rabatiproject/testing-application-backend/model/base"
	"os/user"
)

type ProgrammingQuestionC struct {
	base.ProgrammingQuestion
	Answer string
}

func (receiver *ProgrammingQuestionC) GetCreator() *user.User {
	return nil
}
