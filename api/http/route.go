package http

import (
	"net/http"

	controller "github.com/codespade/stream-server/api/http/controller"
	"github.com/go-chi/chi"
)

func handler() http.Handler {
	h := chi.NewRouter()

	h.Route("/codespade", func(r chi.Router) {
		r.Post("/verify-hash", controller.VerifyHash)

		r.Post("/block-id", controller.BlockID)
	})

	return h
}
