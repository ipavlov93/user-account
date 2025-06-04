package handler

import (
	"context"
	"encoding/json"
	"event-calendar/internal/domain/claims"
	"event-calendar/internal/logger"
	auth "event-calendar/internal/service/authorization"
	firebaseauth "event-calendar/internal/service/firebase"
	"fmt"
	"net/http"
	"strings"

	"go.uber.org/zap"
)

const (
	bearerPrefix        = "Bearer "
	authorizationHeader = "Authorization"

	firebaseClaimsKey = "firebase-claims"
	userClaimsKey     = "user-claims"
)

type AuthMiddleware struct {
	firebaseAuthService firebaseauth.AuthService
	logger              logger.Logger
	providerKeySetURLs  []string // auth provider key set URLs
}

// NewAuthMiddleware set default logger. Use WithLogger() to set custom logger.
func NewAuthMiddleware(
	service firebaseauth.AuthService,
	providerKeySetURLs []string,
) *AuthMiddleware {
	return &AuthMiddleware{
		firebaseAuthService: service,
		providerKeySetURLs:  providerKeySetURLs,
	}
}

// WithLogger sets the logger and returns the *AuthMiddleware
func (m *AuthMiddleware) WithLogger(logger logger.Logger) *AuthMiddleware {
	if logger != nil {
		m.logger = logger
	}
	return m
}

func (m *AuthMiddleware) RequireValidIDToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		token, err := retrieveBearerToken(r)
		if err != nil {
			http.Error(rw,
				fmt.Sprintf("%s: %v", http.StatusText(http.StatusBadRequest), err),
				http.StatusBadRequest)
			return
		}

		idToken, err := m.firebaseAuthService.VerifyIDToken(token)
		if err != nil {
			http.Error(rw,
				http.StatusText(http.StatusUnauthorized),
				http.StatusUnauthorized)
			return
		}

		parsedClaims, err := parseIDTokenClaims(idToken.Claims)
		if err != nil {
			m.logger.Error("parseIDTokenClaims(): parse ID token claims error", zap.Error(err))
			http.Error(rw,
				http.StatusText(http.StatusInternalServerError),
				http.StatusInternalServerError)
			return
		}

		// set claims to context
		ctx := context.WithValue(r.Context(), firebaseClaimsKey, parsedClaims)

		// proceed with the request handling
		next.ServeHTTP(rw, r.WithContext(ctx))
	})
}

func (m *AuthMiddleware) RequireValidAccessToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		accessToken, err := retrieveBearerToken(r)
		if err != nil {
			http.Error(rw,
				fmt.Sprintf("%s: %v", http.StatusText(http.StatusUnauthorized), err),
				http.StatusUnauthorized)
			return
		}

		// Initialize JWK Set client with your JWK Set URLs
		jwks, err := auth.InitializeJWKSetClient(m.providerKeySetURLs)
		if err != nil {
			m.logger.Error("InitializeJWKSetClient()", zap.Error(err))
			http.Error(rw,
				http.StatusText(http.StatusInternalServerError),
				http.StatusInternalServerError)
			return
		}

		// Verify the access token
		userClaims, err := auth.VerifyAccessToken(jwks, accessToken)
		if err != nil {
			http.Error(rw,
				http.StatusText(http.StatusUnauthorized),
				http.StatusUnauthorized)
			return
		}

		r.WithContext(context.WithValue(r.Context(), userClaimsKey, userClaims))

		// Proceed with the request handling
		next.ServeHTTP(rw, r)
	})
}

func retrieveBearerToken(r *http.Request) (string, error) {
	authHeader := r.Header.Get(authorizationHeader)
	if !strings.HasPrefix(authHeader, bearerPrefix) {
		return "", fmt.Errorf("missing or invalid authorization header")
	}
	return strings.TrimPrefix(authHeader, bearerPrefix), nil
}

func parseIDTokenClaims(claimsMap map[string]any) (parsedClaims *claims.FirebaseAuthClaims, err error) {
	claimsJSON, err := json.Marshal(claimsMap)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal claims: %w", err)
	}

	parsedClaims = &claims.FirebaseAuthClaims{}
	err = json.Unmarshal(claimsJSON, parsedClaims)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal claims: %w", err)
	}
	return parsedClaims, nil
}
