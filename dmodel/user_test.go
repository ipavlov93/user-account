package dmodel_test

import (
	"event-calendar/dmodel"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewUser(t *testing.T) {
	// ARRANGE
	user1 := createDummyUser(1)

	// happy flow
	t.Run("should create user and init giving fields", func(t *testing.T) {
		// ACT
		user2 := dmodel.NewUser(
			user1.FirstName,
			user1.LastName,
			user1.EmailAddress,
		)

		// ASSERT
		assert.Equal(t, user2.LastName, user1.LastName)
		assert.Equal(t, user2.LastName, user1.LastName)
		assert.Equal(t, user2.LastName, user1.LastName)

		//if !user2.Equals(user1) {
		//	t.Errorf("got %s but want %s", user2, user1)
		//}
	})
}

// fill all the string fields with given value
func createDummyUser(id int) *dmodel.User {
	return &dmodel.User{
		ID:           int64(id),
		FirstName:    fmt.Sprintf("FirstName %d", id),
		LastName:     fmt.Sprintf("LastName %d", id),
		EmailAddress: fmt.Sprintf("%d@test.com", id),
		//EmailAddress: mail.Address{
		//	Name:    fmt.Sprintf("<Contact %d>", id),
		//	Address: fmt.Sprintf("%d@test.com", id),
		//},
	}
}
