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

type UserProfileRepository interface {
	CreateUserProfile(ctx context.Context, user domain.UserProfile) (int64, error)
	GetUserProfilesCount(ctx context.Context) (int64, error)
	GetUserProfileByID(ctx context.Context, id int64) (user domain.UserProfile, err error)
	GetUserProfileByUserID(ctx context.Context, userID int64) (user domain.UserProfile, err error)
	GetUserProfileByFirebaseUID(ctx context.Context, firebaseUID string) (user domain.UserProfile, err error)
}

type UserAccountRepository interface {
	CreateUserAccount(ctx context.Context, user domain.UserAccount, ignoreDuplicate bool) (userAccountID int64, err error)
	ListUserAccountsByUserID(ctx context.Context, userID int64) (userAccounts []domain.UserAccount, err error)
}
