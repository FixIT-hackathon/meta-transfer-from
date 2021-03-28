package signature

import (
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"gitlab.com/distributed_lab/figure"
	"gitlab.com/distributed_lab/kit/comfig"
	"gitlab.com/distributed_lab/kit/kv"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"reflect"
)

type Signerer interface {
	Signer() Signer
}

type config struct {
	PrivKey *ecdsa.PrivateKey `fig:"eth_signer,required"`
}

type signerer struct {
	signerOnce comfig.Once
	getter     kv.Getter
}

func NewSignerer(getter kv.Getter) Signerer {
	return &signerer{
		getter: getter,
	}
}

func (s *signerer) Signer() Signer {
	return s.signerOnce.Do(func() interface{} {
		var cfg config

		err := figure.
			Out(&cfg).
			With(figure.BaseHooks, hooks).
			From(kv.MustGetStringMap(s.getter, "signer")).
			Please()
		if err != nil {
			panic(errors.Wrap(err, "failed to figure out signer"))
		}

		return NewSigner(cfg.PrivKey)
	}).(Signer)
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
