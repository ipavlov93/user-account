package test

import (
	"event-calendar/internal/domain/claims"
	"testing"

	"event-calendar/internal/domain"

	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T) {
	// ARRANGE
	roles := []claims.Role{
		claims.RoleUser,
	}
	expectedUser := CreateTestUser(1)
	expectedUser.Roles = roles

	// MANDATORY STEP
	//set expected userID = 0
	expectedUser.ID = 0

	// happy flow
	t.Run("should create user and init giving fields", func(t *testing.T) {
		// ACT
		testUser := domain.NewUser(
			expectedUser.FirebaseUID,
			expectedUser.Description,
			roles...,
		)

		// ASSERT
		assert.True(t, testUser.Equals(&expectedUser))
		// use Equals() to ignore CreatedAt comparison
		//assert.Equal(t, expectedUser, testUser)
	})
}
