package service

import (
	"github.com/FixIT-hackathon/meta-transfer-from/internal/service/api"
	"net"
	"net/http"
)

type Service struct{}

func (s Service) Run() error {
	r := api.Router()

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
