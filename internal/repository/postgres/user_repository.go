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

type UserRepository struct {
	// db adapter abstraction
	db sqlx.ExtContext
}

func NewUserRepository(dbAdapter sqlx.ExtContext) UserRepository {
	return UserRepository{
		db: dbAdapter,
	}
}

func (repo UserRepository) GetUsersCount(ctx context.Context) (int64, error) {
	var count int64
	err := sqlx.GetContext(ctx, repo.db, &count,
		`SELECT count(*) FROM users`)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return 0, repository.ErrNoRows
		}
		return 0, err
	}

	return count, nil
}

func (repo UserRepository) GetUserByID(ctx context.Context, id int64) (user domain.User, err error) {
	var userDto dmodel.User
	err = sqlx.GetContext(ctx, repo.db, &userDto,
		`SELECT * FROM users
				WHERE id = $1`, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return user, repository.ErrNoRows
		}
		return user, err
	}
	return mapper.UserDtoToUser(userDto), nil
}

func (repo UserRepository) GetUserByFirebaseUID(ctx context.Context, firebaseUID string) (user domain.User, err error) {
	var userDto dmodel.User
	err = sqlx.GetContext(ctx, repo.db, &userDto,
		`SELECT * FROM users
				WHERE firebase_uid = $1`, firebaseUID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return user, repository.ErrNoRows
		}
		return user, err
	}
	return mapper.UserDtoToUser(userDto), nil
}

func (repo UserRepository) CreateUser(ctx context.Context, user domain.User) (userID int64, err error) {
	err = repo.db.QueryRowxContext(
		ctx,
		`INSERT INTO users (firebase_uid, business_name, first_name, last_name, email_address, organization, description)
				VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id`,
		user.FirebaseUID, user.BusinessName, user.FirstName, user.LastName, user.EmailAddress, user.Organization, user.Description,
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
