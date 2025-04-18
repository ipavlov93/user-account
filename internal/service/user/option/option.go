package option

import (
	"event-calendar/internal/repository"
	"github.com/jmoiron/sqlx"
)

type TxOption struct {
	Tx *sqlx.Tx
}

func ApplyTx[T any](repo repository.WithTx[T], opts *TxOption) T {
	if opts != nil && opts.Tx != nil {
		return repo.WithTx(opts.Tx)
	}
	return repo.(T)
}

type CreateOptions struct {
	TxOption
}

type CreateUserAccountOptions struct {
	TxOption
	AllowDuplicates bool
}
