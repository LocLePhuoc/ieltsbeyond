package writing

import "context"

type TaskRepository interface {
	GetAllTasks1(ctx context.Context, cursor int, limit int) ([]Task1, error)
	GetAllTasks2(ctx context.Context, cursor int, limit int) ([]Task2, error)

	GetTask1(ctx context.Context, taskId string) (Task1, error)
	GetTask2(ctx context.Context, taskId string) (Task2, error)

	UpsertTask1(ctx context.Context, task Task1) (Task1, error)
	UpsertTask2(ctx context.Context, task Task2) (Task2, error)
}

type SubmissionRepository interface {
	GetAllSubmissions(ctx context.Context, cursor int, limit int) ([]Submission, error)
	GetSubmission(ctx context.Context, submissionId string) (Submission, error)
	Upsert(ctx context.Context, submission Submission) (Submission, error)
}

type AssessmentRepository interface {
	GetAllAssessments(ctx context.Context, cursor int, limit int) ([]Assessment, error)
	GetAssessment(ctx context.Context, assessmentId string) (Assessment, error)
	Upsert(ctx context.Context, assessment Assessment) (Assessment, error)
}
