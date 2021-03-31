package managers

import (
	"errors"
	"github.com/rabatiproject/testing-application-backend/beans"
	"github.com/rabatiproject/testing-application-backend/model/base"
)

func CreateNewExam(exam *base.Exam) error {
	return beans.ExamRepository.SaveExam(exam)
}

func AttachQuestionToExam(examId, questionId string) error {
	examExists := beans.ExamRepository.ExamExists(examId)

	if !examExists {
		return errors.New("EXAM DOES NOT EXIST")
	}

	questionExists := beans.QuestionRepository.QuestionExists(questionId)

	if !questionExists {
		return errors.New("QUESTION DOES NO EXIST")
	}

	//examContainsQuestion := beans.ExamQuestionRepository.ExamContainsQuestions(questionId)
	//
	//if examContainsQuestion {
	//	return errors.New("EXAM ALREADY CONTAINS THE QUESTION")
	//} else {
	return beans.ExamQuestionRepository.AddToExam(questionId, examId)
	//}
	//return nil
}
