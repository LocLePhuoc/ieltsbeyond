package writing

import "time"

type Submission struct {
	Id               string
	TaskId           string
	UserId           string
	Paragraphs       []string
	SubmitTime       time.Time
	TimeTakenSeconds int
}
