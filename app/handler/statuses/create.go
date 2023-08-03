package statuses

import (
	"encoding/json"
	"net/http"

	"yatter-backend-go/app/domain/object"
	"yatter-backend-go/app/handler/auth"
)

// Request body for `POST /v1/statuses`
type AddRequest struct {
	Content string `json:"status"`
}

// Handle request for `POST /v1/statuses`
func (h *handler) Create(w http.ResponseWriter, r *http.Request) {
	//ctx := r.Context()
	var req AddRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	account := auth.AccountOf(r)
	if account == nil {
		http.Error(w, "must be authorized", http.StatusUnauthorized)
		return
	}

	status := new(object.Status)
	status.Content = req.Content
	status.AccountID = account.ID

	// panic("Must Implement Account Registration")

	if err := h.sr.CreateStatus(status); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(status); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
