package dmodel

import (
	"fmt"
	"time"
)

type Participant struct {
	ID int64
	// fields set by participant himself or auth provider
	FirstName      string
	LastName       string
	ContactEmail   string
	AvatarFileName string

	UserID      int64
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	// soft delete
	DeletedAt time.Time
}

func NewParticipant(firstName, lastName, emailAddress string, userID ...int64) *Participant {
	uID := int64(0)
	if len(userID) > 0 {
		uID = userID[0]
	}

	return &Participant{
		FirstName:    firstName,
		LastName:     lastName,
		ContactEmail: emailAddress,
		UserID:       uID,
	}
}

func (p Participant) String() string {
	return fmt.Sprintf("ID:%d, FullName:%s %s", p.ID, p.FirstName, p.LastName)
}

func (p Participant) Equals(participant *Participant) bool {
	if p.ID != participant.ID {
		return false
	}
	return p.ContactEmail == participant.ContactEmail
}
