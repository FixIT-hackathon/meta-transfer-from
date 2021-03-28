package data

import (
	"gitlab.com/distributed_lab/logan/v3"
	"github.com/FixIT-hackathon/meta-transfer-from/internal/service/api/resources"
	"time"
)

type Transfers interface {
	New() Transfers
	Create(transfer Transfer) (int64, error)
	Select() ([]Transfer, error)
}

type Transfer struct {
	ID          int64  `db:"id" structs:"-"`
	Sender      string `db:"sender" structs:"sender"`
	Receiver    string `db:"receiver" structs:"receiver"`
	Fee         string `db:"fee" structs:"fee"`
	ERC20       string `db:"erc20" structs:"erc20"`
	Amount      string `db:"amount" structs:"amount"`
	Status      string `db:"status" structs:"status"`
	IsCustomFee bool   `db:"is_custom_fee" structs:"is_custom_fee"`

	R string `db:"r" structs:"r"`
	S string `db:"s" structs:"s"`
	V int    `db:"v" structs:"v"`

	CreatedAt *time.Time `db:"created_at" structs:"-"`
}

func (f *Transfer) Resource() *resources.Transfer {
	return &resources.Transfer{
		ID:        f.ID,
		Amount:    f.Amount,
		Sender:    f.Sender,
		Status:    f.Status,
		Receiver:  f.Receiver,
		Fee:       f.Fee,
		R:         f.R,
		S:         f.S,
		V:         f.V,
		CreatedAt: f.CreatedAt,
	}
}

func (f *Transfer) Fields() logan.F {
	return logan.F{
		"id":     f.ID,
		"sender": f.Sender,
		"amount": f.Amount,
		"status": f.Status,
	}
}
