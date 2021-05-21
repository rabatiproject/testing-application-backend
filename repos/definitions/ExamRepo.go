package definitions

import "github.com/rabatiproject/testing-application-backend/model/base"

type ExamRepo interface {
	SaveExam(*base.Exam) error
	ExamExists(examId string) bool
	GetExam(examId string) *base.Exam
}
