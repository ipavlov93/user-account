package userservice

import (
	"context"
	"event-calendar/internal/domain"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user domain.User) (int64, error)
	GetUsersCount(ctx context.Context) (int64, error)
	GetUserByID(ctx context.Context, id int64) (domain.User, error)
	GetUserByFirebaseUID(ctx context.Context, uuid string) (domain.User, error)
}

type UserAccountRepository interface {
	CreateUserAccount(ctx context.Context, user domain.UserAccount, ignoreDuplicate bool) (userAccountID int64, err error)
	ListUserAccountsByUserID(ctx context.Context, userID int64) (userAccounts []domain.UserAccount, err error)
}
