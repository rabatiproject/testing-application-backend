package base

type MultipleChoiceQuestion struct {
	Question
	Choices []Choice `json:"choices"`
}
