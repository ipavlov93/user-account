package userservice

import (
	"context"

	"event-calendar/internal/domain"
	repository "event-calendar/internal/repository/postgres"

	"github.com/jmoiron/sqlx"
)

type UserAccountService struct {
	dbDriver sqlx.ExtContext
}

func NewUserAccountService(dbDriver sqlx.ExtContext) *UserAccountService {
	return &UserAccountService{
		dbDriver: dbDriver,
	}
}

func (s UserAccountService) ListUserAccountsByUserID(ctx context.Context, userID int64) (userAccountsList []domain.UserAccount, err error) {
	repo := repository.NewUserAccountRepository(s.dbDriver)
	userAccountsList, err = repo.ListUserAccountsByUserID(ctx, userID)
	if err != nil {
		//if errors.Is(err, repository.ErrNoRows) {
		//return customError with status code NotFound
		//}
		return userAccountsList, err
	}

	return userAccountsList, nil
}

// todo: refactor ignoreDuplicate required parameter to WithOption()
func (s UserAccountService) CreateUserAccount(ctx context.Context, userAccount domain.UserAccount, ignoreDuplicate bool) (int64, error) {
	repo := repository.NewUserAccountRepository(s.dbDriver)
	userID, err := repo.CreateUserAccount(ctx, userAccount, ignoreDuplicate)
	if err != nil {
		//if errors.Is(err, repository.ErrDuplicate) {
		//return customError with status code BadRequest
		//}
		return 0, err
	}

	return userID, nil
}
