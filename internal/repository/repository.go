package repository

import "ieltsbeyond/internal/model"

// PostRepository defines the interface for accessing blog posts.
// Implement this interface to swap between file-based and database-backed storage.
type PostRepository interface {
	GetAllPosts() ([]model.Post, error)
	GetPostsByCategory(category model.Category) ([]model.Post, error)
	GetPostBySlug(slug string) (*model.Post, error)
	GetRecentPosts(limit int) ([]model.Post, error)
}

type WritingSubmissionRepository interface {
	GetSubmission(submissionId string) model.WritingSubmission
	UpsertSubmission(submission model.WritingSubmission)
	GetAllSubmissions(userId string) []model.WritingSubmission
}






