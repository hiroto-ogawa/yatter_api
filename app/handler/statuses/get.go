package statuses

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func (h *handler) Get(w http.ResponseWriter, r *http.Request) {

	statusIDStr := chi.URLParam(r, "id")
	statusID, err := strconv.ParseInt(statusIDStr, 10, 64)
	if err != nil {
		http.Error(w, "Malformed parameter id", http.StatusBadRequest)
		return
	}

	status, err := h.sr.GetStatus(statusID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	if status == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	if err := json.NewEncoder(w).Encode(status); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
