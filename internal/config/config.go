package config

import (
	"gitlab.com/distributed_lab/kit/comfig"
	"gitlab.com/distributed_lab/kit/kv"
	"gitlab.com/distributed_lab/kit/pgdb"
)

type Config interface {
	comfig.Logger
	comfig.Listenerer
	pgdb.Databaser
	EthereumClienter
}

type config struct {
	comfig.Logger
	comfig.Listenerer
	pgdb.Databaser
	EthereumClienter
}

func NewConfig(getter kv.Getter) Config {
	return &config{
		Logger:           comfig.NewLogger(getter, comfig.LoggerOpts{}),
		Listenerer:       comfig.NewListenerer(getter),
		Databaser:        pgdb.NewDatabaser(getter),
		EthereumClienter: newEthereumClienter(getter),
	}
}
