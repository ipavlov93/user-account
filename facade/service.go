package facade

import (
	"context"

	"event-calendar/internal/domain"
	"event-calendar/internal/domain/claims"
	"event-calendar/internal/domain/role"
	"event-calendar/internal/option"

	"firebase.google.com/go/v4/auth"
)

type AuthService interface {
	SignUp(ctx context.Context, claims *claims.FirebaseAuthClaims, roles []role.Role) error
	Login(ctx context.Context, claims *claims.FirebaseAuthClaims) error
	Logout(ctx context.Context, token string) error
}

type FirebaseAuthService interface {
	VerifyIDToken(idToken string) (token *auth.Token, err error)
	RevokeRefreshTokens(ctx context.Context, idToken string) error
	ResetCustomClaims(firebaseUUID string, customClaims map[string]any) error
}

type UserService interface {
	CreateUser(ctx context.Context, user domain.User, options *option.TxOption) (int64, error)
	GetUserByID(ctx context.Context, ID int64, options *option.TxOption) (user domain.User, found bool, err error)
	GetUserByUUID(ctx context.Context, uuid string, options *option.TxOption) (user domain.User, found bool, err error)
}

type UserProfileService interface {
	CreateUserProfile(ctx context.Context, user domain.UserProfile, options *option.CreateOptions) (int64, error)
	GetUserProfileByID(ctx context.Context, ID int64, options *option.TxOption) (user domain.UserProfile, found bool, err error)
	GetUserProfileByUserID(ctx context.Context, ID int64, options *option.TxOption) (user domain.UserProfile, found bool, err error)
	GetUserProfileByUUID(ctx context.Context, uuid string, options *option.TxOption) (user domain.UserProfile, found bool, err error)
}

type UserAccountService interface {
	CreateUserAccount(ctx context.Context, userAccount domain.UserAccount, options *option.CreateUserAccountOptions) (int64, error)
	ListUserAccountsByUserID(ctx context.Context, userID int64, options *option.TxOption) (userAccountsList []domain.UserAccount, err error)
}
