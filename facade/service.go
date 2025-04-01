package facade

import (
	"context"

	"event-calendar/internal/domain"
)

type UserService interface {
	GetUserByID(ctx context.Context, id int64) (user domain.User, found bool, err error)
	GetUserByUUID(ctx context.Context, uuid string) (user domain.User, found bool, err error)
	CreateUser(ctx context.Context, user domain.User) (int64, error)
}

type AuthService interface {
	SignUp(ctx context.Context, token string) error
	Login(ctx context.Context, token string) error
	Logout(ctx context.Context, token string) error
}
