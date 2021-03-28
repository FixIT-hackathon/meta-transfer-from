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
	ID       int64  `db:"id" structs:"-"`
	Sender   string `db:"status" structs:"status"`
	Receiver string `db:"tx_hash" structs:"tx_hash"`
	Fee      string `db:"fee" structs:"type"`
	Status   string `db:"sender" structs:"sender"`
	Amount   string `db:"amount" structs:"amount"`

	R        string `db:"r" structs:"r"`
	S        string `db:"r" structs:"s"`
	V        string `db:"r" structs:"v"`

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
