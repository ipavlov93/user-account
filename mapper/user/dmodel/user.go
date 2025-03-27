package dmodel

import (
	"event-calendar/dmodel"
	"event-calendar/internal/domain"
	"event-calendar/smodel"
)

func UserToUserDto(user domain.User) dmodel.User {
	return dmodel.User{
		ID:           user.ID,
		UUID:         user.UUID,
		FirstName:    user.FirstName,
		LastName:     user.LastName,
		EmailAddress: user.EmailAddress,
		Company:      user.Company,
		Description:  user.Description,
	}
}

func UserDtoToUser(user dmodel.User) domain.User {
	return domain.User{
		ID:           user.ID,
		UUID:         user.UUID,
		FirstName:    user.FirstName,
		LastName:     user.LastName,
		EmailAddress: user.EmailAddress,
		Company:      user.Company,
		Description:  user.Description,
	}
}

// MapDto maps Smodel.User to Dmodel.User
func MapDto(user smodel.User) dmodel.User {
	return dmodel.User{
		ID:           user.ID,
		UUID:         user.UUID,
		FirstName:    user.FirstName,
		LastName:     user.LastName,
		EmailAddress: user.EmailAddress.Address,
		Company:      user.Company,
		Description:  user.Description,
	}
}
