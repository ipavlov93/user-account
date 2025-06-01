package database

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
)

type DBAdapter interface {
	GetConnection() *sqlx.DB
	CloseConnection() error
	MustBeginTx(ctx context.Context, options *sql.TxOptions) *sqlx.Tx
	// MustRollbackTxUnlessCommitted name is preferred than MustCommit
	MustRollbackTxUnlessCommitted(tx *sqlx.Tx)
}

type DBMigrator interface {
	CloseConnection() error
	Up(dialect, migrationDirPath string) error
	//Down(dialect, migrationDirPath string) error
}
