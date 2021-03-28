package handlers

import (
	"encoding/json"
	"github.com/FixIT-hackathon/meta-transfer-from/internal/service/api/resources"
	"github.com/FixIT-hackathon/meta-transfer-from/pkg/responses"
	"github.com/ethereum/go-ethereum/common/math"
	signer "github.com/ethereum/go-ethereum/signer/core"
	"github.com/google/jsonapi"
	"math/big"
	"net/http"
)

func Craft(w http.ResponseWriter, r *http.Request) {
	req, err := resources.NewCraftRequest(r)
	if err != nil {
		errorsj := []*jsonapi.ErrorObject{{
			Title:  http.StatusText(http.StatusBadRequest),
			Detail: err.Error(),
		}}
		responses.WriteError(w, http.StatusBadRequest, errorsj)
		return
	}

	bchainID, _ := new(big.Int).SetString(req.ChainID, 10)

	hchainID := math.HexOrDecimal256(*bchainID)

	resp := resources.CraftResponse{
		signer.TypedData{
			Types: signer.Types{
				"TransferFrom": []signer.Type{
					{Name: "receiver", Type: "address"},
					{Name: "amount", Type: "uint256"},
					{Name: "fee", Type: "uint256"},
					{Name: "erc20", Type: "address"},
					{Name: "nonce", Type: "uint256"},
				},
				"EIP712Domain": []signer.Type{
					{Name: "name", Type: "string"},
					{Name: "chainId", Type: "uint256"},
					{Name: "verifyingContract", Type: "address"},
				},
			},
			PrimaryType: "TransferFrom",
			Domain: signer.TypedDataDomain{
				Name:    "RelayerEIP712",
				ChainId: &hchainID,
				VerifyingContract: "0x414e1508153Ff4Eb4F22919C828db9E1715ffDaF",
			},

			Message: signer.TypedDataMessage{
				"receiver": req.Receiver,
				"amount":   req.Amount,
				"fee":      100,
				"erc20":    req.ERC20,
				"nonce":    req.Nonce,
			},
		},
	}

	json.NewEncoder(w).Encode(resp.Map())
}
