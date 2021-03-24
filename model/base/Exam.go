package base

type Exam struct {
	Id                string `json:"id"`
	StartTime         int64  `json:"start_time"`
	EndTime           int64  `json:"end_time"`
	DurationInMinutes int64  `json:"duration_in_minutes"`
	Creator           User   `json:"creator"`
	Title             string `json:"title"`
}
