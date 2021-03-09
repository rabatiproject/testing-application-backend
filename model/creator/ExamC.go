package creator

import "github.com/rabatiproject/testing-application-backend/model/base"

type ExamC struct {
	base.Exam
	Questions []QuestionC
}
