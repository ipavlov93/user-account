package smodel

import (
	"fmt"
	"time"
)

type User struct {
	ID           int64     `json:"id"`
	FirebaseUUID string    `json:"uuid"`
	Description  string    `json:"description"`
	CreatedAt    time.Time `json:"createdAt"`
}

func (p User) String() string {
	return fmt.Sprintf("ID: %d, FirebaseUUID: %s", p.ID, p.FirebaseUUID)
}
