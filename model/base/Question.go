package base

type QuestionType int

const (
	QuestionOpenEnded QuestionType = iota
	QuestionMultipleChoice
	QuestionProgramming
)

type Question struct {
	Id        string       `json:"id"`
	Type      QuestionType `json:"type"`
	Direction string       `json:"direction"`
	Images    []string     `json:"images"`
}
