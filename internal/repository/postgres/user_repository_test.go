package postgres

import (
	"context"
	"event-calendar/internal/test"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
)

func TestGetUsersCount(t *testing.T) {
	// Create a new mock DB
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	// Wrap sql.DB with sqlx.DB
	sqlxDB := sqlx.NewDb(db, "postgres")

	// Instantiate the repository
	repo := UserRepository{db: sqlxDB}

	// Define expected behavior for mock
	rows := sqlmock.NewRows([]string{"count"}).AddRow(10)
	mock.ExpectQuery(`(?i)SELECT COUNT\(\*\) FROM users`).WillReturnRows(rows)

	// ACT
	count, err := repo.GetUsersCount(context.Background())

	// Assertions
	assert.NoError(t, err)
	assert.Equal(t, int64(10), count)

	// Ensure expectations were met
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGetUserByID(t *testing.T) {
	// Setup mock DB
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "postgres")
	repo := UserRepository{db: sqlxDB}

	// Define expected user
	expectedUser := test.CreateTestUser(1)

	// Mock the query
	rows := sqlmock.NewRows(
		[]string{"id", "firebase_uid", "description"}).
		AddRow(expectedUser.ID, expectedUser.FirebaseUID, expectedUser.Description)

	mock.ExpectQuery(`(?i)SELECT \* FROM users WHERE id = \$1`).
		WithArgs(expectedUser.ID).
		WillReturnRows(rows)

	// ACT
	ctx := context.Background()
	user, err := repo.GetUserByID(ctx, expectedUser.ID)

	// Assertions
	assert.NoError(t, err)
	assert.True(t, user.Equals(&expectedUser))
	// use Equals() to ignore CreatedAt comparison
	//assert.Equal(t, expectedUser, user)

	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGetUserByUUID(t *testing.T) {
	// Setup mock DB
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "postgres")
	repo := UserRepository{db: sqlxDB}

	// Define expected user
	expectedUser := test.CreateTestUser(1)

	// Mock the query
	rows := sqlmock.NewRows(
		[]string{"id", "firebase_uid", "description"}).
		AddRow(expectedUser.ID, expectedUser.FirebaseUID, expectedUser.Description)

	mock.ExpectQuery(`(?i)SELECT \* FROM users WHERE firebase_uid = \$1`).
		WithArgs(expectedUser.FirebaseUID).
		WillReturnRows(rows)

	// ACT
	ctx := context.Background()
	user, err := repo.GetUserByFirebaseUID(ctx, expectedUser.FirebaseUID)

	// Assertions
	assert.NoError(t, err)
	assert.True(t, user.Equals(&expectedUser))
	// use Equals() to ignore CreatedAt comparison
	//assert.Equal(t, expectedUser, user)

	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestCreateUser(t *testing.T) {
	// Setup mock DB
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "postgres")
	repo := UserRepository{db: sqlxDB}

	// Define expected user and mock response
	newUser := test.CreateTestUser(1)

	// Define expected behavior for mock
	rows := sqlmock.NewRows([]string{"id"}).AddRow(1)
	mock.ExpectQuery(
		`(?i)INSERT INTO users \(firebase_uid, description\) VALUES \(\$1, \$2\) RETURNING id`).
		WithArgs(newUser.FirebaseUID, newUser.Description).
		WillReturnRows(rows) // Return ID = 1

	// ACT
	userID, err := repo.CreateUser(context.Background(), newUser)

	// Assertions
	assert.NoError(t, err)
	assert.Equal(t, int64(1), userID)

	assert.NoError(t, mock.ExpectationsWereMet())
}
