package dmodel

import (
	"fmt"
	"time"

	"event-calendar/internal/domain"
)

type UserAccount struct {
	ID           int64             `db:"id"`
	UserID       int64             `db:"user_id"`
	IssuerCode   domain.IssuerCode `db:"issuer_code"`
	SubjectUID   string            `db:"subject_uid"` // UID set by Auth Provider
	EmailAddress string            `db:"email_address"`
	CreatedAt    time.Time         `db:"created_at"`
}

func NewUserAccount(
	issuerCode domain.IssuerCode,
	userID int64,
	subjectUID string,
	emailAddress string,
) UserAccount {
	return UserAccount{
		IssuerCode:   issuerCode,
		UserID:       userID,
		SubjectUID:   subjectUID,
		EmailAddress: emailAddress,
	}
}

func (u UserAccount) String() string {
	return fmt.Sprintf("ID:%d, UserID:%d", u.ID, u.UserID)
}
