package dmodel

import (
	"fmt"
	"time"
)

type User struct {
	ID   int64  `db:"id"`
	UUID string `db:"uuid"`
	// fields set by participant himself or auth provider
	FirstName    string `db:"first_name"`
	LastName     string `db:"last_name"`
	EmailAddress string `db:"email_address"`

	Description string    `db:"description"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
	// soft delete
	DeletedAt time.Time `db:"deleted_at"`
}

func NewUser(firstName, lastName, emailAddress string) *User {
	//mail.ParseAddress(email)

	return &User{
		FirstName:    firstName,
		LastName:     lastName,
		EmailAddress: emailAddress,
		//EmailAddress: mail.Address{
		//	Name:    fmt.Sprintf("%s %s", firstName, lastName),
		//	Address: emailAddress,
		//},
	}
}

func (p User) String() string {
	return fmt.Sprintf("ID:%s, FullName:%s %s", p.ID, p.FirstName, p.LastName)
}

func (p User) Equals(participant *User) bool {
	if p.ID != participant.ID {
		return false
	}
	return p.EmailAddress == participant.EmailAddress
}
