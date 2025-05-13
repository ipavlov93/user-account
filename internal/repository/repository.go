// Package repository provides contracts of persistence layer.
// Interfaces includes methods for storage, modification and retrieval.
// It also provides ability to perform these operations within transaction.
package repository

import (
	"context"

	"event-calendar/internal/domain"

	"github.com/jmoiron/sqlx"
)

// WithTx provides ability to run SQL queries within database transaction.
type WithTx[T any] interface {
	// WithTx returns new copy of T with tx.
	WithTx(tx *sqlx.Tx) T
}

// UserRepository defines the contract for user persistence operations.
type UserRepository interface {
	WithTx[UserRepository]
	CreateUser(ctx context.Context, user domain.User) (int64, error)
	GetUsersCount(ctx context.Context) (int64, error)
	GetUserByID(ctx context.Context, ID int64) (domain.User, error)
	GetUserByFirebaseUUID(ctx context.Context, uuid string) (domain.User, error)
}

// UserProfileRepository defines the contract for user profile persistence operations.
type UserProfileRepository interface {
	WithTx[UserProfileRepository]
	CreateUserProfile(ctx context.Context, user domain.UserProfile) (int64, error)
	GetUserProfilesCount(ctx context.Context) (int64, error)
	GetUserProfileByID(ctx context.Context, ID int64) (user domain.UserProfile, err error)
	GetUserProfileByUserID(ctx context.Context, userID int64) (user domain.UserProfile, err error)
	GetUserProfileByFirebaseUUID(ctx context.Context, firebaseUUID string) (user domain.UserProfile, err error)
}

// UserAccountRepository defines the contract for user account persistence operations.
type UserAccountRepository interface {
	WithTx[UserAccountRepository]
	CreateUserAccount(ctx context.Context, user domain.UserAccount, ignoreConflict bool) (userAccountID int64, err error)
	ListUserAccountsByUserID(ctx context.Context, userID int64) (userAccounts []domain.UserAccount, err error)
}
