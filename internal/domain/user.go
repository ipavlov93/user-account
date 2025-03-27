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

	Company     string
	Description string
}

func NewUser(uuid string, firstName, lastName, emailAddress, company, description string) *User {
	return &User{
		UUID:         uuid,
		FirstName:    firstName,
		LastName:     lastName,
		EmailAddress: emailAddress,
		Company:      company,
		Description:  description,
	}
}

func (p User) String() string {
	return fmt.Sprintf("ID:%d, FullName:%s %s", p.ID, p.FirstName, p.LastName)
}

func (p User) Equals(participant *User) bool {
	if p.ID != participant.ID {
		return false
	}
	if p.UUID != participant.UUID {
		return false
	}
	if !strings.EqualFold(p.FirstName, participant.FirstName) {
		return false
	}
	if !strings.EqualFold(p.LastName, participant.LastName) {
		return false
	}
	return strings.EqualFold(p.EmailAddress, participant.EmailAddress)
}
