package config

import (
	"gitlab.com/distributed_lab/kit/kv"
	"gitlab.com/distributed_lab/kit/pgdb"
)

type Config interface {
	pgdb.Databaser
}

type config struct {
	pgdb.Databaser
}

func NewConfig(getter kv.Getter) Config {
	return &config{
		Databaser:        pgdb.NewDatabaser(getter),
	}
}
