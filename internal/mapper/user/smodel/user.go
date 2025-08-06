package smodel

import (
	"user-account/internal/domain"
	"user-account/internal/dto/dmodel"
	"user-account/internal/dto/smodel"
)

func UserToUserDto(user domain.User) smodel.User {
	return smodel.User{
		ID:           user.ID,
		FirebaseUUID: user.FirebaseUUID,
		Description:  user.Description,
		CreatedAt:    user.CreatedAt,
	}
}

// UserDtoToUser
// IMPORTANT: ignore CreatedAt.
func UserDtoToUser(user smodel.User) domain.User {
	return domain.User{
		ID:           user.ID,
		FirebaseUUID: user.FirebaseUUID,
		Description:  user.Description,
	}
}

// MapDto maps Dmodel.User to Smodel.User
func MapDto(user dmodel.User) smodel.User {
	return smodel.User{
		ID:           user.ID,
		FirebaseUUID: user.FirebaseUUID,
		Description:  user.Description.String,
		CreatedAt:    user.CreatedAt,
	}
}
