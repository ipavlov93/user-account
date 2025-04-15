package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"event-calendar/internal/domain/claims"
	firebaseauth "event-calendar/internal/service/authentication"
	auth "event-calendar/internal/service/authorization"
)

const (
	bearerPrefix         = "Bearer "
	authorizationHeader  = "Authorization"
	authenticationHeader = "Authentication"

	tokenClaimsKey = "claims"
)

type AuthMiddleware struct {
	firebaseAuthService firebaseauth.FirebaseAuthService
}

func NewAuthMiddleware(service firebaseauth.FirebaseAuthService) AuthMiddleware {
	return AuthMiddleware{
		firebaseAuthService: service,
	}
}

func (m AuthMiddleware) RequireValidIDToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		token, err := retrieveBearerToken(r)
		if err != nil {
			http.Error(rw,
				fmt.Sprintf("%s: %v", http.StatusText(http.StatusBadRequest), err),
				http.StatusBadRequest)
		}

		idToken, err := m.firebaseAuthService.VerifyIDToken(token)
		if err != nil {
			http.Error(rw,
				http.StatusText(http.StatusUnauthorized),
				http.StatusUnauthorized)
		}

		parsedClaims, err := parseIDTokenClaims(idToken.Claims)
		if err != nil {
			//log.Errorf("Parse ID token claims error: %s", err)
			http.Error(rw,
				http.StatusText(http.StatusInternalServerError),
				http.StatusInternalServerError)
		}

		ctx := context.WithValue(r.Context(), tokenClaimsKey, parsedClaims)

		// Proceed with the request handling
		next.ServeHTTP(rw, r.WithContext(ctx))
	})
}

func RequireValidAccessToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		accessToken, err := retrieveBearerToken(r)
		if err != nil {
			http.Error(rw,
				fmt.Sprintf("%s: %v", http.StatusText(http.StatusUnauthorized), err),
				http.StatusUnauthorized)
		}

		// Initialize JWK Set client with your JWK Set URLs
		jwksURLs := []string{
			"https://www.googleapis.com/oauth2/v3/certs",                   // Google
			"https://login.microsoftonline.com/common/discovery/v2.0/keys", // Azure
		}
		jwks, err := auth.InitializeJWKSetClient(jwksURLs)
		if err != nil {
			http.Error(rw,
				fmt.Sprintf("failed to initialize JWK Set client: %v", err),
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

		r.WithContext(context.WithValue(r.Context(), "userClaims", userClaims))

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

func parseIDTokenClaims(claimsMap map[string]any) (parsed *claims.FirebaseAuthClaims, err error) {
	bt, err := json.Marshal(claimsMap)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(bt, &parsed)
	return parsed, err
}
