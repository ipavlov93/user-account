package authentication

import (
	"context"
	"time"

	"event-calendar/internal/domain/role"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"google.golang.org/api/option"
)

type FirebaseAuthService struct {
	firebaseAuthService *auth.Client
}

func NewAuthService(firebaseCredentialFilePath string) FirebaseAuthService {
	option := option.WithCredentialsFile(firebaseCredentialFilePath)

	ctx := context.Background()
	firebaseApp, err := firebase.NewApp(ctx, nil, option)
	if err != nil {
		panic(err)
	}

	firebaseAuthService, err := firebaseApp.Auth(ctx)
	if err != nil {
		panic(err)
	}

	if firebaseAuthService == nil {
		panic("firebase auth service init failed")
	}

	return FirebaseAuthService{
		firebaseAuthService: firebaseAuthService,
	}
}

// VerifyIDToken uses Firebase SDK (to implement OIDC flow) to verify ID token.
// Returns an err if token is invalid, expired, disabled or revoked.
// TODO: Add key to cache
func (s FirebaseAuthService) VerifyIDToken(idToken string) (token *auth.Token, err error) {
	//if idToken == "" {
	//	return token, ErrInvalidToken
	//}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	token, err = s.firebaseAuthService.VerifyIDTokenAndCheckRevoked(ctx, idToken)
	if err != nil {
		return nil, err
	}
	//if err != nil {
	//	switch {
	//	case auth.IsIDTokenExpired(err):
	//		return token, ErrExpiredToken
	//	case auth.IsUserDisabled(err):
	//		return token, ErrDisabledToken
	//	case auth.IsIDTokenRevoked(err):
	//		return token, ErrRevokedToken
	//	}
	//
	//	return token, err
	//}

	// TODO:
	// 1. Add key to (session) cache if SetRolePrivilegesToClaims doesn't work
	// 1. test auth after SetRolePrivilegesToClaims, does token (access or id?) contains firebaseUID

	return token, nil
}

// TODO: remove token from cache
func (s FirebaseAuthService) RevokeRefreshTokens(ctx context.Context, idToken string) error {
	token, err := s.firebaseAuthService.VerifyIDToken(ctx, idToken)
	if err != nil {
		return err
	}

	err = s.firebaseAuthService.RevokeRefreshTokens(context.Background(), token.UID)
	if err != nil {
		return err
	}

	// TODO: remove auth tokens from cache

	return nil
}

// SetRolePrivilegesToClaims sets custom claims to future ID tokens.
// It sets "processed" custom claim to mark user in Firestore as synced with back-end.
// The new custom claims will propagate to the user's ID token the next time a new one is issued.
// Note: this operation always overwrites the user's existing custom claims.
func (s FirebaseAuthService) SetRolePrivilegesToClaims(firebaseUID string, roles []role.Role) error {
	customClaims := map[string]any{
		"roles": roles,
	}

	return s.firebaseAuthService.SetCustomUserClaims(context.Background(), firebaseUID, customClaims)
}

// todo: checkExistingCustomClaims in token claims before override them (all the custom claims).
//func (s FirebaseAuthService) checkExistingCustomClaims(firebaseUID string, roles []claims.Role) error {
