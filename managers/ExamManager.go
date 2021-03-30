package managers

import (
	"errors"
	"github.com/rabatiproject/testing-application-backend/beans"
	"github.com/rabatiproject/testing-application-backend/model/base"
	"github.com/rabatiproject/testing-application-backend/utils"
)

func CreateNewExam(exam *base.Exam) error {
	return beans.ExamRepository.SaveExam(exam)
}

func AttachQuestionToExam(examId, questionId string) error {
	if utils.IsEmpty(examId) || utils.IsEmpty(questionId) {
		return errors.New("Empty values")
	}

	return nil
}
