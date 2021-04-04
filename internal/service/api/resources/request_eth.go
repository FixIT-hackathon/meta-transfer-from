package resources

import (
	"encoding/json"
	"github.com/ethereum/go-ethereum/common"
	"net/http"
)

type EthRequest struct {
	ERC20  common.Address `json:"erc20"`
	Sender common.Address `json:"sender"`
}

func NewEthRequest(r *http.Request) (*EthRequest, error) {
	var req EthRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}

	return &req, nil
}
