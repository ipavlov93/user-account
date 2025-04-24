package userservice

import (
	"context"

	"event-calendar/internal/domain"
	"event-calendar/internal/option"
	"event-calendar/internal/repository"
)

type UserService struct {
	userRepository repository.UserRepository
}

func NewUserService(
	userRepository repository.UserRepository,
) *UserService {
	return &UserService{
		userRepository: userRepository,
	}
}

func (s UserService) GetUserByID(
	ctx context.Context,
	ID int64,
	options *option.TxOption,
) (
	obj domain.User,
	found bool,
	err error,
) {
	// inject tx into repository
	repo := option.ApplyTx(s.userRepository, options)

	user, err := repo.GetUserByID(ctx, ID)
	if err != nil {
		//if errors.Is(err, repository.ErrNoRows) {
		//return customError with status code NotFound
		//}
		return domain.User{}, false, err
	}

	return user, user.HasValidID(), nil
}

func (s UserService) GetUserByUUID(
	ctx context.Context,
	uuid string,
	options *option.TxOption,
) (
	obj domain.User,
	found bool,
	err error,
) {
	// inject tx into repository
	repo := option.ApplyTx(s.userRepository, options)

	user, err := repo.GetUserByFirebaseUID(ctx, uuid)
	if err != nil {
		//if errors.Is(err, repository.ErrNoRows) {
		//return customError with status code NotFound
		//}
		return domain.User{}, false, err
	}

	return user, user.HasValidID(), nil
}

// CreateUser supplies options as struct instance instead of functional-style WithOption() calls.
// Pass options as the nil if you want to apply default behaviour.
func (s UserService) CreateUser(
	ctx context.Context,
	user domain.User,
	options *option.TxOption,
) (int64, error) {
	// inject tx into repository
	repo := option.ApplyTx(s.userRepository, options)

	userID, err := repo.CreateUser(ctx, user)
	if err != nil {
		//if errors.Is(err, repository.ErrDuplicate) {
		//return customError with status code BadRequest
		//}
		return 0, err
	}

	return userID, nil
}
