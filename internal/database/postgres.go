package database

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" //The postgres driver
)

type PostgresAdapter struct {
	sqlxDB *sqlx.DB
}

func NewPostgresAdapter(host string, port int, user, password, dbname string) PostgresAdapter {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	connection, err := sqlx.Connect("postgres", psqlInfo)
	if err != nil {
		panic(fmt.Sprintf("failed to connect to database, %s", err))
	}

	return PostgresAdapter{sqlxDB: connection}
}

func (db *PostgresAdapter) GetConnection() *sqlx.DB {
	return db.sqlxDB
}

// MustBeginTx tries to begin transaction, panics if begin has failed.
func (db *PostgresAdapter) MustBeginTx(ctx context.Context, options *sql.TxOptions) *sqlx.Tx {
	tx, err := db.sqlxDB.BeginTxx(ctx, options)
	if err != nil {
		panic(fmt.Sprintf("failed to start transaction, %s", err))
	}
	return tx
}

// mustRollbackTx tries to rollback transaction, panics if rollback failed.
func mustRollbackTx(tx *sqlx.Tx) {
	err := tx.Rollback()
	if err != nil {
		panic(fmt.Sprintf("failed to rollback transaction, %s", err))
	}
}

// MustRollbackTxUnlessCommitted commits the not nil transaction.
// It tries to rollback tx if commit has failed.
// MustRollbackTxUnlessCommitted name is preferred than MustCommit.
func (db *PostgresAdapter) MustRollbackTxUnlessCommitted(tx *sqlx.Tx) {
	if tx == nil {
		return
	}
	if err := tx.Commit(); err != nil {
		mustRollbackTx(tx)
	}
}
