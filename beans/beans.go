package beans

import (
	"github.com/rabatiproject/testing-application-backend/repos/definitions"
	"github.com/rabatiproject/testing-application-backend/repos/implementation"
)

var (
	ExamRepository         definitions.ExamRepo         = implementation.NewDynamoDbRepo()
	UserRepository         definitions.UserRepo         = implementation.NewUserRepo()
	QuestionRepository     definitions.QuestionRepo     = implementation.NewQuestionRepo()
	ExamQuestionRepository definitions.ExamQuestionRepo = implementation.NewExamQuestionRepo()
)
