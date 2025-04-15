package claims

// UserClaims represents extracted claims from an access token
type UserClaims struct {
	Subject string `json:"sub"` // User ID
	Issuer  string `json:"iss"` // Auth provider (Google, Azure)
	Email   string `json:"email"`

	FirebaseUID string `json:"firebase_uid"`
	Roles       []Role `json:"roles"`
}
