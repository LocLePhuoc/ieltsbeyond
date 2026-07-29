package handler

import (
	"fmt"
	"ieltsbeyond/internal/postgres"
	"ieltsbeyond/internal/utils"
	"log"
	"net/http"
	"strconv"
)

type WritingTaskHandler struct {
	db postgres.WritingTaskRepository
}

func NewWritingTaskHandler(db postgres.WritingTaskRepository) *WritingTaskHandler {
	return &WritingTaskHandler{db: db}
}

const defaultTask1Limit = 20

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

func getLimitParam(r *http.Request) (int, error) {
	limit := defaultTask1Limit
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
