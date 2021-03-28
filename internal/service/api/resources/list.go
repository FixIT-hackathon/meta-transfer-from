package resources

import (
	"encoding/json"
	"net/http"
	"time"
)

type ListRequest struct {
}

func NewListRequest(r *http.Request) (*ListRequest, error) {
	var req ListRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}

	return &req, nil
}

type Transfer struct {
	ID        int64  `json:"id"`
	Sender    string `json:"sender"`
	Receiver  string `json:"receiver"`
	Amount    string `json:"amount"`
	Status    string `json:"status"`
	Fee       string `json:"fee"`
	CreatedAt time.Time
}
