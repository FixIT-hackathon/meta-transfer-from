package resources

import (
	"encoding/json"
	signer "github.com/ethereum/go-ethereum/signer/core"
	"net/http"
)

type PushRequest struct {
	signer.TypedData `json:"data"`

	Signature string `json:"signature_string"`
	Sender    string `json:"sender"`
	Amount    string `json:"amount"`
	Receiver  string `json:"receiver"`
	Fee       string `json:"fee"`
	ERC20     string `json:"erc20"`
}

func NewPushRequest(r *http.Request) (*PushRequest, error) {
	var req PushRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}

	return &req, nil
}
