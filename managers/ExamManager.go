package managers

import (
	"github.com/rabatiproject/testing-application-backend/beans"
	"github.com/rabatiproject/testing-application-backend/model/base"
)

func CreateNewExam(exam *base.Exam) error {
	return beans.ExamRepository.SaveExam(exam)
}
