package smodel

import (
	"net/mail"

	"user-account/internal/domain"
	"user-account/internal/dto/dmodel"
	"user-account/internal/dto/smodel"
)

func ProfileToProfileDto(profile domain.UserProfile) smodel.UserProfile {
	return smodel.UserProfile{
		ID:     profile.ID,
		UserID: profile.ID,
		ContactEmail: mail.Address{
			Name:    profile.BusinessName,
			Address: profile.ContactEmail,
		},
		FirstName:      profile.FirstName,
		LastName:       profile.LastName,
		Organization:   profile.Organization,
		BusinessName:   profile.BusinessName,
		Description:    profile.Description,
		AvatarFileName: profile.AvatarFileName,
		CreatedAt:      profile.CreatedAt,
	}
}

// IMPORTANT: ignore CreatedAt.
func ProfileDtoToProfile(profile smodel.UserProfile) domain.UserProfile {
	return domain.UserProfile{
		ID:           profile.ID,
		UserID:       profile.UserID,
		FirstName:    profile.FirstName,
		LastName:     profile.LastName,
		ContactEmail: profile.ContactEmail.Address,
		Organization: profile.Organization,
		BusinessName: profile.BusinessName,
		Description:  profile.Description,
	}
}

// MapUserProfileDtos maps Dmodel.UserProfile to Smodel.UserProfile.
func MapUserProfileDtos(profile dmodel.UserProfile) smodel.UserProfile {
	return smodel.UserProfile{
		ID:        profile.ID,
		UserID:    profile.UserID,
		FirstName: profile.FirstName,
		LastName:  profile.LastName,
		ContactEmail: mail.Address{
			Name:    profile.BusinessName.String,
			Address: profile.ContactEmail,
		},
		Organization: profile.Organization,
		BusinessName: profile.BusinessName.String,
		Description:  profile.Description.String,
		CreatedAt:    profile.CreatedAt,
	}
}
