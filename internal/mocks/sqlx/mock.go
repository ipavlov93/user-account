package sqlx

import "github.com/jmoiron/sqlx"

type ExtContext interface {
	sqlx.ExtContext
}
