package domain

import (
	"fmt"
	"strings"
)

// issuer is a domain-specific enumeration of known issuers (Auth providers).
// It is unexported to prevent invalid or inconsistent values being created outside this package.
// It enforces safe construction via NewIssuerCode() and safe setting value via SetValidIssuer().
type issuer string

const (
	UnknownIssuer issuer = "UNKNOWN"
	EmailPassword issuer = "EMAIL_PASSWORD"
	GoogleIssuer  issuer = "GOOGLE"
)

var stateIssuerCodeNames = map[issuer]string{
	EmailPassword: "password",
	GoogleIssuer:  "google.com", // "Google",
}

func (c issuer) String() string {
	s, ok := stateIssuerCodeNames[c]
	if !ok {
		return fmt.Sprintf("Issuer(%s)", string(c))
	}
	return s
}

// NewIssuerCode creates a valid issuerCode instance.
// If the input string does not match a known issuer name (not recognized), it returns UnknownIssuer.
func NewIssuerCode(iss string) issuer {
	return SetValidIssuer(issuer(iss))
}

// SetValidIssuer ensures that an issuer value is valid.
// If issuer's value is absent in the known issuer list, returns UnknownIssuer.
// Useful when input is already an issuer type but needs validation.
func SetValidIssuer(iss issuer) issuer {
	for issuerName, val := range stateIssuerCodeNames {
		if strings.EqualFold(val, string(iss)) {
			return issuerName
		}
	}
	return UnknownIssuer
}

// UserAccount represents the registered (authenticated at least once) user of Bookly application.
type UserAccount struct {
	ID           int64
	UserID       int64
	IssuerCode   issuer
	SubjectUID   string // UID set by Auth Provider
	EmailAddress string
	ContactName  string
}

// NewUserAccount safely constructs a new UserAccount instance,
// mapping input strings to domain-safe types like issuerCode via constructors.
func NewUserAccount(
	issuerCode string,
	userID int64,
	subjectUID string,
	email string,
	contactName string,
) UserAccount {
	return UserAccount{
		IssuerCode:   NewIssuerCode(issuerCode),
		UserID:       userID,
		SubjectUID:   subjectUID,
		EmailAddress: email,
		ContactName:  contactName,
	}
}

func (u UserAccount) String() string {
	return fmt.Sprintf("ID:%d, UserID:%d, EmailAddress: %s", u.ID, u.UserID, u.EmailAddress)
}
