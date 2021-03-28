package signature

import (
	"github.com/ethereum/go-ethereum/common/hexutil"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

type Parameters struct {
	R string `json:"r"`
	S string `json:"s"`
	V int    `json:"v"`
}

func ParseSignatureParameters(signature []byte) (*Parameters, error) {
	if len(signature) != 65 {
		return nil, errors.New("bad signature")
	}
	params := Parameters{}

	params.R = hexutil.Encode(signature[:32])
	params.S = hexutil.Encode(signature[32:64])
	params.V = 27 + int(signature[64])

	return &params, nil
}
