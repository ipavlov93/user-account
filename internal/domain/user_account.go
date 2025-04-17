package domain

import (
	"fmt"
	"strings"
)

type IssuerCode int

const (
	UnknownIssuer IssuerCode = 0
	EmailPassword IssuerCode = 1
	GoogleIssuer  IssuerCode = 2
	//AzureIssuer  IssuerCode = 3
)

var stateIssuerCodeNames = map[IssuerCode]string{
	EmailPassword: "password",
	GoogleIssuer:  "google.com",
	//GoogleIssuer: "Google",
}

func (c IssuerCode) String() string {
	s, ok := stateIssuerCodeNames[c]
	if !ok {
		return fmt.Sprintf("IssuerCode(%d)", c)
	}
	return s
}

func NewIssuerCode(issuer string) IssuerCode {
	for issuerCode, val := range stateIssuerCodeNames {
		if strings.EqualFold(val, issuer) {
			return issuerCode
		}
	}
	return UnknownIssuer
}

type UserAccount struct {
	ID           int64
	UserID       int64
	IssuerCode   IssuerCode
	SubjectUID   string // UID set by Auth Provider
	EmailAddress string
	ContactName  string
}

func NewUserAccount(
	issuerCode IssuerCode,
	userID int64,
	subjectUID string,
	email string,
	contactName string,
) UserAccount {
	return UserAccount{
		IssuerCode:   issuerCode,
		UserID:       userID,
		SubjectUID:   subjectUID,
		EmailAddress: email,
		ContactName:  contactName,
	}
}

func (u UserAccount) String() string {
	return fmt.Sprintf("ID:%d, UserID:%d, EmailAddress: %s", u.ID, u.UserID, u.EmailAddress)
}
