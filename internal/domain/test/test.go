package test

import (
	"fmt"
	"time"
	"user-account/internal/domain"
	"user-account/internal/domain/claims"
)

// file test.go contains helper function that is used in tests.

// CreateTestClaims fill all fields based on given value
func CreateTestClaims(value string) claims.FirebaseAuthClaims {
	return claims.FirebaseAuthClaims{
		AuthTime:      0,
		Email:         fmt.Sprintf("%s@example.com", value),
		EmailVerified: false,
		FirebaseIdentityClaims: claims.IdentityClaims{
			SignInProvider: value,
			//Identities: nil,
		},
		Name:       value,
		PictureURL: value,
		UserID:     value,
	}
}

// CreateTestUser fill all fields based on given value.
// IMPORTANT: set CreatedAt as time.Now().
func CreateTestUser(ID int) domain.User {
	return domain.User{
		ID:           int64(ID),
		FirebaseUUID: fmt.Sprint(ID),
		Description:  fmt.Sprintf("Description%d", ID),
		CreatedAt:    time.Now(),
	}
}

// CreateTestUserAccount fill all fields based on given value.
// IMPORTANT: ignore UserID.
func CreateTestUserAccount(ID int) domain.UserAccount {
	return domain.UserAccount{
		ID:           int64(ID),
		Issuer:       domain.NewIssuer(fmt.Sprint(ID)),
		SubjectUID:   fmt.Sprintf("SubjectUID%d", ID),
		EmailAddress: fmt.Sprintf("%d@test.com", ID),
	}
}

// CreateTestUserProfile fill all fields based on given value.
// IMPORTANT:
// - ignore UserID
// - set CreatedAt as time.Now().
func CreateTestUserProfile(ID int) domain.UserProfile {
	return domain.UserProfile{
		ID:             int64(ID),
		FirstName:      fmt.Sprintf("FirstName%d", ID),
		LastName:       fmt.Sprintf("LastName%d", ID),
		BusinessName:   fmt.Sprintf("BusinessName%d", ID),
		ContactEmail:   fmt.Sprintf("%d@test.com", ID),
		Organization:   fmt.Sprintf("Organization%d", ID),
		Description:    fmt.Sprintf("Description%d", ID),
		AvatarFileName: fmt.Sprintf("AvatarFileName%d", ID),
		CreatedAt:      time.Now(),
	}
}
