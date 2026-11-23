package handler

import (
	"ieltsbeyond/internal/utils"
	"log"
	"net/http"
)

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
