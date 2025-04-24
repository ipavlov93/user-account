package userservice

import (
	"context"

	"event-calendar/internal/domain"
	"event-calendar/internal/option"
	"event-calendar/internal/repository"
)

type UserAccountService struct {
	userAccountRepository repository.UserAccountRepository
}

func NewUserAccountService(
	userAccountRepository repository.UserAccountRepository,
) *UserAccountService {
	return &UserAccountService{
		userAccountRepository: userAccountRepository,
	}
}

func (s UserAccountService) ListUserAccountsByUserID(
	ctx context.Context,
	userID int64,
	options *option.TxOption,
) (
	userAccounts []domain.UserAccount,
	err error,
) {
	// inject tx into repository
	repo := option.ApplyTx(s.userAccountRepository, options)

	userAccounts, err = repo.ListUserAccountsByUserID(ctx, userID)
	if err != nil {
		//if errors.Is(err, repository.ErrNoRows) {
		//return customError with status code NotFound
		//}
		return nil, err
	}

	return userAccounts, nil
}

// CreateUserAccount supplies options as struct instance instead of functional-style WithOption() calls.
// Pass options as the nil if you want to apply default behaviour.
func (s UserAccountService) CreateUserAccount(
	ctx context.Context,
	userAccount domain.UserAccount,
	options *option.CreateUserAccountOptions,
) (int64, error) {
	repo := s.userAccountRepository

	// inject tx into repository
	if options != nil {
		repo = option.ApplyTx(s.userAccountRepository, &options.TxOption)
	}

	allowDuplicates := false
	if options != nil {
		allowDuplicates = options.AllowDuplicates
	}

	userID, err := repo.CreateUserAccount(ctx, userAccount, allowDuplicates)
	if err != nil {
		//if errors.Is(err, repository.ErrDuplicate) {
		//return customError with status code BadRequest
		//}
		return 0, err
	}

	return userID, nil
}
