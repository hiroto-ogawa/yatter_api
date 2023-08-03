package statuses

import (
	"net/http"
	"yatter-backend-go/app/domain/repository"
	"yatter-backend-go/app/handler/auth"

	"github.com/go-chi/chi/v5"
)

// Implementation of handler
type handler struct {
	sr repository.Status
}

// Create Handler for `/v1/accounts/`
func NewRouter(ar repository.Account, sr repository.Status) http.Handler {
	r := chi.NewRouter()

	h := &handler{sr}

	// 認証が必要なルート
	r.Group(func(r chi.Router) {
		r.Use(auth.Middleware(ar))
		r.Post("/", h.Create)
	})

	// 認証をスキップするルート
	r.Group(func(r chi.Router) {
		r.Get("/{id}", h.Get)
	})

	return r
}
