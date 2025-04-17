package domain

import (
	"fmt"
	"time"
)

type UserProfile struct {
	ID int64
	// fields set by participant himself
	FirstName      string
	LastName       string
	BusinessName   string
	ContactEmail   string
	UserID         int64
	Organization   string
	Description    string
	AvatarFileName string
	CreatedAt      time.Time
}

func NewUserProfile(
	userID int64,
	firstName string,
	lastName string,
	businessName string,
	emailAddress string,
	organization string,
	description string,
	avatarFileName string,
) UserProfile {
	return UserProfile{
		UserID:         userID,
		FirstName:      firstName,
		LastName:       lastName,
		BusinessName:   businessName,
		ContactEmail:   emailAddress,
		Organization:   organization,
		Description:    description,
		AvatarFileName: avatarFileName,
	}
}

func (p UserProfile) String() string {
	return fmt.Sprintf("ID:%d, FullName:%s %s", p.ID, p.FirstName, p.LastName)
}

func (p UserProfile) Equals(profile *UserProfile) bool {
	if profile == nil {
		return false
	}
	if p.ID != profile.ID {
		return false
	}
	return p.UserID == profile.UserID
}
