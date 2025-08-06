package domain

import (
	"fmt"
	"time"

	"user-account/internal/domain/role"
)

// User represents the registered (authenticated at least once) user
type User struct {
	ID           int64
	FirebaseUUID string // UUID of the authenticated user stored in the Firestore
	Roles        []role.Role
	Description  string
	CreatedAt    time.Time
}

func NewUser(
	firebaseUUID string,
	description string,
	roles ...role.Role,
) User {
	return User{
		Roles:        roles,
		FirebaseUUID: firebaseUUID,
		Description:  description,
	}
}

func (u User) String() string {
	return fmt.Sprintf("ID: %d, FirebaseUUID: %s", u.ID, u.FirebaseUUID)
}

// Equals compare critical user fields with given user.
// Users will be considered as distinct if scope of roles isn't identical.
func (u User) Equals(user *User) bool {
	if user == nil {
		return false
	}
	if u.ID != user.ID {
		return false
	}
	if u.FirebaseUUID != user.FirebaseUUID {
		return false
	}

	if len(u.Roles) != len(user.Roles) || !role.IsGivenRolesPresent(user.Roles, u.Roles) {
		return false
	}
	return true
}

func (u User) HasValidID() bool {
	return u.ID > 0
}
