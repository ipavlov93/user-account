package smodel

import (
	"fmt"
	"net/mail"

	"event-calendar/internal/domain"
	"event-calendar/internal/dto/dmodel"
	"event-calendar/internal/dto/smodel"
)

func UserToUserDto(user domain.User) smodel.User {
	return smodel.User{
		ID:          user.ID,
		FirebaseUID: user.FirebaseUID,
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		EmailAddress: mail.Address{
			Name:    fmt.Sprintf("%s %s", user.FirstName, user.LastName),
			Address: user.EmailAddress,
		},
		Organization: user.Organization,
		Description:  user.Description,
	}
}

func UserDtoToUser(user smodel.User) domain.User {
	return domain.User{
		ID:           user.ID,
		FirebaseUID:  user.FirebaseUID,
		FirstName:    user.FirstName,
		LastName:     user.LastName,
		EmailAddress: user.EmailAddress.Address,
		Organization: user.Organization,
		Description:  user.Description,
	}
}

// MapDto maps Dmodel.User to Smodel.User
func MapDto(user dmodel.User) smodel.User {
	return smodel.User{
		ID:          user.ID,
		FirebaseUID: user.FirebaseUID,
		FirstName:   user.FirstName.String,
		LastName:    user.LastName.String,
		EmailAddress: mail.Address{
			Name:    fmt.Sprintf("%s %s", user.FirstName, user.LastName),
			Address: user.EmailAddress,
		},
		Organization: user.Organization.String,
		Description:  user.Description.String,
	}
}
