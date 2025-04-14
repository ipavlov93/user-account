package userservice

import (
	"context"

	"event-calendar/internal/domain"
	repository "event-calendar/internal/repository/postgres"

	"github.com/jmoiron/sqlx"
)

type UserService struct {
	dbDriver sqlx.ExtContext
}

func NewUserService(dbDriver sqlx.ExtContext) *UserService {
	return &UserService{
		dbDriver: dbDriver,
	}
}

func (s UserService) GetUserByID(ctx context.Context, id int64) (user domain.User, found bool, err error) {
	repo := repository.NewUserRepository(s.dbDriver)

	user, err = repo.GetUserByID(ctx, id)
	if err != nil {
		//if errors.Is(err, repository.ErrNoRows) {
		//return customError with status code NotFound
		//}
		return user, false, err
	}

	return user, user.HasValidID(), nil
}

func (s UserService) GetUserByUUID(ctx context.Context, uuid string) (user domain.User, found bool, err error) {
	repo := repository.NewUserRepository(s.dbDriver)

	user, err = repo.GetUserByFirebaseUID(ctx, uuid)
	if err != nil {
		//if errors.Is(err, repository.ErrNoRows) {
		//return customError with status code NotFound
		//}
		return user, false, err
	}

	return user, user.HasValidID(), nil
}

func (s UserService) CreateUser(ctx context.Context, user domain.User) (int64, error) {
	repo := repository.NewUserRepository(s.dbDriver)

	userID, err := repo.CreateUser(ctx, user)
	if err != nil {
		//if errors.Is(err, repository.ErrDuplicate) {
		//return customError with status code BadRequest
		//}
		return 0, err
	}

	return userID, nil
}
