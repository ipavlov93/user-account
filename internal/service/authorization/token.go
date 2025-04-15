package authorization

import (
	"context"
	"crypto"
	"encoding/json"
	"fmt"

	"event-calendar/internal/domain/claims"

	"github.com/MicahParks/jwkset"
	"github.com/golang-jwt/jwt/v5"
)

func InitializeJWKSetClient(jwksURLs []string) (jwkset.Storage, error) {
	jwks, err := jwkset.NewDefaultHTTPClient(jwksURLs)
	if err != nil {
		return nil, fmt.Errorf("failed to create JWK Set client: %w", err)
	}
	return jwks, nil
}

// VerifyAccessToken verifies the JWT access token and extracts user claims
func VerifyAccessToken(jwks jwkset.Storage, accessToken string) (userClaims claims.UserClaims, err error) {
	// Parse the JWT token
	token, err := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
		// Extract the key ID (kid) from the token header
		kid, ok := token.Header["kid"].(string)
		if !ok {
			return nil, fmt.Errorf("missing kid in token header")
		}

		// Retrieve the public key from the JWK Set
		ctx := context.Background()
		jwk, err := jwks.KeyRead(ctx, kid)
		if err != nil {
			return nil, fmt.Errorf("failed to read key from JWK Set: %w", err)
		}

		// Obtain the public key from the JWK
		// Ensure key is valid
		key := jwk.Key()
		if _, ok := key.(crypto.PublicKey); !ok {
			return nil, fmt.Errorf("unexpected key type: %T", key)
		}

		// Return the public key for token verification
		return jwk.Key(), nil
	})
	if err != nil {
		return userClaims, fmt.Errorf("failed to parse token: %w", err)
	}

	// Check if the token is valid
	if !token.Valid {
		return userClaims, fmt.Errorf("invalid token")
	}

	// Validate token claims
	claimsMap, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return userClaims, fmt.Errorf("failed to parse token claims")
	}

	userClaims, err = parseUserClaims(claimsMap)
	if err != nil {
		return userClaims, fmt.Errorf("failed to parse user claims: %w", err)
	}

	// TODO: delete commented code after testing the flow
	//// Extract standard and custom claims
	//userClaims = claims.UserClaims{
	//	Subject: claimsMap["sub"].(string),
	//	Issuer:  claimsMap["iss"].(string),
	//}
	//
	//// Required custom claim
	//if firebaseUID, exists := claimsMap["firebase_uid"].(string); exists {
	//	userClaims.FirebaseUID = firebaseUID
	//}
	//
	//// Optional claims
	//if email, exists := claimsMap["email"].(string); exists {
	//	userClaims.Email = email
	//}
	//if roles, exists := claimsMap["roles"].([]claims.Role); exists {
	//	userClaims.Roles = roles
	//}

	return userClaims, nil
}

func parseUserClaims(claimsMap map[string]any) (parsedUserClaims claims.UserClaims, err error) {
	claimsJSON, err := json.Marshal(claimsMap)
	if err != nil {
		return parsedUserClaims, fmt.Errorf("failed to marshal claims: %w", err)
	}

	parsedUserClaims = claims.UserClaims{}
	err = json.Unmarshal(claimsJSON, &parsedUserClaims)
	if err != nil {
		return parsedUserClaims, fmt.Errorf("failed to unmarshal claims: %w", err)
	}

	return parsedUserClaims, nil
}
