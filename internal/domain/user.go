package domain

import (
	"event-calendar/internal/domain/claims"
	"fmt"
)

type User struct {
	ID          int64
	FirebaseUID string
	Roles       []claims.Role
	// fields set by participant himself or identity/auth provider
	BusinessName string
	FirstName    string
	LastName     string
	EmailAddress string

	Organization string
	Description  string
}

func NewUser(
	firebaseUID string,
	businessName string,
	firstName string,
	lastName string,
	emailAddress string,
	organization string,
	description string,
	roles ...claims.Role,
) User {
	return User{
		Roles:        roles,
		FirebaseUID:  firebaseUID,
		BusinessName: businessName,
		FirstName:    firstName,
		LastName:     lastName,
		EmailAddress: emailAddress,
		Organization: organization,
		Description:  description,
	}
}

func (u User) String() string {
	return fmt.Sprintf("ID: %d, FirebaseUID: %s, BusinessName: %s, FullName: %s %s, EmailAddress: %s",
		u.ID, u.FirebaseUID, u.BusinessName, u.FirstName, u.LastName, u.EmailAddress)
}

func (u User) Equals(user *User) bool {
	if user == nil {
		return false
	}
	if u.ID != user.ID {
		return false
	}
	if u.FirebaseUID != user.FirebaseUID {
		return false
	}
	return true
}

func (u User) HasValidID() bool {
	return u.ID > 0
}
