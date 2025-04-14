package dmodel

import (
	"database/sql"
	"fmt"
	"time"
)

type User struct {
	ID           int64  `db:"id"`
	FirebaseUID  string `db:"firebase_uid"`
	EmailAddress string `db:"email_address"`
	// fields set by user himself
	BusinessName sql.NullString `db:"business_name"`
	FirstName    sql.NullString `db:"first_name"`
	LastName     sql.NullString `db:"last_name"`
	Organization sql.NullString `db:"organization"`
	Description  sql.NullString `db:"description"`
	CreatedAt    time.Time      `db:"created_at"`
	UpdatedAt    time.Time      `db:"updated_at"`
	// soft delete
	DeletedAt sql.NullTime `db:"deleted_at"`
}

func (p User) String() string {
	return fmt.Sprintf("ID: %d, FirebaseUID: %s, EmailAddress: %s",
		p.ID, p.FirebaseUID, p.EmailAddress)
}
