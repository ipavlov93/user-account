package smodel

import (
	"fmt"
	"net/mail"
)

type User struct {
	ID           int64        `json:"id"`
	FirebaseUID  string       `json:"uuid"`
	EmailAddress mail.Address `json:"email_address"`
	// fields set by user himself
	BusinessName string `json:"businessName"`
	FirstName    string `json:"firstName"`
	LastName     string `json:"lastName"`
	Organization string `json:"organization"`
	Description  string `json:"description"`
}

func (p User) String() string {
	return fmt.Sprintf("ID: %d, FirebaseUID: %s, BusinessName: %s, FullName: %s %s, EmailAddress: %s",
		p.ID, p.FirebaseUID, p.BusinessName, p.FirstName, p.LastName, p.EmailAddress)
}
