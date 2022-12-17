package http

import (
	"net/http"

	"github.com/go-chi/chi"
)

func handler() http.Handler {
	h := chi.NewRouter()

	h.Route("/codespade", func(r chi.Router) {
		r.Post("/verify-hash", VerifyHash)

		r.Post("/block-id", BlockID)
	})

	return h
}
