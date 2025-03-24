package model_test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/mail"
	"secret-santa/model"
	"testing"
)

func TestNewUser(t *testing.T) {
	// ARRANGE
	user1 := createDummyUser("1")

	// happy flow
	t.Run("should create user and init giving fields", func(t *testing.T) {
		// ACT
		user2 := model.NewUser(
			user1.FirstName,
			user1.LastName,
			user1.EmailAddress.Address,
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
func createDummyUser(id string) *model.User {
	return &model.User{
		ID:        id,
		FirstName: fmt.Sprintf("FirstName %s", id),
		LastName:  fmt.Sprintf("LastName %s", id),
		EmailAddress: mail.Address{
			Name:    fmt.Sprintf("<Contact %s>", id),
			Address: fmt.Sprintf("%s@test.com", id),
		},
	}
}
