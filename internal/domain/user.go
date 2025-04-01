package domain

import (
	"fmt"
	"strings"
)

type User struct {
	ID   int64
	UUID string
	// fields set by participant himself or auth provider
	FirstName    string
	LastName     string
	EmailAddress string

	Organization string
	Description  string
}

func NewUser(
	uuid string,
	firstName string,
	lastName string,
	emailAddress string,
	organization string,
	description string,
) User {
	return User{
		UUID:         uuid,
		FirstName:    firstName,
		LastName:     lastName,
		EmailAddress: emailAddress,
		Organization: organization,
		Description:  description,
	}
}

func (u User) String() string {
	return fmt.Sprintf("ID:%d, FullName:%s %s", u.ID, u.FirstName, u.LastName)
}

func (u User) Equals(participant *User) bool {
	if u.ID != participant.ID {
		return false
	}
	if u.UUID != participant.UUID {
		return false
	}
	if !strings.EqualFold(u.FirstName, participant.FirstName) {
		return false
	}
	if !strings.EqualFold(u.LastName, participant.LastName) {
		return false
	}
	return strings.EqualFold(u.EmailAddress, participant.EmailAddress)
}

func (u User) HasValidID() bool {
	return u.ID > 0
}
