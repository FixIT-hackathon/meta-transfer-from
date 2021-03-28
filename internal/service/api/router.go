package api

import (
	"github.com/FixIT-hackathon/meta-transfer-from/internal/service/api/handlers"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
)

func Router() chi.Router {
	r := chi.NewRouter()

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	r.Post("/craft", handlers.Craft)

	return r
}
