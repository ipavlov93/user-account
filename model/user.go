package model

import (
	"fmt"
	"net/mail"
	"time"
)

type User struct {
	ID string // type string to store hash or GUID
	// fields set by participant himself or auth provider
	FirstName    string
	LastName     string
	EmailAddress mail.Address

	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   time.Time // soft delete
}

func NewUser(firstName, lastName, email string) *User {
	//mail.ParseAddress(email)

	return &User{
		FirstName: firstName,
		LastName:  lastName,
		EmailAddress: mail.Address{
			Name:    fmt.Sprintf("%s %s", firstName, lastName),
			Address: email,
		},
	}
}

func (p User) String() string {
	return fmt.Sprintf("ID:%s, FullName:%s %s", p.ID, p.FirstName, p.LastName)
}

func (p User) Equals(participant *User) bool {
	if p.ID != participant.ID {
		return false
	}
	return p.EmailAddress.Address == participant.EmailAddress.Address
}
