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

type UserRepositoryPostgres struct {
	// dbDriver adapter abstraction
	dbDriver sqlx.ExtContext
}

func NewUserRepository(dbAdapter sqlx.ExtContext) UserRepositoryPostgres {
	return UserRepositoryPostgres{
		dbDriver: dbAdapter,
	}
}

func (repo UserRepositoryPostgres) WithTx(tx *sqlx.Tx) repository.UserRepository {
	return UserRepositoryPostgres{
		dbDriver: tx,
	}
}

func (repo UserRepositoryPostgres) GetUsersCount(ctx context.Context) (int64, error) {
	var count int64
	err := sqlx.GetContext(ctx, repo.dbDriver, &count,
		`SELECT count(*) FROM users`)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return 0, repository.ErrNoRows
		}
		return 0, err
	}

	return count, nil
}

func (repo UserRepositoryPostgres) GetUserByID(ctx context.Context, id int64) (obj domain.User, err error) {
	var userDto dmodel.User
	err = sqlx.GetContext(ctx, repo.dbDriver, &userDto,
		`SELECT * FROM users
				WHERE id = $1`, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return domain.User{}, repository.ErrNoRows
		}
		return domain.User{}, err
	}
	return mapper.UserDtoToUser(userDto), nil
}

func (repo UserRepositoryPostgres) GetUserByFirebaseUID(ctx context.Context, firebaseUID string) (obj domain.User, err error) {
	var userDto dmodel.User
	err = sqlx.GetContext(ctx, repo.dbDriver, &userDto,
		`SELECT * FROM users
				WHERE firebase_uid = $1`, firebaseUID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return domain.User{}, repository.ErrNoRows
		}
		return domain.User{}, err
	}
	return mapper.UserDtoToUser(userDto), nil
}

// CreateUser
// IMPORTANT: ignore given Roles, CreatedAt values.
func (repo UserRepositoryPostgres) CreateUser(ctx context.Context, user domain.User) (userID int64, err error) {
	err = repo.dbDriver.QueryRowxContext(
		ctx,
		`INSERT INTO users (firebase_uid, description) VALUES ($1, $2) RETURNING id`,
		user.FirebaseUID, user.Description,
	).Scan(&userID)
	if err != nil {
		if len(err.Error()) > 50 {
			if err.Error()[:50] == pqDuplicateErr {
				return 0, repository.ErrDuplicate
			}
		}
		return 0, err
	}
	return userID, nil
}
