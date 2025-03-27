package userservice

import (
	"context"

	mapper "event-calendar/mapper/smodel"
	"event-calendar/smodel"
)

type UserService struct {
	repository UserRepository
}

func NewUserService(repository UserRepository) *UserService {
	return &UserService{
		repository,
	}
}

func (s UserService) GetUserByID(ctx context.Context, id int64) (userDto smodel.User, err error) {
	user, err := s.repository.GetUserByID(ctx, id)
	if err != nil {
		//if errors.Is(err, repository.ErrNoRows) {
		//return customError with status code NotFound
		//}
		return userDto, err
	}

	return mapper.UserToUserDto(user), nil
}

func (s UserService) GetUserByUUID(ctx context.Context, uuid string) (userDto smodel.User, err error) {
	user, err := s.repository.GetUserByUUID(ctx, uuid)
	if err != nil {
		//if errors.Is(err, repository.ErrNoRows) {
		//return customError with status code NotFound
		//}
		return userDto, err
	}

	return mapper.UserToUserDto(user), nil
}

func (s UserService) CreateUser(ctx context.Context, user smodel.User) (int64, error) {
	ID, err := s.repository.CreateUser(ctx, mapper.UserDtoToUser(user))
	if err != nil {
		//if errors.Is(err, repository.ErrDuplicate) {
		//return customError with status code BadRequest
		//}
		return 0, err
	}

	return ID, nil
}
