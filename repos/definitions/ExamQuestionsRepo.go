package definitions

type ExamQuestionRepo interface {
	AddToExam(questionId, examId string) error
	ExamContainsQuestions(questionId string) bool
}
