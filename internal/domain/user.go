package domain

import (
	"fmt"
	"time"

	"event-calendar/internal/domain/claims"
)

// User represents the registered (authenticated at least once) user of Bookly application.
type User struct {
	ID          int64
	FirebaseUID string // UUID of the authenticated user stored in the Firestore
	Roles       []claims.Role
	Description string
	CreatedAt   time.Time
}

func NewUser(
	firebaseUID string,
	description string,
	roles ...claims.Role,
) User {
	return User{
		Roles:       roles,
		FirebaseUID: firebaseUID,
		Description: description,
	}
}

func (u User) String() string {
	return fmt.Sprintf("ID: %d, FirebaseUID: %s", u.ID, u.FirebaseUID)
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
	if u.FirebaseUID != user.FirebaseUID {
		return false
	}

	if len(u.Roles) != len(user.Roles) || !claims.IsGivenRolesPresent(user.Roles, u.Roles) {
		return false
	}
	return true
}

func (u User) HasValidID() bool {
	return u.ID > 0
}
