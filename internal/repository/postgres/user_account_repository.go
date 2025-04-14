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

type UserAccountRepository struct {
	dbDriver sqlx.ExtContext
}

func NewUserAccountRepository(dbDriver sqlx.ExtContext) UserAccountRepository {
	return UserAccountRepository{
		dbDriver: dbDriver,
	}
}

//func (repo UserAccountRepository) GetUserAccountByID(ctx context.Context, id int64) (user domain.UserAccount, err error) {
//	var userAccountDto dmodel.UserAccount
//	err = sqlx.GetContext(ctx, repo.dbDriver, &userAccountDto,
//		`SELECT * FROM user_accounts
//				WHERE id = $1`, id)
//	if err != nil {
//		if errors.Is(err, sql.ErrNoRows) {
//			return user, repository.ErrNoRows
//		}
//		return user, err
//	}
//	return mapper.UserAccountDtoToUserAccount(userAccountDto), nil
//}

func (repo UserAccountRepository) ListUserAccountsByUserID(ctx context.Context, userID int64) (userAccounts []domain.UserAccount, err error) {
	var userAccountDtos []dmodel.UserAccount

	err = sqlx.SelectContext(ctx, repo.dbDriver, &userAccountDtos,
		`SELECT * FROM user_accounts
				WHERE user_id = $1`, userID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return userAccounts, repository.ErrNoRows
		}
		return userAccounts, err
	}
	return mapper.DtosToUserAccounts(userAccountDtos), nil
}

func (repo UserAccountRepository) CreateUserAccount(ctx context.Context, user domain.UserAccount, ignoreDuplicate bool) (userAccountID int64, err error) {
	if ignoreDuplicate {
		return repo.createUserAccountIgnoreDuplicate(ctx, user)
	}
	return repo.createUserAccount(ctx, user)
}

func (repo UserAccountRepository) createUserAccountIgnoreDuplicate(ctx context.Context, user domain.UserAccount) (userAccountID int64, err error) {
	err = repo.dbDriver.QueryRowxContext(
		ctx,
		`INSERT INTO user_accounts (user_id, issuer_code, subject_uid, email_address)
				VALUES ($1, $2, $3, $4) ON CONFLICT DO NOTHING RETURNING id`,
		user.UserID, user.IssuerCode, user.SubjectUID, user.EmailAddress,
	).Scan(&userAccountID)
	if err != nil {
		if len(err.Error()) > 50 {
			if err.Error()[:50] == "pq: duplicate key value violates unique constraint" {
				return 0, repository.ErrDuplicate
			}
		}
		return 0, err
	}
	return userAccountID, nil
}

func (repo UserAccountRepository) createUserAccount(ctx context.Context, user domain.UserAccount) (userAccountID int64, err error) {
	err = repo.dbDriver.QueryRowxContext(
		ctx,
		`INSERT INTO user_accounts (user_id, issuer_code, subject_uid, email_address)
				VALUES ($1, $2, $3, $4) RETURNING id`,
		user.UserID, user.IssuerCode, user.SubjectUID, user.EmailAddress,
	).Scan(&userAccountID)
	if err != nil {
		if len(err.Error()) > 50 {
			if err.Error()[:50] == "pq: duplicate key value violates unique constraint" {
				return 0, repository.ErrDuplicate
			}
		}
		return 0, err
	}
	return userAccountID, nil
}
