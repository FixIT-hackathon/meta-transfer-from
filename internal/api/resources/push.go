package resources

import (
	"encoding/json"
	"github.com/ethereum/go-ethereum/common"
	signer "github.com/ethereum/go-ethereum/signer/core"
	"math/big"
	"net/http"
)

type PushRequest struct {
	signer.TypedData
}

func NewCraftRequest(r *http.Request) (*CraftRequest, error) {
	var req CraftRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}

	return &req, nil
}