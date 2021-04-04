package service

import (
	"github.com/FixIT-hackathon/meta-transfer-from/internal/config"
	"github.com/FixIT-hackathon/meta-transfer-from/internal/service/api"
	"net"
	"net/http"
)

type Service struct{}

func (s Service) Run(cfg config.Config) error {
	r := api.Router(cfg)

	listener, err := net.Listen("tcp", ":8011")
	if err != nil {
		panic(err)
	}

	err = http.Serve(listener, r)
	if err != nil {
		return err
	}

	return nil
}
