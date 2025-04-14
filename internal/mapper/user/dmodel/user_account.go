package dmodel

import (
	"event-calendar/internal/domain"
	"event-calendar/internal/dto/dmodel"
)

func UserAccountToDto(userAccount domain.UserAccount) dmodel.UserAccount {
	return dmodel.UserAccount{
		ID:           userAccount.ID,
		UserID:       userAccount.UserID,
		IssuerCode:   userAccount.IssuerCode,
		SubjectUID:   userAccount.SubjectUID,
		EmailAddress: userAccount.EmailAddress,
	}
}

func DtoToUserAccount(userAccount dmodel.UserAccount) domain.UserAccount {
	return domain.UserAccount{
		ID:           userAccount.ID,
		UserID:       userAccount.UserID,
		IssuerCode:   userAccount.IssuerCode,
		SubjectUID:   userAccount.SubjectUID,
		EmailAddress: userAccount.EmailAddress,
	}
}

func DtosToUserAccounts(userAccounts []dmodel.UserAccount) []domain.UserAccount {
	result := make([]domain.UserAccount, 0, len(userAccounts))
	for _, userAccount := range userAccounts {
		result = append(result, domain.UserAccount{
			ID:           userAccount.ID,
			UserID:       userAccount.UserID,
			IssuerCode:   userAccount.IssuerCode,
			SubjectUID:   userAccount.SubjectUID,
			EmailAddress: userAccount.EmailAddress,
		})
	}
	return result
}
