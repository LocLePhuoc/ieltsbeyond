package writing

import "time"

type AssessmentCriteria string

const (
	AssessmentCriteriaTaskAchievement             AssessmentCriteria = "Task Achievement" // or "Task Response" for Task 2
	AssessmentCriteriaCoherenceAndCohesion        AssessmentCriteria = "Coherence and Cohesion"
	AssessmentCriteriaLexicalResource             AssessmentCriteria = "Lexical Resource"
	AssessmentCriteriaGrammaticalRangeAndAccuracy AssessmentCriteria = "Grammatical Range and Accuracy"
)

type criterion struct {
	Band       float64  `json:"band"`
	Strengths  []string `json:"strengths"`
	Weaknesses []string `json:"weaknesses"`
	Comment    string   `json:"comment"`
}

type errorCorrection struct {
	Original    string `json:"original"`
	Correction  string `json:"correction"`
	Type        string `json:"type"`
	Explanation string `json:"explanation"`
}

type Assessment struct {
	Id           string    `json:"id" bson:"_id"`
	TaskId       string    `json:"task_id" bson:"taskId"`
	SubmissionId string    `json:"submission_id" bson:"submissionId"`
	CreatedAt    time.Time `json:"created_at" bson:"createdAt"`
	Criteria     struct {
		TaskAchievement          criterion `json:"task_achievement"`
		CoherenceAndCohesion     criterion `json:"coherence_and_cohesion"`
		LexicalResource          criterion `json:"lexical_resource"`
		GrammaticalRangeAccuracy criterion `json:"grammatical_range_accuracy"`
	} `json:"criteria"`
	DataAccuracyCheck struct {
		CorrectPoints   []string `json:"correct_points"`
		IncorrectPoints []string `json:"incorrect_points"`
	} `json:"data_accuracy_check"`
	ErrorCorrections       []errorCorrection `json:"error_corrections"`
	OverviewFeedback       string            `json:"overview_feedback"`
	TopPrioritiesToImprove []string          `json:"top_priorities_to_improve"`
}
