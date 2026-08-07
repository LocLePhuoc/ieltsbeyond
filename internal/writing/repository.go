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
	GetSubmission(ctx context.Context, collectionName string, submissionId string) (*Submission, error)
	UpsertWritingSubmission(ctx context.Context, collectionName string, submission Submission) (*Submission, error)
}

type AssessmentRepository interface {
	GetAssessment(ctx context.Context, collectionName string, assessmentId string) (*Assessment, error)
	UpsertWritingAssessment(ctx context.Context, collectionName string, assessment Assessment) (*Assessment, error)
}
