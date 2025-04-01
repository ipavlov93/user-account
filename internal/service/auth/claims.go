package auth

import "database/sql"

type FirebaseAuthClaims struct {
	AuthTime      float64        `json:"auth_time"`
	Email         string         `json:"email"`
	EmailVerified bool           `json:"email_verified"`
	Firebase      FirebaseClaims `json:"firebase"`
	Name          string         `json:"name"`
	PictureURL    string         `json:"picture"`
	UserID        string         `json:"user_id"`
}

type FirebaseClaims struct {
	Identities     Identities
	SignInProvider string `json:"sign_in_provider"`
}

type Identities struct {
	Email  IdentitiesStorage `json:"email"`
	Google IdentitiesStorage `json:"google.com"`
	//Microsoft IdentitiesStorage `json:"microsoft.com"`
}

type IdentitiesStorage []string

func (s IdentitiesStorage) IsEmpty() bool {
	return len(s) == 0
}

func (s IdentitiesStorage) GetIdentity() (id sql.NullString) {
	if s.IsEmpty() {
		return id
	}
	return sql.NullString{
		String: s[0],
		Valid:  true,
	}
}
