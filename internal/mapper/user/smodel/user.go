package smodel

import (
	"event-calendar/internal/domain"
	"event-calendar/internal/dto/dmodel"
	"event-calendar/internal/dto/smodel"
)

func UserToUserDto(user domain.User) smodel.User {
	return smodel.User{
		ID:          user.ID,
		FirebaseUID: user.FirebaseUID,
		Description: user.Description,
		CreatedAt:   user.CreatedAt,
	}
}

// UserDtoToUser
// IMPORTANT: ignore CreatedAt.
func UserDtoToUser(user smodel.User) domain.User {
	return domain.User{
		ID:          user.ID,
		FirebaseUID: user.FirebaseUID,
		Description: user.Description,
	}
}

// MapDto maps Dmodel.User to Smodel.User
func MapDto(user dmodel.User) smodel.User {
	return smodel.User{
		ID:          user.ID,
		FirebaseUID: user.FirebaseUID,
		Description: user.Description.String,
		CreatedAt:   user.CreatedAt,
	}
}
