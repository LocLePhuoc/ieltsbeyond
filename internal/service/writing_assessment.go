package service

import (
	"context"
	"encoding/json"
	"fmt"
	"ieltsbeyond/internal/llm"
	"ieltsbeyond/internal/storage"
	"ieltsbeyond/internal/utils"
	"ieltsbeyond/internal/writing"
	"net/http"
	"strings"
	"time"

	"github.com/google/uuid"
)

type WritingAssessService struct {
	taskRepo     writing.TaskRepository
	storage      *storage.Client
	provider     llm.Provider
	systemPrompt string
}

func NewWritingAssessmentService(taskRepo writing.TaskRepository, storage *storage.Client, provider llm.Provider, systemPrompt string) *WritingAssessService {
	return &WritingAssessService{taskRepo: taskRepo, storage: storage, provider: provider, systemPrompt: systemPrompt}
}

func (s *WritingAssessService) AssessTask1(ctx context.Context, submission writing.Submission) (*writing.Assessment, error) {
	task, err := s.taskRepo.GetTask1(ctx, submission.TaskId)
	if err != nil {
		return nil, fmt.Errorf("Get task 1 failed for task: %s, %w", submission.TaskId, err)
	}
	var raw []byte
	if parts := strings.SplitN(task.ImageKey, "/", 2); len(parts) == 2 {
		bucket, key := parts[0], parts[1]
		raw, err = s.storage.GetObjectAsBytes(ctx, bucket, key)
		if err != nil {
			return nil, err
		}
	}
	if raw == nil {
		return nil, fmt.Errorf("No image found for task 1: %s", submission.TaskId)
	}

	mediaType := http.DetectContentType(raw)
	if !strings.HasPrefix(mediaType, "image/") {
		return nil, fmt.Errorf("Saved object is not image, found: %s", mediaType)
	}
	parts := []llm.UserPart{
		llm.TextPart(fmt.Sprintf("ĐỀ BÀI:\n%s\n\nHÌNH ẢNH BIỂU ĐỒ:", task.Question)),
		llm.ImagePart(raw, mediaType),
		llm.TextPart(fmt.Sprintf("\nBÀI VIẾT CỦA THÍ SINH:\n%s", strings.Join(submission.Paragraphs, "\n"))),
	}

	text, err := s.provider.Complete(ctx, s.systemPrompt, parts)
	if err != nil {
		return nil, err
	}

	jsonText := utils.ExtractLLMJSON(text)
	var result writing.Assessment
	if err := json.Unmarshal([]byte(jsonText), &result); err != nil {
		return nil, fmt.Errorf("JSON Parse failed: %w ---- model returned ---\n%s", err, jsonText)
	}
	result.Id = uuid.New().String()
	result.TaskId = submission.TaskId
	result.SubmissionId = submission.Id
	result.CreatedAt = time.Now().UTC()
	return &result, nil
}
