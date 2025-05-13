package userservice

import (
	"context"
	"event-calendar/internal/domain"
	"event-calendar/internal/option"
	"event-calendar/internal/repository"
	"fmt"
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
		return domain.User{}, false, fmt.Errorf("service.GetUserByID: %w", err)
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

	user, err := repo.GetUserByFirebaseUUID(ctx, uuid)
	if err != nil {
		return domain.User{}, false, fmt.Errorf("service.GetUserByUUID: %w", err)
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

		return 0, fmt.Errorf("service.CreateUser: %w", err)
	}

	return userID, nil
}
