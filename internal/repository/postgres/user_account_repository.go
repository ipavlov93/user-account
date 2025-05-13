// Package postgres is part of repository.
// Package provides implementations of persistence layer interfaces.
// It includes interactions with PostgreSQL for user account storage, modification and retrieval.
package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"event-calendar/internal/domain"
	"event-calendar/internal/dto/dmodel"
	errs "event-calendar/internal/error"
	"event-calendar/internal/logger"
	mapper "event-calendar/internal/mapper/user/dmodel"
	"event-calendar/internal/repository"

	"github.com/jmoiron/sqlx"
)

// UserAccountRepositoryPostgres implements UserAccountRepository using sqlx.
// It provides methods to manage user accounts in PostgreSQL.
type UserAccountRepositoryPostgres struct {
	// dbDriver abstraction
	dbDriver sqlx.ExtContext
	logger   logger.Logger
}

func NewUserAccountRepository(dbDriver sqlx.ExtContext) *UserAccountRepositoryPostgres {
	return &UserAccountRepositoryPostgres{
		dbDriver: dbDriver,
	}
}

// WithLogger sets the logger and returns the *UserAccountRepositoryPostgres
func (repo *UserAccountRepositoryPostgres) WithLogger(logger logger.Logger) *UserAccountRepositoryPostgres {
	if logger != nil {
		repo.logger = logger
	}
	return repo
}

// WithTx returns new copy of UserAccountRepository with new dbDriver.
func (repo *UserAccountRepositoryPostgres) WithTx(tx *sqlx.Tx) repository.UserAccountRepository {
	return &UserAccountRepositoryPostgres{
		dbDriver: tx,
		logger:   repo.logger,
	}
}

// ListUserAccountsByUserID retrieves a user list by user ID.
// Returns errs.ErrUserAccountNotFound if no matching record exists.
func (repo *UserAccountRepositoryPostgres) ListUserAccountsByUserID(ctx context.Context, userID int64) (userAccounts []domain.UserAccount, err error) {
	var accounts []dmodel.UserAccount
	err = sqlx.SelectContext(ctx, repo.dbDriver, &accounts,
		`SELECT * FROM user_accounts
				WHERE user_id = $1`, userID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errs.ErrUserAccountNotFound.WithInfo(fmt.Sprintf("repository.ListUserAccountsByUserID: user accounts not found for userID=%d", userID))
		}
		return nil, errs.ErrDB.WithInfo(fmt.Sprintf("repository.ListUserAccountsByUserID: %v", err))
	}
	return mapper.MapUserAccounts(accounts), nil
}

// CreateUserAccount inserts a new user account into the database.
// It accepts ignoreConflict option to ignore duplicate conflict.
// If ignoreConflict is true then no error would be returned after try to create duplicate.
// IMPORTANT: ignore given CreatedAt value.
func (repo *UserAccountRepositoryPostgres) CreateUserAccount(ctx context.Context, user domain.UserAccount, ignoreConflict bool) (userAccountID int64, err error) {
	if ignoreConflict {
		return repo.createUserAccountIgnoreConflict(ctx, user)
	}
	return repo.createUserAccount(ctx, user)
}

func (repo *UserAccountRepositoryPostgres) createUserAccountIgnoreConflict(ctx context.Context, user domain.UserAccount) (userAccountID int64, err error) {
	err = repo.dbDriver.QueryRowxContext(
		ctx,
		`INSERT INTO user_accounts (user_id, issuer, subject_uid, email_address, contact_name)
				VALUES ($1, $2, $3, $4, $5) ON CONFLICT DO NOTHING RETURNING id`,
		user.UserID, user.Issuer, user.SubjectUID, user.EmailAddress, user.ContactName,
	).Scan(&userAccountID)
	if err != nil {
		return 0, errs.ErrDB.WithInfo(fmt.Sprintf("repository.createUserAccountIgnoreConflict: %v", err))
	}
	return userAccountID, nil
}

func (repo *UserAccountRepositoryPostgres) createUserAccount(ctx context.Context, user domain.UserAccount) (userAccountID int64, err error) {
	err = repo.dbDriver.QueryRowxContext(
		ctx,
		`INSERT INTO user_accounts (user_id, issuer, subject_uid, email_address, contact_name)
				VALUES ($1, $2, $3, $4, $5) RETURNING id`,
		user.UserID, user.Issuer, user.SubjectUID, user.EmailAddress, user.ContactName,
	).Scan(&userAccountID)
	if err != nil {
		errorInfo := fmt.Sprintf("repository.createUserAccount: %v", err)

		if len(err.Error()) > 50 && err.Error()[:50] == pqDuplicateErr {
			return 0, errs.ErrUserAccountExists.WithInfo(errorInfo)
		}
		return 0, errs.ErrDB.WithInfo(errorInfo)
	}
	return userAccountID, nil
}
