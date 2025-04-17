package dmodel

import (
	"database/sql"

	"event-calendar/internal/domain"
	"event-calendar/internal/dto/dmodel"
	"event-calendar/internal/dto/smodel"
)

// UserToUserDto
// IMPORTANT: ignore field Roles, CreatedAt.
func UserToUserDto(user domain.User) dmodel.User {
	return dmodel.User{
		ID:          user.ID,
		FirebaseUID: user.FirebaseUID,
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
		ID:          user.ID,
		FirebaseUID: user.FirebaseUID,
		Description: user.Description.String,
		CreatedAt:   user.CreatedAt,
	}
}

// MapUserDtos maps Smodel.UserID to Dmodel.UserID
// IMPORTANT: ignore field Roles, CreatedAt.
func MapUserDtos(user smodel.User) dmodel.User {
	return dmodel.User{
		ID:          user.ID,
		FirebaseUID: user.FirebaseUID,
		Description: func(desc string) sql.NullString {
			return sql.NullString{
				String: desc,
				Valid:  len(desc) > 0,
			}
		}(user.Description),
	}
}
