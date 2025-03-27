package dmodel

import (
	"event-calendar/dmodel"
	"event-calendar/internal/domain"
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
