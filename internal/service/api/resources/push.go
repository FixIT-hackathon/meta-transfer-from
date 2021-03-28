package resources

import (
	"encoding/json"
	"github.com/FixIT-hackathon/meta-transfer-from/pkg/signature"
	signer "github.com/ethereum/go-ethereum/signer/core"
	"net/http"
)

type PushRequest struct {
	signer.TypedData
	Signature signature.Parameters `json:"signature"`
}

func NewPushRequest(r *http.Request) (*CraftRequest, error) {
	var req CraftRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}

	return &req, nil
}
