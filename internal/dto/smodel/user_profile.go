package smodel

import (
	"fmt"
	"net/mail"
	"time"
)

type UserProfile struct {
	ID int64 `json:"id"`
	// fields set by participant himself or auth provider
	FirstName      string       `json:"first_name"`
	LastName       string       `json:"last_name"`
	BusinessName   string       `json:"business_name"`
	ContactEmail   mail.Address `json:"contact_email"`
	UserID         int64        `json:"user_id"`
	Organization   string       `json:"organization"`
	Description    string       `json:"description"`
	AvatarFileName string       `json:"avatar_file_name"`
	CreatedAt      time.Time    `json:"created_at"`
}

func (p UserProfile) String() string {
	return fmt.Sprintf(
		"ID: %d, FullName: %s %s, EmailAddress: %s",
		p.ID,
		p.FirstName,
		p.LastName,
		p.ContactEmail,
	)
}
