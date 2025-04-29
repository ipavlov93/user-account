package error

import (
	"net/http"
)

var (
	ErrDB           = New("USER_ACCOUNT_SV_DB_ERROR", "database error")
	ErrDBNoRows     = New("USER_ACCOUNT_SV_DB_NO_ROWS", "database no rows found")
	ErrDBConstraint = New("USER_ACCOUNT_SV_DB_CONSTRAINT", "duplicate record").WithCode(http.StatusConflict)

	ErrRequestRequired = New("USER_ACCOUNT_SV_REQUEST_REQUIRED", "request required").WithCode(http.StatusBadRequest)
	ErrInvalidArgument = New("USER_ACCOUNT_SV_BAD_REQUEST", "bad request argument(s)").WithCode(http.StatusBadRequest)

	ErrNotFound            = New("USER_ACCOUNT_SV_NOT_FOUND", "not found").WithCode(http.StatusNotFound)
	ErrUserNotFound        = New("USER_ACCOUNT_SV_USER_NOT_FOUND", "user not found").WithCode(http.StatusNotFound)
	ErrUserAccountNotFound = New("USER_ACCOUNT_SV_USER_ACCOUNT_NOT_FOUND", "user account not found").WithCode(http.StatusNotFound)
	ErrUserProfileNotFound = New("USER_ACCOUNT_SV_USER_PROFILE_NOT_FOUND", "user profile not found").WithCode(http.StatusNotFound)

	ErrUserExists = New("USER_ACCOUNT_SV_USER_EXISTS", "user already exists").WithCode(http.StatusConflict)
)
