package dmodel

import (
	"database/sql"

	"user-account/internal/domain"
	"user-account/internal/dto/dmodel"
)

// UserAccountToDto
// IMPORTANT: ignore field CreatedAt.
func UserAccountToDto(userAccount domain.UserAccount) dmodel.UserAccount {
	return dmodel.UserAccount{
		ID:           userAccount.ID,
		UserID:       userAccount.UserID,
		Issuer:       userAccount.Issuer.String(),
		SubjectUID:   userAccount.SubjectUID,
		EmailAddress: userAccount.EmailAddress,
		ContactName: sql.NullString{
			Valid:  len(userAccount.ContactName) > 0,
			String: userAccount.ContactName,
		},
	}
}

func DtoToUserAccount(userAccount dmodel.UserAccount) domain.UserAccount {
	return domain.UserAccount{
		ID:           userAccount.ID,
		UserID:       userAccount.UserID,
		Issuer:       domain.NewIssuer(userAccount.Issuer),
		SubjectUID:   userAccount.SubjectUID,
		EmailAddress: userAccount.EmailAddress,
		ContactName:  userAccount.ContactName.String,
	}
}

// MapUserAccounts maps []dmodel.User to []domain.User
func MapUserAccounts(userAccounts []dmodel.UserAccount) []domain.UserAccount {
	result := make([]domain.UserAccount, 0, len(userAccounts))
	for _, userAccount := range userAccounts {
		result = append(result, domain.UserAccount{
			ID:           userAccount.ID,
			UserID:       userAccount.UserID,
			Issuer:       domain.NewIssuer(userAccount.Issuer),
			SubjectUID:   userAccount.SubjectUID,
			EmailAddress: userAccount.EmailAddress,
			ContactName:  userAccount.ContactName.String,
		})
	}
	return result
}
