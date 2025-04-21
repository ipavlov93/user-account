package dmodel

import (
	"database/sql"
	"event-calendar/internal/domain"
	"event-calendar/internal/dto/dmodel"
)

// UserAccountToDto
// IMPORTANT: ignore field CreatedAt.
func UserAccountToDto(userAccount domain.UserAccount) dmodel.UserAccount {
	return dmodel.UserAccount{
		ID:           userAccount.ID,
		UserID:       userAccount.UserID,
		IssuerCode:   userAccount.IssuerCode.String(),
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
		IssuerCode:   domain.NewIssuerCode(userAccount.IssuerCode),
		SubjectUID:   userAccount.SubjectUID,
		EmailAddress: userAccount.EmailAddress,
		ContactName:  userAccount.ContactName.String,
	}
}

func DtosToUserAccounts(userAccounts []dmodel.UserAccount) []domain.UserAccount {
	result := make([]domain.UserAccount, 0, len(userAccounts))
	for _, userAccount := range userAccounts {
		result = append(result, domain.UserAccount{
			ID:           userAccount.ID,
			UserID:       userAccount.UserID,
			IssuerCode:   domain.NewIssuerCode(userAccount.IssuerCode),
			SubjectUID:   userAccount.SubjectUID,
			EmailAddress: userAccount.EmailAddress,
			ContactName:  userAccount.ContactName.String,
		})
	}
	return result
}
