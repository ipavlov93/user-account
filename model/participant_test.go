package model_test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/mail"
	"secret-santa/model"
	"testing"
)

func TestNewParticipant(t *testing.T) {
	// ARRANGE
	participant1 := createDummyParticipant("1")

	// happy flow
	t.Run("should create participant and init giving fields (no user ID)", func(t *testing.T) {
		// ACT
		participant2 := model.NewParticipant(
			participant1.FirstName,
			participant1.LastName,
			participant1.ContactEmail.Address,
		)

		// ASSERT
		assert.Equal(t, participant2.LastName, participant1.LastName)
		assert.Equal(t, participant2.LastName, participant1.LastName)
		assert.Equal(t, participant2.ContactEmail.Address, participant1.ContactEmail.Address)

		//if !participant2.Equals(participant1) {
		//	t.Errorf("got %s but want %s", participant2, participant1)
		//}
	})
	// happy flow
	t.Run("should create participant and init giving fields (with user ID)", func(t *testing.T) {
		// ARRANGE
		user1 := createDummyUser("1")

		// ACT
		participant2 := model.NewParticipant(
			participant1.FirstName,
			participant1.LastName,
			participant1.ContactEmail.Address,
			user1.ID,
		)

		// ASSERT
		assert.Equal(t, participant2.User.ID, user1.ID)
		assert.Equal(t, participant2.FirstName, participant1.FirstName)
		assert.Equal(t, participant2.LastName, participant1.LastName)
		assert.Equal(t, participant2.ContactEmail.Address, participant1.ContactEmail.Address)

		//if !participant2.Equals(participant1) {
		//	t.Errorf("got %s but want %s", participant2, participant1)
		//}
	})
}

// fill all the string fields with given value
func createDummyParticipant(id string) *model.Participant {
	return &model.Participant{
		ID:        id,
		FirstName: fmt.Sprintf("FirstName %s", id),
		LastName:  fmt.Sprintf("LastName %s", id),
		ContactEmail: mail.Address{
			Name:    fmt.Sprintf("<Contact %s>", id),
			Address: fmt.Sprintf("%s@test.com", id),
		},
	}
}
