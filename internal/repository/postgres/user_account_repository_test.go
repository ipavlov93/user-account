package postgres

import (
	"context"
	"event-calendar/internal/domain"
	"event-calendar/internal/test"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
)

func TestListUserAccountByUserID(t *testing.T) {
	// Setup mock DB
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "postgres")
	repo := NewUserAccountRepository(sqlxDB)

	// Define expected user
	expectedUserAccount1 := test.CreateTestUserAccount(1)
	expectedUserAccount2 := test.CreateTestUserAccount(2)
	expectedUserAccount1.UserID = 1
	expectedUserAccount2.UserID = 1

	expectedUserAccounts := []domain.UserAccount{
		expectedUserAccount1,
		expectedUserAccount2,
	}

	// Mock the query
	rows := sqlmock.NewRows(
		[]string{"id", "user_id", "issuer_code", "subject_uid", "email_address", "contact_name"}).
		AddRow(expectedUserAccount1.ID, expectedUserAccount1.UserID, expectedUserAccount1.IssuerCode, expectedUserAccount1.SubjectUID, expectedUserAccount1.EmailAddress, expectedUserAccount1.ContactName).
		AddRow(expectedUserAccount2.ID, expectedUserAccount2.UserID, expectedUserAccount2.IssuerCode, expectedUserAccount2.SubjectUID, expectedUserAccount2.EmailAddress, expectedUserAccount2.ContactName)

	mock.ExpectQuery(`(?i)SELECT \* FROM user_accounts WHERE user_id = \$1`).
		WithArgs(expectedUserAccount1.UserID).
		WillReturnRows(rows)

	// ACT
	ctx := context.Background()
	userAccounts, err := repo.ListUserAccountsByUserID(ctx, expectedUserAccount1.UserID)

	// Assertions
	assert.NoError(t, err)
	assert.Equal(t, expectedUserAccounts, userAccounts)

	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestCreateUserAccount(t *testing.T) {
	// Setup mock DB
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "postgres")
	repo := NewUserAccountRepository(sqlxDB)

	// Define expected user and mock response
	newUserAccount := test.CreateTestUserAccount(1)

	userID := int64(1)
	newUserAccount.UserID = userID

	// Define expected behavior for mock
	rows := sqlmock.NewRows([]string{"id"}).AddRow(1)
	mock.ExpectQuery(
		`(?i)INSERT INTO user_accounts \(user_id, issuer_code, subject_uid, email_address, contact_name\) VALUES \(\$1, \$2, \$3, \$4, \$5\) RETURNING id`).
		WithArgs(newUserAccount.UserID, newUserAccount.IssuerCode, newUserAccount.SubjectUID, newUserAccount.EmailAddress, newUserAccount.ContactName).
		WillReturnRows(rows) // Return ID = 1

	// ACT
	userAccountID, err := repo.CreateUserAccount(context.Background(), newUserAccount, false)

	// Assertions
	assert.NoError(t, err)
	assert.Equal(t, userID, userAccountID)

	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestCreateUserAccountIgnoreDuplicate(t *testing.T) {
	// Setup mock DB
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "postgres")
	repo := NewUserAccountRepository(sqlxDB)

	// Define expected user and mock response
	newUserAccount := test.CreateTestUserAccount(1)
	newUserAccount.UserID = 1

	// Define expected behavior for mock
	rows := sqlmock.NewRows([]string{"id"}).AddRow(1)
	mock.ExpectQuery(
		`(?i)INSERT INTO user_accounts \(user_id, issuer_code, subject_uid, email_address, contact_name\) VALUES \(\$1, \$2, \$3, \$4, \$5\) ON CONFLICT DO NOTHING RETURNING id`).
		WithArgs(newUserAccount.UserID, newUserAccount.IssuerCode, newUserAccount.SubjectUID, newUserAccount.EmailAddress, newUserAccount.ContactName).
		WillReturnRows(rows) // Return ID = 1

	// ACT
	userAccountID, err := repo.CreateUserAccount(context.Background(), newUserAccount, true)

	// Assertions
	assert.NoError(t, err)
	assert.Equal(t, int64(1), userAccountID)

	assert.NoError(t, mock.ExpectationsWereMet())
}
