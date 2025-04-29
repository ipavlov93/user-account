package userservice

import (
	"context"

	"event-calendar/internal/domain"
	"event-calendar/internal/option"
	"event-calendar/internal/repository"
)

type UserProfileService struct {
	userProfileRepository repository.UserProfileRepository
}

func NewUserProfileService(
	userProfileRepository repository.UserProfileRepository,
) *UserProfileService {
	return &UserProfileService{
		userProfileRepository: userProfileRepository,
	}
}

func (s UserProfileService) GetUserProfileByID(
	ctx context.Context,
	ID int64,
	options *option.TxOption,
) (
	userProfile domain.UserProfile,
	found bool,
	err error,
) {
	// inject tx into repository
	repo := option.ApplyTx(s.userProfileRepository, options)

	userProfile, err = repo.GetUserProfileByID(ctx, ID)
	if err != nil {
		//if errors.Is(err, repository.ErrNoRows) {
		//return customError with status code NotFound
		//}
		return domain.UserProfile{}, false, err
	}

	return userProfile, true, nil
}

func (s UserProfileService) GetUserProfileByUserID(
	ctx context.Context,
	userID int64,
	options *option.TxOption,
) (
	userProfile domain.UserProfile,
	found bool,
	err error,
) {
	// inject tx into repository
	repo := option.ApplyTx(s.userProfileRepository, options)

	userProfile, err = repo.GetUserProfileByUserID(ctx, userID)
	if err != nil {
		//if errors.Is(err, repository.ErrNoRows) {
		//return customError with status code NotFound
		//}
		return domain.UserProfile{}, false, err
	}

	return userProfile, true, nil
}

func (s UserProfileService) GetUserProfileByUUID(
	ctx context.Context,
	uuid string,
	options *option.TxOption,
) (
	userProfile domain.UserProfile,
	found bool,
	err error,
) {
	// inject tx into repository
	repo := option.ApplyTx(s.userProfileRepository, options)

	userProfile, err = repo.GetUserProfileByFirebaseUID(ctx, uuid)
	if err != nil {
		//if errors.Is(err, repository.ErrNoRows) {
		//return customError with status code NotFound
		//}
		return domain.UserProfile{}, false, err
	}

	return userProfile, true, nil
}

// CreateUserProfile supplies options as struct instance instead of functional-style WithOption() calls.
// Pass options as the nil if you want to apply default behaviour.
func (s UserProfileService) CreateUserProfile(
	ctx context.Context,
	userProfile domain.UserProfile,
	options *option.CreateOptions,
) (int64, error) {
	// inject tx into repository
	repo := option.ApplyTx(s.userProfileRepository, &options.TxOption)

	userProfileID, err := repo.CreateUserProfile(ctx, userProfile)
	if err != nil {
		//if errors.Is(err, repository.ErrDuplicate) {
		//return customError with status code BadRequest
		//}
		return 0, err
	}

	return userProfileID, nil
}
