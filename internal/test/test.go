package test

import (
	"event-calendar/internal/domain"
	"event-calendar/internal/domain/claims"
	"fmt"
)

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
		SubjectUID: value,
	}
}

// CreateTestUser fill all fields based on given value
func CreateTestUser(id int) domain.User {
	return domain.User{
		ID:           int64(id),
		FirebaseUID:  fmt.Sprint(id),
		BusinessName: fmt.Sprintf("BusinessName%d", id),
		FirstName:    fmt.Sprintf("FirstName%d", id),
		LastName:     fmt.Sprintf("LastName%d", id),
		EmailAddress: fmt.Sprintf("%d@test.com", id),
		Organization: fmt.Sprintf("Organization%d", id),
		Description:  fmt.Sprintf("Description%d", id),
	}
}

// CreateTestUserAccount fill all fields based on given value
func CreateTestUserAccount(id int) domain.UserAccount {
	return domain.UserAccount{
		ID:           int64(id),
		IssuerCode:   domain.IssuerCode(id),
		SubjectUID:   fmt.Sprintf("SubjectUID%d", id),
		EmailAddress: fmt.Sprintf("%d@test.com", id),
	}
}

// CreateTestParticipant fill all fields based on given value
func CreateTestParticipant(id int) domain.Participant {
	return domain.Participant{
		ID:             int64(id),
		FirstName:      fmt.Sprintf("FirstName%d", id),
		LastName:       fmt.Sprintf("LastName%d", id),
		ContactEmail:   fmt.Sprintf("%d@test.com", id),
		Organization:   fmt.Sprintf("Organization%d", id),
		Description:    fmt.Sprintf("Description%d", id),
		AvatarFileName: fmt.Sprintf("AvatarFileName%d", id),
	}
}
