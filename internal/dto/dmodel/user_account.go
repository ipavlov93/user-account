package dmodel

import (
	"database/sql"
	"fmt"
	"time"
)

type UserAccount struct {
	ID           int64          `db:"id"`
	UserID       int64          `db:"user_id"`
	IssuerCode   string         `db:"issuer"`
	SubjectUID   string         `db:"subject_uid"` // UID set by Auth Provider
	EmailAddress string         `db:"email_address"`
	ContactName  sql.NullString `db:"contact_name"`
	CreatedAt    time.Time      `db:"created_at"`
}

func NewUserAccount(
	issuerCode string,
	userID int64,
	subjectUID string,
	emailAddress string,
	contactName string,
) UserAccount {
	return UserAccount{
		IssuerCode:   issuerCode,
		UserID:       userID,
		SubjectUID:   subjectUID,
		EmailAddress: emailAddress,
		ContactName: sql.NullString{
			Valid:  len(contactName) > 0,
			String: contactName,
		},
	}
}

func (u UserAccount) String() string {
	return fmt.Sprintf("ID:%d, UserID:%d", u.ID, u.UserID)
}
