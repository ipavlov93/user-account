package claims

// UserClaims represents extracted claims from an access token
type UserClaims struct {
	FirebaseUID string `json:"user_id"`
	Subject     string `json:"sub"` // Auth provider user ID
	Issuer      string `json:"iss"` // Auth provider (Google, Azure)
	Email       string `json:"email"`
	ContactName string `json:"name"`

	Roles []Role `json:"roles"`
}
