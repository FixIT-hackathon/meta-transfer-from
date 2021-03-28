package resources

import (
	"encoding/json"
	"github.com/ethereum/go-ethereum/common"
	signer "github.com/ethereum/go-ethereum/signer/core"
	"net/http"
)

type CraftRequest struct {
	Sender   common.Address `json:"sender"`
	Receiver common.Address `json:"receiver"`
	ERC20    common.Address `json:"erc20"`
	Nonce    string         `json:"nonce"`
	ChainID  string         `json:"chain_id"`
	Amount   string         `json:"amount"`
	Fee      string         `json:"fee"`
}

func NewCraftRequest(r *http.Request) (*CraftRequest, error) {
	var req CraftRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}

	return &req, nil
}

type CraftResponse struct {
	signer.TypedData
}
