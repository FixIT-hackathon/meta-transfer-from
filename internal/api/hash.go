package api

import (
	"github.com/FixIT-hackathon/meta-transfer-from/internal/api/resources"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	signer "github.com/ethereum/go-ethereum/signer/core"
	sha3 "github.com/miguelmota/go-solidity-sha3"
	"math/big"
)

func HashRequest(req resources.TransferFromRequest) string {
	return hexutil.Encode(sha3.SoliditySHA3(
		sha3.Uint8(req.),
		sha3.Address(req.Sender),
		sha3.String(m.TxHash),
		sha3.Uint256(m.Amount),
	))
}


func HashRequest{

	typedDataHash, _ := signerData.HashStruct(signerData.PrimaryType, signerData.Message)
	domainSeparator, _ := signerData.HashStruct("EIP712Domain", signerData.Domain.Map())

	rawData := []byte(fmt.Sprintf("\x19\x01%s%s", string(domainSeparator), string(typedDataHash)))
	challengeHash := crypto.Keccak256Hash(rawData)
}