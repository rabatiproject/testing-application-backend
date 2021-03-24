package managers

import (
	"github.com/rabatiproject/testing-application-backend/model/base"
	"github.com/rabatiproject/testing-application-backend/repos/definitions"
	"github.com/rabatiproject/testing-application-backend/repos/implmentaions"
)

var (
	exams          []*base.Exam
	examRepository definitions.ExamRepo = implmentaions.NewDynamoDbRepo()
)

func CreateNewExam(exam *base.Exam) error {
	return examRepository.SaveExam(exam)
}

func GetExam(id int64) *base.Exam {
	return findExamFromSlice(id)
}

func DeleteExam(id int64) error {

	return nil
}

func findExamFromSlice(id int64) *base.Exam {

	return nil
}
