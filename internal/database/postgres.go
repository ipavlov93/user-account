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

// MustNewPostgresAdapter tries to connect to Postgres database, panics if connection has failed.
// WARNING: sslmode is disabled. It's not recommended use disabled sslmode in production.
// MustNewPostgresAdapter returns PostgresAdapter
func MustNewPostgresAdapter(host string, port int, user, password, dbname string) PostgresAdapter {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	// panics due to failed connection to database
	connection := sqlx.MustConnect("postgres", psqlInfo)

	return PostgresAdapter{sqlxDB: connection}
}

func (db *PostgresAdapter) GetConnection() *sqlx.DB {
	return db.sqlxDB
}

func (db *PostgresAdapter) CloseConnection() error {
	if db.sqlxDB == nil {
		return nil
	}
	return db.sqlxDB.Close()
}

// MustBeginTx tries to begin transaction, panics if begin has failed.
func (db *PostgresAdapter) MustBeginTx(ctx context.Context, options *sql.TxOptions) *sqlx.Tx {
	tx, err := db.sqlxDB.BeginTxx(ctx, options)
	if err != nil {
		panic(fmt.Sprintf("failed to start transaction, %s", err))
	}
	return tx
}

// mustRollbackTx tries to rollback not nil transaction, panics if rollback failed.
func mustRollbackTx(tx *sqlx.Tx) {
	if tx == nil {
		return
	}
	err := tx.Rollback()
	if err != nil {
		panic(fmt.Sprintf("failed to rollback transaction, %s", err))
	}
}

// MustRollbackTxUnlessCommitted commits not nil transaction.
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
