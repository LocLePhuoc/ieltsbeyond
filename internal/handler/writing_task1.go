package handler

import (
	"encoding/json"
	logger "ieltsbeyond/internal/logging"
	"ieltsbeyond/internal/middleware"
	"ieltsbeyond/internal/utils"
	"ieltsbeyond/internal/writing"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

// TODO: fill in the Mongo collection name used for Task 1 submissions
const task1SubmissionCollection = "writing_task1"
const task1AssessmentCollection = "writing_task1_assessments"

type SubmitTask1request struct {
	TaskId string `json:"task_id"`
	Answer string `json:"answer"`
}

func (wh *WritingTaskHandler) HandlerGetAllTask1(w http.ResponseWriter, r *http.Request) {
	limit, err := getLimitParam(r)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, "invalid limit")
		return
	}
	cursor, err := getCursorParam(r)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, "invalid cursor")
		return
	}
	tasks1, err := wh.db.GetAllTasks1(r.Context(), cursor, limit)
	if err != nil {
		log.Printf("Error fetching task1 list: %v", err)
		utils.WriteError(w, http.StatusInternalServerError, "failed to load tasks")
		return
	}
	utils.WriteJSON(w, http.StatusOK, tasks1)
}

func (wh *WritingTaskHandler) HandlerGetTask1(w http.ResponseWriter, r *http.Request) {
	taskId := chi.URLParam(r, "id")
	task, err := wh.db.GetTask1(r.Context(), taskId)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, "Cannot get Writing Task 1")
		return
	}
	if parts := strings.SplitN(task.ImageKey, "/", 2); len(parts) == 2 {
		bucket, key := parts[0], parts[1]
		signedURL, err := wh.objectStorage.GetObjectURL(r.Context(), bucket, key)
		if err != nil {
			log.Printf("Cannot get presigned URL: %v", err)
		} else {
			task.ImageKey = signedURL
		}
	}
	utils.WriteJSON(w, http.StatusOK, task)
}

func (wh *WritingTaskHandler) HandlerSubmitTask1(w http.ResponseWriter, r *http.Request) {
	userId, ok := middleware.UserIdFromContext(r)
	if !ok {
		utils.WriteError(w, http.StatusUnauthorized, "unauthorized")
		return
	}
	var req SubmitTask1request
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.WriteError(w, http.StatusBadRequest, "invalid request body")
		return
	}
	if req.Answer == "" {
		utils.WriteError(w, http.StatusBadRequest, "answer is required")
		return
	}

	submission := writing.Submission{
		Id:         uuid.New().String(),
		UserId:     userId,
		TaskId:     req.TaskId,
		Paragraphs: strings.Split(req.Answer, "\n"),
		SubmitTime: time.Now().UTC(),
	}

	saved, err := wh.submissionRepo.UpsertWritingSubmission(r.Context(), task1SubmissionCollection, submission)
	if err != nil {
		logger.Instance.Errorw("Error saving submission.", "error", err)
		utils.WriteError(w, http.StatusInternalServerError, "failed to save submission")
		return
	}
	utils.WriteJSON(w, http.StatusOK, saved)
}

func (wh *WritingTaskHandler) HandlerAssessTask1(w http.ResponseWriter, r *http.Request) {
	submissionId := r.URL.Query().Get("submission_id")
	if submissionId == "" {
		utils.WriteError(w, http.StatusUnauthorized, "invalid submission id")
		return
	}
	submission, err := wh.submissionRepo.GetSubmission(r.Context(), task1SubmissionCollection, submissionId)
	if err != nil {
		logger.Instance.Errorw("Error getting submission.", "error", err)
		utils.WriteError(w, http.StatusInternalServerError, "failed to get submission")
		return
	}
	assessment, err := wh.assessService.AssessTask1(r.Context(), *submission)
	if err != nil {
		logger.Instance.Errorw("Failed to assess task1.", "error", err)
		utils.WriteError(w, http.StatusInternalServerError, "failed to assess task1")
		return
	}

	if _, err := wh.assessmentRepo.UpsertWritingAssessment(r.Context(), task1AssessmentCollection, *assessment); err != nil {
		logger.Instance.Errorw("Error saving assessment.", "error", err)
	}

	utils.WriteJSON(w, http.StatusOK, assessment)
}
