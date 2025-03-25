package dmodel

import (
	"fmt"
	"net/mail"
	"time"
)

type Participant struct {
	ID int64
	// fields set by participant himself or auth provider
	FirstName      string
	LastName       string
	ContactEmail   mail.Address
	AvatarFileName string

	User        User
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   time.Time // soft delete
}

func NewParticipant(firstName, lastName, emailAddress string, userID ...int64) *Participant {
	//mail.ParseAddress(email)

	uID := int64(0)
	if len(userID) > 0 {
		uID = userID[0]
	}

	return &Participant{
		FirstName: firstName,
		LastName:  lastName,
		ContactEmail: mail.Address{
			Name:    fmt.Sprintf("%s %s", firstName, lastName),
			Address: emailAddress,
		},
		User: User{
			ID: uID,
		},
	}
}

func (p Participant) String() string {
	return fmt.Sprintf("ID:%s, FullName:%s %s", p.ID, p.FirstName, p.LastName)
}

func (p Participant) Equals(participant *Participant) bool {
	if p.ID != participant.ID {
		return false
	}
	return p.ContactEmail.Address == participant.ContactEmail.Address
}
