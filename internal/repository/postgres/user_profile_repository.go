package postgres

import (
	"context"
	"database/sql"
	"errors"

	"event-calendar/internal/domain"
	"event-calendar/internal/dto/dmodel"
	mapper "event-calendar/internal/mapper/user/dmodel"
	"event-calendar/internal/repository"

	"github.com/jmoiron/sqlx"
)

type UserProfileRepository struct {
	// db adapter abstraction
	db sqlx.ExtContext
}

func NewUserProfileRepository(dbAdapter sqlx.ExtContext) UserProfileRepository {
	return UserProfileRepository{
		db: dbAdapter,
	}
}

func (repo UserProfileRepository) GetUserProfilesCount(ctx context.Context) (int64, error) {
	var count int64
	err := sqlx.GetContext(ctx, repo.db, &count,
		`SELECT count(*) FROM user_profiles`)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return 0, repository.ErrNoRows
		}
		return 0, err
	}

	return count, nil
}

func (repo UserProfileRepository) GetUserProfileByID(ctx context.Context, id int64) (user domain.UserProfile, err error) {
	var userProfileDto dmodel.UserProfile
	err = sqlx.GetContext(ctx, repo.db, &userProfileDto,
		`SELECT * FROM user_profiles
				WHERE id = $1`, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return user, repository.ErrNoRows
		}
		return user, err
	}
	return mapper.ProfileDtoToProfile(userProfileDto), nil
}

func (repo UserProfileRepository) GetUserProfileByUserID(ctx context.Context, userID int64) (user domain.UserProfile, err error) {
	var userProfileDto dmodel.UserProfile
	err = sqlx.GetContext(ctx, repo.db, &userProfileDto,
		`SELECT * FROM user
    			LEFT JOIN user_profiles ON user.id = user_profiles.user_id
				WHERE user.id = $1`, userID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return user, repository.ErrNoRows
		}
		return user, err
	}
	return mapper.ProfileDtoToProfile(userProfileDto), nil
}

func (repo UserProfileRepository) GetUserProfileByFirebaseUID(ctx context.Context, firebaseUID string) (user domain.UserProfile, err error) {
	var userProfileDto dmodel.UserProfile
	err = sqlx.GetContext(ctx, repo.db, &userProfileDto,
		`SELECT * FROM user
    			LEFT JOIN user_profiles ON user.id = user_profiles.user_id
				WHERE user.firebase_uid = $1`, firebaseUID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return user, repository.ErrNoRows
		}
		return user, err
	}
	return mapper.ProfileDtoToProfile(userProfileDto), nil
}

// CreateUserProfile
// IMPORTANT: ignore given CreatedAt value.
func (repo UserProfileRepository) CreateUserProfile(ctx context.Context, user domain.UserProfile) (userID int64, err error) {
	err = repo.db.QueryRowxContext(
		ctx,
		`INSERT INTO user_profiles (first_name, last_name, user_id, business_name, contact_email, organization, avatar_file_name, description) VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id`,
		user.FirstName, user.LastName, user.UserID, user.BusinessName, user.ContactEmail, user.Organization, user.AvatarFileName, user.Description,
	).Scan(&userID)
	if err != nil {
		if len(err.Error()) > 50 {
			if err.Error()[:50] == pqDuplicateErr {
				return 0, repository.ErrDuplicate
			}
		}
		return 0, err
	}
	return userID, nil
}
