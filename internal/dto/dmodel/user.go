package dmodel

import (
	"database/sql"
	"fmt"
	"time"
)

type User struct {
	ID           int64          `db:"id"`
	FirebaseUUID string         `db:"firebase_uuid"`
	Description  sql.NullString `db:"description"`
	CreatedAt    time.Time      `db:"created_at"`
	UpdatedAt    time.Time      `db:"updated_at"`
	// soft delete
	DeletedAt sql.NullTime `db:"deleted_at"`
}

func (p User) String() string {
	return fmt.Sprintf("ID: %d, FirebaseUUID: %s", p.ID, p.FirebaseUUID)
}
