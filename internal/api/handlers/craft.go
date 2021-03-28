package handlers

import (
	"encoding/json"
	"github.com/FixIT-hackathon/meta-transfer-from/internal/api/resources"
	"github.com/FixIT-hackathon/meta-transfer-from/pkg/responses"
	"github.com/ethereum/go-ethereum/common/math"
	signer "github.com/ethereum/go-ethereum/signer/core"
	"github.com/google/jsonapi"
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

	chainID := math.HexOrDecimal256(req.ChainID)

	json.NewEncoder(w).Encode(&resources.CraftResponse{
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
					{Name: "Name", Type: "string"},
					{Name: "chainId", Type: "uint256"},
					{Name: "ve"},
				},
			},
			PrimaryType: "Challenge",
			Domain: signer.TypedDataDomain{
				Name:    "RelayerEIP712",
				ChainId: &chainID,
			},

			Message: signer.TypedDataMessage{
				"receiver": req.Receiver,
				"amount":   req.Amount,
				"fee":      req.Fee,
				"erc20":    req.ERC20,
				"nonce":    req.Nonce,
			},
		},
	})
}
