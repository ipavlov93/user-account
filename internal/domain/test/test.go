package test

import (
	"event-calendar/internal/domain"
	"fmt"
)

// CreateTestUser fill all fields based on given value
func CreateTestUser(id int) domain.User {
	return domain.User{
		ID:           int64(id),
		UUID:         fmt.Sprint(id),
		FirstName:    fmt.Sprintf("FirstName%d", id),
		LastName:     fmt.Sprintf("LastName%d", id),
		EmailAddress: fmt.Sprintf("%d@test.com", id),
		Company:      fmt.Sprintf("Company%d", id),
		Description:  fmt.Sprintf("Description%d", id),
	}
}
