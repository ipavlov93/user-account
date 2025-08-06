// Package postgres is part of repository.
// Package provides implementations of persistence layer interfaces.
// It includes interactions with PostgreSQL for user storage, modification and retrieval.
package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"user-account/internal/domain"
	"user-account/internal/dto/dmodel"
	errs "user-account/internal/error"
	"user-account/internal/logger"
	mapper "user-account/internal/mapper/user/dmodel"
	"user-account/internal/repository"

	"github.com/jmoiron/sqlx"
)

// UserRepositoryPostgres implements UserRepository using sqlx.
// It provides methods to manage users in PostgreSQL.
type UserRepositoryPostgres struct {
	// dbDriver abstraction
	dbDriver sqlx.ExtContext
	logger   logger.Logger
}

func NewUserRepository(dbDriver sqlx.ExtContext) *UserRepositoryPostgres {
	return &UserRepositoryPostgres{
		dbDriver: dbDriver,
	}
}

// WithLogger sets the logger and returns the *UserRepositoryPostgres
func (repo *UserRepositoryPostgres) WithLogger(logger logger.Logger) *UserRepositoryPostgres {
	if logger != nil {
		repo.logger = logger
	}
	return repo
}

// WithTx returns new copy of UserAccountRepository with new dbDriver.
func (repo *UserRepositoryPostgres) WithTx(tx *sqlx.Tx) repository.UserRepository {
	return &UserRepositoryPostgres{
		dbDriver: tx,
		logger:   repo.logger,
	}
}

// GetUsersCount retrieves a user total count.
func (repo *UserRepositoryPostgres) GetUsersCount(ctx context.Context) (int64, error) {
	var count int64
	err := sqlx.GetContext(ctx, repo.dbDriver, &count,
		`SELECT count(*) FROM users`)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return 0, nil
		}
		return 0, errs.ErrDB.WithInfo(fmt.Sprintf("repository.GetUserByID: %v", err))
	}

	return count, nil
}

// GetUserByID retrieves a user by its unique ID.
// Returns errs.ErrUserNotFound if no matching record exists.
func (repo *UserRepositoryPostgres) GetUserByID(ctx context.Context, ID int64) (obj domain.User, err error) {
	var userDto dmodel.User
	err = sqlx.GetContext(ctx, repo.dbDriver, &userDto,
		`SELECT * FROM users
				WHERE id = $1`, ID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return domain.User{}, errs.ErrUserNotFound.WithInfo(
				fmt.Sprintf("repository.GetUserByID: user not found for ID=%d", ID))
		}
		return domain.User{}, errs.ErrDB.WithInfo(fmt.Sprintf("repository.GetUserByID: %v", err))
	}
	return mapper.UserDtoToUser(userDto), nil
}

// GetUserByFirebaseUUID retrieves a user by its unique Firebase UID.
// Returns errs.ErrUserNotFound if no matching record exists.
func (repo *UserRepositoryPostgres) GetUserByFirebaseUUID(ctx context.Context, firebaseUUID string) (obj domain.User, err error) {
	var userDto dmodel.User
	err = sqlx.GetContext(ctx, repo.dbDriver, &userDto,
		`SELECT * FROM users
				WHERE firebase_uuid = $1`, firebaseUUID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return domain.User{}, errs.ErrUserNotFound.WithInfo(
				fmt.Sprintf("repository.GetUserByFirebaseUUID: user not found for UUID=%s", firebaseUUID))
		}
		return domain.User{}, errs.ErrDB.WithInfo(fmt.Sprintf("repository.GetUserByFirebaseUUID: %v", err))
	}
	return mapper.UserDtoToUser(userDto), nil
}

// CreateUser inserts a new user into the database.
// Returns an error if the operation fails or the user ID already exists.
// IMPORTANT: ignore given Roles, CreatedAt values.
func (repo *UserRepositoryPostgres) CreateUser(ctx context.Context, user domain.User) (userID int64, err error) {
	err = repo.dbDriver.QueryRowxContext(
		ctx,
		`INSERT INTO users (firebase_uuid, description) VALUES ($1, $2) RETURNING id`,
		user.FirebaseUUID, user.Description,
	).Scan(&userID)
	if err != nil {
		errorInfo := fmt.Sprintf("repository.CreateUser: %v", err)

		if len(err.Error()) > 50 && err.Error()[:50] == pqDuplicateErr {
			return 0, errs.ErrUserExists.WithInfo(errorInfo)
		}
		return 0, errs.ErrDB.WithInfo(errorInfo)
	}
	return userID, nil
}
