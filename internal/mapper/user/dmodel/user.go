package dmodel

import (
	"database/sql"
	"user-account/internal/domain"
	"user-account/internal/dto/dmodel"
	"user-account/internal/dto/smodel"
)

// UserToUserDto
// IMPORTANT: ignore field Roles, CreatedAt.
func UserToUserDto(user domain.User) dmodel.User {
	return dmodel.User{
		ID:           user.ID,
		FirebaseUUID: user.FirebaseUUID,
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
		FirebaseUUID: user.FirebaseUUID,
		Description:  user.Description.String,
		CreatedAt:    user.CreatedAt,
	}
}

// MapUserDtos maps Smodel.User to Dmodel.User
// IMPORTANT: ignore field Roles, CreatedAt.
func MapUserDtos(user smodel.User) dmodel.User {
	return dmodel.User{
		ID:           user.ID,
		FirebaseUUID: user.FirebaseUUID,
		Description: func(desc string) sql.NullString {
			return sql.NullString{
				String: desc,
				Valid:  len(desc) > 0,
			}
		}(user.Description),
	}
}
