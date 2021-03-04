package base

type Exam struct {
	Id                int64
	StartTime         int64
	EndTime           int64
	DurationInMinutes int64
	Questions         []Question
	Creator           Teacher
	Title             string
}
