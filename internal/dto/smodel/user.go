package smodel

import (
	"fmt"
	"time"
)

type User struct {
	ID          int64     `json:"id"`
	FirebaseUID string    `json:"uuid"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"createdAt"`
}

func (p User) String() string {
	return fmt.Sprintf("ID: %d, FirebaseUID: %s", p.ID, p.FirebaseUID)
}
