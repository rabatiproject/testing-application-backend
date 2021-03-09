package taker

import "github.com/rabatiproject/testing-application-backend/model/base"

type ExamT struct {
	base.Exam
	Questions []base.Question
	Taker     base.User
}
