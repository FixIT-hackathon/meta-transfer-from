package resources

import (
	"net/http"
	"time"
)

type ListRequest struct {
}

func NewListRequest(r *http.Request) (*ListRequest, error) {
	var req ListRequest

	return &req, nil
}

type Transfer struct {
	ID        int64      `json:"id"`
	Sender    string     `json:"sender"`
	Receiver  string     `json:"receiver"`
	Amount    string     `json:"amount"`
	Status    string     `json:"status"`
	Fee       string     `json:"fee"`
	CreatedAt *time.Time `json:"created_at,omitempty"`

	R string `json:"r"`
	S string `json:"s"`
	V int    `json:"v"`
}
