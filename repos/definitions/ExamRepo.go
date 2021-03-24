package definitions

import "github.com/rabatiproject/testing-application-backend/model/base"

type ExamRepo interface {
	SaveExam(*base.Exam) error
}
