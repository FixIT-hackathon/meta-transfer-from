package api

import (
	"github.com/FixIT-hackathon/meta-transfer-from/internal/config"
	"github.com/FixIT-hackathon/meta-transfer-from/internal/data"
	"github.com/FixIT-hackathon/meta-transfer-from/internal/service/api/handlers"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/kit/pgdb"
	"log"
)

func Router(cfg config.Config) chi.Router {
	r := chi.NewRouter()

	db, err := pgdb.Open(pgdb.Opts{
		URL:                "postgres://auction:auction@localhost:5432/auction?sslmode=disable",
		MaxOpenConnections: 15,
		MaxIdleConnections: 15,
	})
	if err != nil {
		log.Fatal(err)
	}

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"https://*", "http://*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		MaxAge:         300,
	}),
		ape.CtxMiddleware(
			handlers.CtxTransfersQ(data.NewTransfers(db)),
			handlers.CtxEthClient(cfg.Client()),
		),
	)

	r.Post("/craft", handlers.Craft)
	r.Post("/push", handlers.Push)
	r.Post("/ethereum", handlers.RequestEthereum)
	r.Get("/list", handlers.List)

	return r
}
