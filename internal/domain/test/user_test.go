package test

import (
	"testing"

	"user-account/internal/domain"
	"user-account/internal/domain/role"

	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T) {
	// ARRANGE
	roles := []role.Role{
		role.User,
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
			expectedUser.FirebaseUUID,
			expectedUser.Description,
			roles...,
		)

		// ASSERT
		assert.True(t, testUser.Equals(&expectedUser))
		// use Equals() to ignore CreatedAt comparison
		//assert.Equal(t, expectedUser, testUser)
	})
}
