package smodel

import (
	"event-calendar/internal/domain"
	"event-calendar/smodel"
	"fmt"
	"net/mail"
)

func UserToUserDto(user domain.User) smodel.User {
	return smodel.User{
		ID:        user.ID,
		UUID:      user.UUID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		EmailAddress: mail.Address{
			Name:    fmt.Sprintf("%s %s", user.FirstName, user.LastName),
			Address: user.EmailAddress,
		},
		Company:     user.Company,
		Description: user.Description,
	}
}

func UserDtoToUser(user smodel.User) domain.User {
	return domain.User{
		ID:           user.ID,
		UUID:         user.UUID,
		FirstName:    user.FirstName,
		LastName:     user.LastName,
		EmailAddress: user.EmailAddress.Address,
		Company:      user.Company,
		Description:  user.Description,
	}
}
