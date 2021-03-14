package managers

import (
	"github.com/rabatiproject/testing-application-backend/model/base"
	"math/rand"
)

var (
	exams []*base.Exam
)

func CreateNewExam(exam *base.Exam) error {
	if exam.Id == 0 {
		exam.Id = rand.Int63()
	}
	exams = append(exams, exam)
	return nil
}

func GetExam(id int64) *base.Exam {
	return findExamFromSlice(id)
}

func DeleteExam(id int64) error {
	for i, v := range exams {
		if v.Id == id {
			exams = append(exams[:i], exams[i+1:]...)
		}
	}
	return nil
}

func findExamFromSlice(id int64) *base.Exam {
	for _, v := range exams {
		if v.Id == id {
			return v
		}
	}
	return nil
}
