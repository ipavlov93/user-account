package postgres

import (
	"context"
	"event-calendar/internal/domain/test"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
)

func TestGetUserProfilesCount(t *testing.T) {
	// Create a new mock DB
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	// Wrap sql.DB with sqlx.DB
	sqlxDB := sqlx.NewDb(db, "postgres")

	// Instantiate the repository
	repo := UserProfileRepositoryPostgres{dbDriver: sqlxDB}

	// Define expected behavior for mock
	rows := sqlmock.NewRows([]string{"count"}).AddRow(10)
	mock.ExpectQuery(`(?i)SELECT COUNT\(\*\) FROM user_profiles`).WillReturnRows(rows)

	// ACT
	count, err := repo.GetUserProfilesCount(context.Background())

	// Assertions
	assert.NoError(t, err)
	assert.Equal(t, int64(10), count)

	// Ensure expectations were met
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGetUserProfileByID(t *testing.T) {
	// Setup mock DB
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "postgres")
	repo := UserProfileRepositoryPostgres{dbDriver: sqlxDB}

	// Define expected userProfile
	expectedUserProfile := test.CreateTestUserProfile(1)
	// optionally set userID
	expectedUserProfile.UserID = 1

	// Mock the query
	rows := sqlmock.NewRows(
		[]string{"id", "first_name", "last_name", "user_id", "business_name", "contact_email", "organization", "avatar_file_name", "description"}).
		AddRow(expectedUserProfile.ID, expectedUserProfile.FirstName, expectedUserProfile.LastName, expectedUserProfile.UserID, expectedUserProfile.BusinessName, expectedUserProfile.ContactEmail, expectedUserProfile.Organization, expectedUserProfile.AvatarFileName, expectedUserProfile.Description)

	mock.ExpectQuery(`(?i)SELECT \* FROM user_profiles WHERE id = \$1`).
		WithArgs(expectedUserProfile.ID).
		WillReturnRows(rows)

	// ACT
	ctx := context.Background()
	userProfile, err := repo.GetUserProfileByID(ctx, expectedUserProfile.ID)

	// Assertions
	assert.NoError(t, err)
	assert.True(t, userProfile.Equals(&expectedUserProfile))
	// use Equals() to ignore CreatedAt comparison
	//assert.Equal(t, expectedUser, user)

	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGetUserProfileByUserID(t *testing.T) {
	// Setup mock DB
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "postgres")
	repo := UserProfileRepositoryPostgres{dbDriver: sqlxDB}

	// Define expected userProfile
	user := test.CreateTestUser(1)
	expectedUserProfile := test.CreateTestUserProfile(1)
	// optionally set userID
	expectedUserProfile.UserID = user.ID

	// Mock the query
	rows := sqlmock.NewRows(
		[]string{"id", "first_name", "last_name", "user_id", "business_name", "contact_email", "organization", "avatar_file_name", "description"}).
		AddRow(expectedUserProfile.ID, expectedUserProfile.FirstName, expectedUserProfile.LastName, expectedUserProfile.UserID, expectedUserProfile.BusinessName, expectedUserProfile.ContactEmail, expectedUserProfile.Organization, expectedUserProfile.AvatarFileName, expectedUserProfile.Description)

	mock.ExpectQuery(`(?i)SELECT \* FROM user LEFT JOIN user_profiles ON user.id = user_profiles.user_id WHERE user.id = \$1`).
		WithArgs(user.ID).
		WillReturnRows(rows)

	// ACT
	ctx := context.Background()
	userProfile, err := repo.GetUserProfileByUserID(ctx, user.ID)

	// Assertions
	assert.NoError(t, err)
	assert.True(t, userProfile.Equals(&expectedUserProfile))
	// use Equals() to ignore CreatedAt comparison
	//assert.Equal(t, expectedUserProfile, userProfile)

	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGetUserProfileByUUID(t *testing.T) {
	// Setup mock DB
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "postgres")
	repo := UserProfileRepositoryPostgres{dbDriver: sqlxDB}

	// Define expected userProfile
	user := test.CreateTestUser(1)
	expectedUserProfile := test.CreateTestUserProfile(1)
	// optionally set userID
	expectedUserProfile.UserID = user.ID

	// Mock the query
	rows := sqlmock.NewRows(
		[]string{"id", "first_name", "last_name", "user_id", "business_name", "contact_email", "organization", "avatar_file_name", "description"}).
		AddRow(expectedUserProfile.ID, expectedUserProfile.FirstName, expectedUserProfile.LastName, expectedUserProfile.UserID, expectedUserProfile.BusinessName, expectedUserProfile.ContactEmail, expectedUserProfile.Organization, expectedUserProfile.AvatarFileName, expectedUserProfile.Description)

	mock.ExpectQuery(`(?i)SELECT \* FROM user LEFT JOIN user_profiles ON user.id = user_profiles.user_id WHERE user.firebase_uid = \$1`).
		WithArgs(user.FirebaseUUID).
		WillReturnRows(rows)

	// ACT
	ctx := context.Background()
	userProfile, err := repo.GetUserProfileByFirebaseUUID(ctx, user.FirebaseUUID)

	// Assertions
	assert.NoError(t, err)
	assert.True(t, userProfile.Equals(&expectedUserProfile))
	// use Equals() to ignore CreatedAt comparison
	//assert.Equal(t, expectedUserProfile, userProfile)

	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestCreateUserProfile(t *testing.T) {
	// Setup mock DB
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "postgres")
	repo := UserProfileRepositoryPostgres{dbDriver: sqlxDB}

	// Define expected user and mock response
	newUser := test.CreateTestUserProfile(1)

	// Define expected behavior for mock
	rows := sqlmock.NewRows([]string{"id"}).AddRow(1)
	mock.ExpectQuery(
		`(?i)INSERT INTO user_profiles \(first_name, last_name, user_id, business_name, contact_email, organization, avatar_file_name, description\) VALUES \(\$1, \$2, \$3, \$4, \$5, \$6, \$7, \$8\) RETURNING id`).
		WithArgs(newUser.FirstName, newUser.LastName, newUser.UserID, newUser.BusinessName, newUser.ContactEmail, newUser.Organization, newUser.AvatarFileName, newUser.Description).
		WillReturnRows(rows) // Return ID = 1

	// ACT
	userID, err := repo.CreateUserProfile(context.Background(), newUser)

	// Assertions
	assert.NoError(t, err)
	assert.Equal(t, int64(1), userID)

	assert.NoError(t, mock.ExpectationsWereMet())
}
