package test

import (
	"testing"
	"user-account/internal/domain"

	"github.com/stretchr/testify/assert"
)

func TestNewProfile(t *testing.T) {
	// ARRANGE
	expectedProfile := CreateTestUserProfile(1)
	// MANDATORY STEP
	//set expected profile ID = 0
	expectedProfile.ID = 0

	expectedUserID := int64(1)
	expectedProfile.UserID = expectedUserID

	// happy flow
	t.Run("should create profile and init giving fields (with user ID)", func(t *testing.T) {
		// ACT
		testProfile := domain.NewUserProfile(
			expectedUserID,
			expectedProfile.FirstName,
			expectedProfile.LastName,
			expectedProfile.BusinessName,
			expectedProfile.ContactEmail,
			expectedProfile.Organization,
			expectedProfile.Description,
			expectedProfile.AvatarFileName,
		)

		// ASSERT
		assert.True(t, testProfile.Equals(&expectedProfile))
		// use Equals() to ignore CreatedAt comparison
		//assert.Equal(t, expectedProfile, testProfile)
	})
}
