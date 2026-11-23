package handler

import (
	"fmt"
	"ieltsbeyond/internal/repository/mongodb"
	"ieltsbeyond/internal/repository/postgres"
	"ieltsbeyond/internal/service"
	"ieltsbeyond/internal/storage"
	"net/http"
	"strconv"
)

type WritingTaskHandler struct {
	db             postgres.WritingTaskRepository
	submissionRepo *mongodb.SubmissionRepository
	assessmentRepo *mongodb.AssessmentRepository
	assessService  *service.WritingAssessService
	objectStorage  *storage.Client
}

func NewWritingTaskHandler(db postgres.WritingTaskRepository,
	submissionRepo *mongodb.SubmissionRepository,
	assessmentRepo *mongodb.AssessmentRepository,
	assessService *service.WritingAssessService,
	storage *storage.Client) *WritingTaskHandler {

	return &WritingTaskHandler{
		db:             db,
		submissionRepo: submissionRepo,
		assessmentRepo: assessmentRepo,
		assessService:  assessService,
		objectStorage:  storage,
	}
}

const defaultTaskLimit = 20

func getLimitParam(r *http.Request) (int, error) {
	limit := defaultTaskLimit
	if limitParam := r.URL.Query().Get("limit"); limitParam != "" {
		parsed, err := strconv.Atoi(limitParam)
		if err != nil || parsed < 0 {
			return limit, fmt.Errorf("Invalid limit")
		}
		limit = parsed
	}
	return limit, nil
}

func getCursorParam(r *http.Request) (int, error) {
	cursor := 0
	if cursorParam := r.URL.Query().Get("cursor"); cursorParam != "" {
		parsed, err := strconv.Atoi(cursorParam)
		if err != nil || parsed < 0 {
			return cursor, fmt.Errorf("Invalid cursor")
		}
		cursor = parsed
	}
	return cursor, nil
}
