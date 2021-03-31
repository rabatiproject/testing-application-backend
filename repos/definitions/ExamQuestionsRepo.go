package definitions

type ExamQuestionRepo interface {
	AddToExam(questionId, examId string) error
	ExamContainsQuestions(examId string, questionId string) bool
}
