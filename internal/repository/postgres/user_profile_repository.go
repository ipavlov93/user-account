// Package postgres is part of repository.
// Package provides implementations of persistence layer interfaces.
// It includes interactions with PostgreSQL for user profile storage, modification and retrieval.
package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"user-account/internal/domain"
	"user-account/internal/dto/dmodel"
	errs "user-account/internal/error"
	"user-account/internal/logger"
	mapper "user-account/internal/mapper/user/dmodel"
	"user-account/internal/repository"

	"github.com/jmoiron/sqlx"
)

// UserProfileRepositoryPostgres implements UserProfileRepository using sqlx.
// It provides methods to manage user profiles in PostgreSQL.
type UserProfileRepositoryPostgres struct {
	// dbDriver abstraction
	dbDriver sqlx.ExtContext
	logger   logger.Logger
}

func NewUserProfileRepository(dbDriver sqlx.ExtContext) *UserProfileRepositoryPostgres {
	return &UserProfileRepositoryPostgres{
		dbDriver: dbDriver,
	}
}

// WithLogger sets the logger and returns the *UserProfileRepositoryPostgres
func (repo *UserProfileRepositoryPostgres) WithLogger(logger logger.Logger) *UserProfileRepositoryPostgres {
	if logger != nil {
		repo.logger = logger
	}
	return repo
}

// WithTx returns new copy of UserAccountRepository with new dbDriver.
func (repo *UserProfileRepositoryPostgres) WithTx(tx *sqlx.Tx) repository.UserProfileRepository {
	return &UserProfileRepositoryPostgres{
		dbDriver: tx,
		logger:   repo.logger,
	}
}

// GetUserProfilesCount retrieves a user profile total count.
func (repo *UserProfileRepositoryPostgres) GetUserProfilesCount(ctx context.Context) (int64, error) {
	var count int64
	err := sqlx.GetContext(ctx, repo.dbDriver, &count,
		`SELECT count(*) FROM user_profiles`)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return 0, nil
		}
		return 0, errs.ErrDB.WithInfo(fmt.Sprintf("repository.GetUserProfilesCount: %v", err))
	}

	return count, nil
}

// GetUserProfileByID retrieves a user profile by its unique ID.
// Returns errs.ErrUserProfileNotFound if no matching record exists.
func (repo *UserProfileRepositoryPostgres) GetUserProfileByID(ctx context.Context, ID int64) (obj domain.UserProfile, err error) {
	var userProfileDto dmodel.UserProfile
	err = sqlx.GetContext(ctx, repo.dbDriver, &userProfileDto,
		`SELECT * FROM user_profiles
				WHERE id = $1`, ID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return domain.UserProfile{}, errs.ErrUserProfileNotFound.WithInfo(fmt.Sprintf("repository.GetUserProfileByID: user profile not found for ID=%d", ID))
		}
		return domain.UserProfile{}, errs.ErrDB.WithInfo(fmt.Sprintf("repository.GetUserProfileByID: %v", err))
	}
	return mapper.ProfileDtoToProfile(userProfileDto), nil
}

// GetUserProfileByUserID retrieves a user profile by user ID.
// Returns errs.ErrUserProfileNotFound if no matching record exists.
func (repo *UserProfileRepositoryPostgres) GetUserProfileByUserID(ctx context.Context, userID int64) (obj domain.UserProfile, err error) {
	var userProfileDto dmodel.UserProfile
	err = sqlx.GetContext(ctx, repo.dbDriver, &userProfileDto,
		`SELECT * FROM user
    			LEFT JOIN user_profiles ON user.id = user_profiles.user_id
				WHERE user.id = $1`, userID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return domain.UserProfile{}, errs.ErrUserProfileNotFound.WithInfo(
				fmt.Sprintf("repository.GetUserProfileByUserID: user profile not found for userID=%d", userID))
		}
		return domain.UserProfile{}, errs.ErrDB.WithInfo(fmt.Sprintf("repository.GetUserProfileByUserID: %v", err))
	}
	return mapper.ProfileDtoToProfile(userProfileDto), nil
}

// GetUserProfileByFirebaseUUID retrieves a user profile by user UID.
// Returns errs.ErrUserProfileNotFound if no matching record exists.
func (repo *UserProfileRepositoryPostgres) GetUserProfileByFirebaseUUID(ctx context.Context, firebaseUUID string) (obj domain.UserProfile, err error) {
	var userProfileDto dmodel.UserProfile
	err = sqlx.GetContext(ctx, repo.dbDriver, &userProfileDto,
		`SELECT * FROM user
    			LEFT JOIN user_profiles ON user.id = user_profiles.user_id
				WHERE user.firebase_uuid = $1`, firebaseUUID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return domain.UserProfile{}, errs.ErrUserProfileNotFound.WithInfo(
				fmt.Sprintf("repository.GetUserProfileByFirebaseUUID: user profile not found for UUID=%s", firebaseUUID))
		}
		return domain.UserProfile{}, errs.ErrDB.WithInfo(fmt.Sprintf("repository.GetUserProfileByFirebaseUUID: %v", err))
	}
	return mapper.ProfileDtoToProfile(userProfileDto), nil
}

// CreateUserProfile inserts a new user profile into the database.
// Returns an error if the operation fails or the user ID already exists.
// IMPORTANT: ignore given CreatedAt value.
func (repo *UserProfileRepositoryPostgres) CreateUserProfile(ctx context.Context, user domain.UserProfile) (userID int64, err error) {
	err = repo.dbDriver.QueryRowxContext(
		ctx,
		`INSERT INTO user_profiles (first_name, last_name, user_id, business_name, contact_email, organization, avatar_file_name, description) VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id`,
		user.FirstName, user.LastName, user.UserID, user.BusinessName, user.ContactEmail, user.Organization, user.AvatarFileName, user.Description,
	).Scan(&userID)
	if err != nil {
		errorInfo := fmt.Sprintf("repository.CreateUserProfile: %v", err)

		if len(err.Error()) > 50 && err.Error()[:50] == pqDuplicateErr {
			return 0, errs.ErrUserProfileExists.WithInfo(errorInfo)
		}
		return 0, errs.ErrDB.WithInfo(errorInfo)
	}
	return userID, nil
}
