package domain

import (
	"fmt"
	"time"

	"event-calendar/internal/domain/claims"
)

type User struct {
	ID          int64
	FirebaseUID string
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
