package beans

import (
	"github.com/rabatiproject/testing-application-backend/repos/definitions"
	"github.com/rabatiproject/testing-application-backend/repos/implmentaions"
)

var (
	ExamRepository definitions.ExamRepo = implmentaions.NewDynamoDbRepo()
	UserRepository definitions.UserRepo = implmentaions.NewUserRepo()
)
