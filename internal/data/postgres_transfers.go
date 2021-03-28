package data

import (
	"fmt"
	"sync"

	sq "github.com/Masterminds/squirrel"
	"github.com/fatih/structs"
	"gitlab.com/distributed_lab/kit/pgdb"
)

const (
	transfersTable = "transfers"
)


var transfersSelect = sq.Select(fmt.Sprintf("%s.*", transfersTable)).From(transfersTable)

type transfers struct {
	db           *pgdb.DB
	joinLotsOnce sync.Once
	sql          sq.SelectBuilder
}

func NewTransfers(db *pgdb.DB) Transfers {
	return &transfers{
		db:  db,
		sql: transfersSelect,
	}
}

func (q *transfers) New() Transfers {
	return NewTransfers(q.db.Clone())
}

func (q *transfers) Create(transfer Transfer) (int64, error) {
	clauses := structs.Map(transfer)

	var id int64
	stmt := sq.Insert(transfersTable).SetMap(clauses).Suffix("returning id")
	err := q.db.Get(&id, stmt)

	return id, err
}

//func (q *transfers) SetStatus(nonce string, status string) error {
//	stmt := sq.Update(transfersTable).
//		Where("tx_hash = ?", nonce).
//		Set("status", status)
//	res, err := q.db.ExecWithResult(stmt)
//	if err != nil {
//		return errors.Wrap(err, "unable to update row")
//	}
//	rowsAffected, err := res.RowsAffected()
//	if err != nil {
//		return errors.Wrap(err, "unable to get affected rows")
//	}
//	if rowsAffected == 0 {
//		return ErrNotFound
//	}
//	return nil
//}

func (q *transfers) Select() ([]Transfer, error) {
	var result []Transfer
	err := q.db.Select(&result, q.sql)
	return result, err
}
