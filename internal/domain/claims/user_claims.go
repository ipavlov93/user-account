package claims

import (
	"user-account/internal/domain/role"
)

// UserClaims represents extracted claims from ID token that contains custom claims.
type UserClaims struct {
	FirebaseUUID string `json:"user_id"`
	Subject      string `json:"sub"` // Auth provider user ID
	Issuer       string `json:"iss"` // Auth provider (Google, Azure)
	Email        string `json:"email"`
	ContactName  string `json:"name"`

	// Custom claims apply to users already signed in with supported providers (Email/Password, Google, etc).
	// For example, (already signed) user can have access control defined using custom claims.
	Roles []role.Role `json:"roles"`
}
