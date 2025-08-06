package firebase

import (
	"context"
	"time"
	errs "user-account/internal/error"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"google.golang.org/api/option"
)

type AuthService struct {
	authClient *auth.Client
}

func NewAuthService(firebaseCredentialFilePath string) AuthService {
	option := option.WithCredentialsFile(firebaseCredentialFilePath)

	ctx := context.Background()
	firebaseApp, err := firebase.NewApp(ctx, nil, option)
	if err != nil {
		panic(err)
	}

	authClient, err := firebaseApp.Auth(ctx)
	if err != nil {
		panic(err)
	}

	if authClient == nil {
		panic("firebase auth service init failed")
	}

	return AuthService{
		authClient: authClient,
	}
}

// VerifyIDToken uses Firebase SDK (to implement OIDC flow) to verify ID token.
// Returns an err if token is invalid, expired, disabled or revoked.
// TODO: Add key to cache
func (s AuthService) VerifyIDToken(idToken string) (token *auth.Token, err error) {
	if idToken == "" {
		return token, errs.ErrInvalidToken
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	token, err = s.authClient.VerifyIDTokenAndCheckRevoked(ctx, idToken)
	if err != nil {
		switch {
		case auth.IsIDTokenExpired(err):
			return token, errs.ErrExpiredToken
		case auth.IsUserDisabled(err):
			return token, errs.ErrDisabledToken
		case auth.IsIDTokenRevoked(err):
			return token, errs.ErrRevokedToken
		}
		return token, err
	}

	// TODO:
	// 1. Add key to (session) cache if ResetCustomClaims doesn't work

	return token, nil
}

// TODO: remove token from cache
func (s AuthService) RevokeRefreshTokens(ctx context.Context, idToken string) error {
	token, err := s.authClient.VerifyIDToken(ctx, idToken)
	if err != nil {
		return err
	}

	err = s.authClient.RevokeRefreshTokens(context.Background(), token.UID)
	if err != nil {
		return err
	}

	// TODO: remove auth tokens from cache

	return nil
}

// ResetCustomClaims resets custom user claims to future ID tokens (read SetCustomUserClaims docs).
// The new custom claims will propagate to the user's ID token the next time a new one is issued.
// Note: this operation always overwrites the user's existing custom claims.
func (s AuthService) ResetCustomClaims(firebaseUUID string, customClaims map[string]any) error {
	return s.authClient.SetCustomUserClaims(context.Background(), firebaseUUID, customClaims)
}

// todo: checkExistingCustomClaims in token claims before override them (all the custom claims).
//func (s AuthService) checkExistingCustomClaims(firebaseUUID string, roles []claims.Role) error {
