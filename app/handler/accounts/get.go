package accounts

import (
	"encoding/json"
	"net/http"

	"yatter-backend-go/app/domain/object"

	"github.com/go-chi/chi/v5"
)

// Handle request for `GET /v1/accounts/{username}`
func (h *handler) Get(w http.ResponseWriter, r *http.Request) {
	//ctx := r.Context()
	account := new(object.Account)
	account.Username = chi.URLParam(r, "username")

	test, err := h.ar.GetAccount(account)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(test); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
