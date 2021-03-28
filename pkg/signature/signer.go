package signature

import (
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/crypto"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

type Hasher interface {
	Hash() []byte
}

//go:generate moq -pkg mocks -out  ./mocks/signer.go . Signer
type Signer interface {
	Sign(Hasher) (*Parameters, error)
}

type signer struct {
	privKey *ecdsa.PrivateKey
}

func NewSigner(privKey *ecdsa.PrivateKey) Signer {
	return &signer{privKey: privKey}
}

func (s *signer) Sign(hasher Hasher) (*Parameters, error) {
	signature, err := crypto.Sign(hasher.Hash(), s.privKey)
	if err != nil {
		return nil, errors.Wrap(err, "failed to sign hash")
	}

	return ParseSignatureParameters(signature)
}
