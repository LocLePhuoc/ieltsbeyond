package handler

import (
	"encoding/json"
	"fmt"
	"ieltsbeyond/internal/llm"
	"ieltsbeyond/internal/middleware"
	"ieltsbeyond/internal/repository/mongodb"
	"ieltsbeyond/internal/repository/postgres"
	"ieltsbeyond/internal/storage"
	"ieltsbeyond/internal/utils"
	"ieltsbeyond/internal/writing"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

// TODO: fill in the Mongo collection name used for Task 1 submissions
const task1SubmissionCollection = "writing_task1"

type WritingTaskHandler struct {
	db             postgres.WritingTaskRepository
	submissionRepo *mongodb.SubmissionRepository
	llmService     *llm.Provider
}

func NewWritingTaskHandler(db postgres.WritingTaskRepository, submissionRepo *mongodb.SubmissionRepository, llmService *llm.Provider) *WritingTaskHandler {
	return &WritingTaskHandler{db: db, submissionRepo: submissionRepo, llmService: llmService}
}

const defaultTaskLimit = 20

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

func (wh *WritingTaskHandler) HandlerGetAllTask2(w http.ResponseWriter, r *http.Request) {
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
	tasks2, err := wh.db.GetAllTasks2(r.Context(), cursor, limit)
	if err != nil {
		log.Printf("Error fetching task2 list: %v", err)
		utils.WriteError(w, http.StatusInternalServerError, "failed to load tasks")
		return
	}
	utils.WriteJSON(w, http.StatusOK, tasks2)
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
		signedURL, err := storage.Instance.GetObjectURL(r.Context(), bucket, key)
		if err != nil {
			log.Printf("Cannot get presigned URL: %v", err)
		} else {
			task.ImageKey = signedURL
		}
	}
	utils.WriteJSON(w, http.StatusOK, task)
}

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

type SubmitTask1request struct {
	TaskId string `json:"task_id"`
	Answer string `json:"answer"`
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
		log.Printf("Error saving submission: %v", err)
		utils.WriteError(w, http.StatusInternalServerError, "failed to save submission")
		return
	}
	utils.WriteJSON(w, http.StatusOK, saved)
}
