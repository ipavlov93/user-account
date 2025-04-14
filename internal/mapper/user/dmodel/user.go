package dmodel

import (
	"database/sql"

	"event-calendar/internal/domain"
	"event-calendar/internal/dto/dmodel"
	"event-calendar/internal/dto/smodel"
)

func UserToUserDto(user domain.User) dmodel.User {
	return dmodel.User{
		ID:          user.ID,
		FirebaseUID: user.FirebaseUID,
		BusinessName: func(businessName string) sql.NullString {
			return sql.NullString{
				String: businessName,
				Valid:  len(businessName) > 0,
			}
		}(user.BusinessName),
		FirstName: func(firstName string) sql.NullString {
			return sql.NullString{
				String: firstName,
				Valid:  len(firstName) > 0,
			}
		}(user.FirstName),
		LastName: func(lastName string) sql.NullString {
			return sql.NullString{
				String: lastName,
				Valid:  len(lastName) > 0,
			}
		}(user.LastName),
		EmailAddress: user.EmailAddress,
		Organization: func(organization string) sql.NullString {
			return sql.NullString{
				String: organization,
				Valid:  len(organization) > 0,
			}
		}(user.Organization),
		Description: func(desc string) sql.NullString {
			return sql.NullString{
				String: desc,
				Valid:  len(desc) > 0,
			}
		}(user.Description),
	}
}

func UserDtoToUser(user dmodel.User) domain.User {
	return domain.User{
		ID:           user.ID,
		FirebaseUID:  user.FirebaseUID,
		BusinessName: user.BusinessName.String,
		FirstName:    user.FirstName.String,
		LastName:     user.LastName.String,
		EmailAddress: user.EmailAddress,
		Organization: user.Organization.String,
		Description:  user.Description.String,
	}
}

// MapDto maps Smodel.User to Dmodel.User
func MapDto(user smodel.User) dmodel.User {
	return dmodel.User{
		ID:          user.ID,
		FirebaseUID: user.FirebaseUID,
		BusinessName: func(businessName string) sql.NullString {
			return sql.NullString{
				String: businessName,
				Valid:  len(businessName) > 0,
			}
		}(user.BusinessName),
		FirstName: func(firstName string) sql.NullString {
			return sql.NullString{
				String: firstName,
				Valid:  len(firstName) > 0,
			}
		}(user.FirstName),
		LastName: func(lastName string) sql.NullString {
			return sql.NullString{
				String: lastName,
				Valid:  len(lastName) > 0,
			}
		}(user.LastName),
		EmailAddress: user.EmailAddress.Address,
		Organization: func(organization string) sql.NullString {
			return sql.NullString{
				String: organization,
				Valid:  len(organization) > 0,
			}
		}(user.Organization),
		Description: func(desc string) sql.NullString {
			return sql.NullString{
				String: desc,
				Valid:  len(desc) > 0,
			}
		}(user.Description),
	}
}
