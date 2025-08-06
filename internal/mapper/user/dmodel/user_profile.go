package dmodel

import (
	"database/sql"
	"user-account/internal/domain"
	"user-account/internal/dto/dmodel"
	"user-account/internal/dto/smodel"
)

// ProfileToProfileDto
// IMPORTANT: ignore field CreatedAt.
func ProfileToProfileDto(profile domain.UserProfile) dmodel.UserProfile {
	return dmodel.UserProfile{
		ID:           profile.ID,
		UserID:       profile.ID,
		ContactEmail: profile.ContactEmail,
		FirstName:    profile.FirstName,
		LastName:     profile.LastName,
		Organization: profile.Organization,
		BusinessName: func(businessName string) sql.NullString {
			return sql.NullString{
				String: businessName,
				Valid:  len(businessName) > 0,
			}
		}(profile.BusinessName),
		Description: func(desc string) sql.NullString {
			return sql.NullString{
				String: desc,
				Valid:  len(desc) > 0,
			}
		}(profile.Description),
		AvatarFileName: func(avatarFileName string) sql.NullString {
			return sql.NullString{
				String: avatarFileName,
				Valid:  len(avatarFileName) > 0,
			}
		}(profile.AvatarFileName),
	}
}

func ProfileDtoToProfile(profile dmodel.UserProfile) domain.UserProfile {
	return domain.UserProfile{
		ID:           profile.ID,
		UserID:       profile.UserID,
		FirstName:    profile.FirstName,
		LastName:     profile.LastName,
		ContactEmail: profile.ContactEmail,
		Organization: profile.Organization,
		BusinessName: profile.BusinessName.String,
		Description:  profile.Description.String,
		CreatedAt:    profile.CreatedAt,
	}
}

// MapUserProfileDtos maps Smodel.UserProfile to Dmodel.UserProfile.
// IMPORTANT: ignore field CreatedAt.
func MapUserProfileDtos(profile smodel.UserProfile) dmodel.UserProfile {
	return dmodel.UserProfile{
		ID:           profile.ID,
		UserID:       profile.UserID,
		FirstName:    profile.FirstName,
		LastName:     profile.LastName,
		ContactEmail: profile.ContactEmail.Address,
		Organization: profile.Organization,
		BusinessName: func(businessName string) sql.NullString {
			return sql.NullString{
				String: businessName,
				Valid:  len(businessName) > 0,
			}
		}(profile.BusinessName),
		Description: func(desc string) sql.NullString {
			return sql.NullString{
				String: desc,
				Valid:  len(desc) > 0,
			}
		}(profile.Description),
	}
}
