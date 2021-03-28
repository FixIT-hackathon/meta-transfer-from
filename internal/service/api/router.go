package api

import (
	"github.com/FixIT-hackathon/meta-transfer-from/internal/config"
	"github.com/FixIT-hackathon/meta-transfer-from/internal/data"
	"github.com/FixIT-hackathon/meta-transfer-from/internal/service/api/handlers"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"gitlab.com/distributed_lab/ape"
)

func Router(cfg config.Config) chi.Router {
	r := chi.NewRouter()

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"https://*", "http://*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		MaxAge:         300, // Maximum value not ignored by any of major browsers
	}),
		ape.CtxMiddleware(
			handlers.CtxTransfersQ(data.NewTransfers(cfg.DB())),
		),
	)

	r.Post("/craft", handlers.Craft)
	r.Post("/push", handlers.Push)
	r.Get("/list", handlers.Push)

	return r
}
