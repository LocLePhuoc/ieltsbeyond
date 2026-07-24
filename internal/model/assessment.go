package model

type AssessmentCriteria string

const (
	AssessmentCriteriaTaskAchievement       AssessmentCriteria = "Task Achievement" // or "Task Response" for Task 2
	AssessmentCriteriaCoherenceAndCohesion  AssessmentCriteria = "Coherence and Cohesion"
	AssessmentCriteriaLexicalResource       AssessmentCriteria = "Lexical Resource"
	AssessmentCriteriaGrammaticalRangeAndAccuracy AssessmentCriteria = "Grammatical Range and Accuracy"
)

type Assessment struct {
	Id string
	SubmissionId string
	Criteria AssessmentCriteria
	Comment string
}