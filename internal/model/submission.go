package model

import "time"

type WritingSubmission struct {
	Id               string
	TaskId           string
	UserId           string
	Paragraphs       []string
	SubmitTime       time.Time
	TimeTakenSeconds int
}
