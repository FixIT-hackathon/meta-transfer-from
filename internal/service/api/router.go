package api

import (
	"github.com/FixIT-hackathon/meta-transfer-from/internal/service/api/handlers"
	"github.com/go-chi/chi"
)

func Router() chi.Router {
	r := chi.NewRouter()

	r.Post("/craft", handlers.Craft)

	return r
}
