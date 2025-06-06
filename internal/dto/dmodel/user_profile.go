package dmodel

import (
	"database/sql"
	"fmt"
	"time"
)

type UserProfile struct {
	ID int64 `db:"id"`
	// fields set by participant himself or auth provider
	FirstName      string         `db:"first_name"`
	LastName       string         `db:"last_name"`
	BusinessName   sql.NullString `db:"business_name"`
	ContactEmail   string         `db:"contact_email"`
	UserID         int64          `db:"user_id"`
	Organization   string         `db:"organization"`
	Description    sql.NullString `db:"description"`
	AvatarFileName sql.NullString `db:"avatar_file_name"`
	CreatedAt      time.Time      `db:"created_at"`
	UpdatedAt      time.Time      `db:"updated_at"`
	// soft delete
	DeletedAt sql.NullTime `db:"deleted_at"`
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
