package postgres

import (
	"context"
	"database/sql"
	"errors"

	"event-calendar/internal/domain"
	"event-calendar/internal/dto/dmodel"
	mapper "event-calendar/internal/mapper/user/dmodel"
	"event-calendar/internal/repository"

	"github.com/jmoiron/sqlx"
)

type UserAccountRepositoryPostgres struct {
	// dbDriver adapter abstraction
	dbDriver sqlx.ExtContext
}

func NewUserAccountRepository(dbDriver sqlx.ExtContext) UserAccountRepositoryPostgres {
	return UserAccountRepositoryPostgres{
		dbDriver: dbDriver,
	}
}

func (repo UserAccountRepositoryPostgres) WithTx(tx *sqlx.Tx) repository.UserAccountRepository {
	return UserAccountRepositoryPostgres{
		dbDriver: tx,
	}
}

func (repo UserAccountRepositoryPostgres) ListUserAccountsByUserID(ctx context.Context, userID int64) (userAccounts []domain.UserAccount, err error) {
	var accounts []dmodel.UserAccount

	err = sqlx.SelectContext(ctx, repo.dbDriver, &accounts,
		`SELECT * FROM user_accounts
				WHERE user_id = $1`, userID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, repository.ErrNoRows
		}
		return nil, err
	}
	return mapper.MapUserAccounts(accounts), nil
}

func (repo UserAccountRepositoryPostgres) CreateUserAccount(ctx context.Context, user domain.UserAccount, ignoreDuplicate bool) (userAccountID int64, err error) {
	if ignoreDuplicate {
		return repo.createUserAccountIgnoreDuplicate(ctx, user)
	}
	return repo.createUserAccount(ctx, user)
}

func (repo UserAccountRepositoryPostgres) createUserAccountIgnoreDuplicate(ctx context.Context, user domain.UserAccount) (userAccountID int64, err error) {
	err = repo.dbDriver.QueryRowxContext(
		ctx,
		`INSERT INTO user_accounts (user_id, issuer, subject_uid, email_address, contact_name)
				VALUES ($1, $2, $3, $4, $5) ON CONFLICT DO NOTHING RETURNING id`,
		user.UserID, user.Issuer, user.SubjectUID, user.EmailAddress, user.ContactName,
	).Scan(&userAccountID)
	if err != nil {
		if len(err.Error()) > 50 {
			if err.Error()[:50] == pqDuplicateErr {
				return 0, repository.ErrDuplicate
			}
		}
		return 0, err
	}
	return userAccountID, nil
}

func (repo UserAccountRepositoryPostgres) createUserAccount(ctx context.Context, user domain.UserAccount) (userAccountID int64, err error) {
	err = repo.dbDriver.QueryRowxContext(
		ctx,
		`INSERT INTO user_accounts (user_id, issuer, subject_uid, email_address, contact_name)
				VALUES ($1, $2, $3, $4, $5) RETURNING id`,
		user.UserID, user.Issuer, user.SubjectUID, user.EmailAddress, user.ContactName,
	).Scan(&userAccountID)
	if err != nil {
		if len(err.Error()) > 50 {
			if err.Error()[:50] == pqDuplicateErr {
				return 0, repository.ErrDuplicate
			}
		}
		return 0, err
	}
	return userAccountID, nil
}
