package facade

import (
	"context"
	"database/sql"

	"event-calendar/internal/domain"
	"event-calendar/internal/domain/claims"

	"firebase.google.com/go/v4/auth"
)

type AuthService interface {
	SignUp(ctx context.Context, claims *claims.FirebaseAuthClaims) error
	Login(ctx context.Context, claims *claims.FirebaseAuthClaims) error
	Logout(ctx context.Context, token string) error
}

type FirebaseAuthService interface {
	VerifyIDToken(idToken string) (token *auth.Token, err error)
	RevokeRefreshTokens(ctx context.Context, idToken string) error
	SetFirebaseUID(firebaseUID string) error
}

type UserService interface {
	CreateUser(ctx context.Context, user domain.User, tx ...*sql.Tx) (int64, error)
	GetUserByID(ctx context.Context, id int64) (user domain.User, found bool, err error)
	GetUserByUUID(ctx context.Context, uuid string) (user domain.User, found bool, err error)
}

type UserAccountService interface {
	CreateUserAccount(ctx context.Context, userAccount domain.UserAccount, ignoreDuplicate bool, tx ...*sql.Tx) (int64, error)
	ListUserAccountsByUserID(ctx context.Context, userID int64) (userAccountsList []domain.UserAccount, err error)
}
