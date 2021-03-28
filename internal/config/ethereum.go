package config

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"gitlab.com/distributed_lab/figure"
	"gitlab.com/distributed_lab/kit/comfig"
	"gitlab.com/distributed_lab/kit/kv"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"reflect"
)

type EthereumClienter interface {
	Client() *ethclient.Client
}

type ethereum struct {
	once comfig.Once
	getter kv.Getter
}

func newEthereumClienter(getter kv.Getter) EthereumClienter {
	return &ethereum{getter: getter}
}


func (e *ethereum) Client() *ethclient.Client {
	return e.once.Do(func() interface{} {
		var cfg struct{
			RPC string `figure:"rpc"`
		}
		err := figure.
			Out(&cfg).
			With(figure.BaseHooks, hooks).
			From(kv.MustGetStringMap(e.getter, "etherem")).
			Please()
		if err != nil {
			panic(errors.Wrap(err, "failed to figure out signer"))
		}

		client, err := ethclient.Dial(cfg.RPC)
		if err != nil {
			panic(fmt.Sprintf("failed to dial rpc %s: %s", cfg.RPC, err.Error()))
		}

		return client
	}).(*ethclient.Client)
}


var hooks = figure.Hooks{
	"common.Address": func(value interface{}) (reflect.Value, error) {
		switch v := value.(type) {
		case string:
			if !common.IsHexAddress(v) {
				// provide value does not look like valid address
				return reflect.Value{}, errors.New("invalid address")
			}
			return reflect.ValueOf(common.HexToAddress(v)), nil
		default:
			return reflect.Value{}, fmt.Errorf("unsupported conversion from %T", value)
		}
	},

	"*ecdsa.PrivateKey": func(value interface{}) (reflect.Value, error) {
		switch v := value.(type) {
		case string:
			privKey, err := crypto.HexToECDSA(v)
			if err != nil {
				return reflect.Value{}, errors.Wrap(err, "invalid hex private key")
			}
			return reflect.ValueOf(privKey), nil
		default:
			return reflect.Value{}, fmt.Errorf("unsupported conversion from %T", value)
		}
	},
}
