package writing

import "time"

type Submission struct {
	Id               string    `json:"id" bson:"_id"`
	TaskId           string    `json:"task_id" bson:"taskId"`
	UserId           string    `json:"user_id" bson:"userId"`
	Paragraphs       []string  `json:"paragraphs" bson:"paragraphs"`
	SubmitTime       time.Time `json:"submit_time" bson:"submitTime"`
	TimeTakenSeconds int       `json:"time_taken_seconds" bson:"timeTakenSeconds"`
}
