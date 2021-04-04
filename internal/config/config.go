package config

import (
	"gitlab.com/distributed_lab/kit/kv"
	"gitlab.com/distributed_lab/kit/pgdb"
)

type Config interface {
	pgdb.Databaser
	EthereumClienter
}

type config struct {
	pgdb.Databaser
	EthereumClienter
}

func NewConfig(getter kv.Getter) Config {
	return &config{
		Databaser:        pgdb.NewDatabaser(getter),
		EthereumClienter: newEthereumClienter(getter),
	}
}
