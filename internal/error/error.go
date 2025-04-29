package error

import (
	"context"
	"fmt"
	"maps"
	"net/http"
)

const DefaultLocale = "en_GB"

type TError string

func (t TError) String() string {
	return string(t)
}

var langFn func(ctx context.Context, key string) (locale string, msg string)

// SetLanguageFunc
// determine function who returns locale & translated message by key
// locale must extract from metadata of context
// if no present returns EMPTY locale. will be set default message without locale
func SetLanguageFunc(fn func(ctx context.Context, key string) (string, string)) {
	langFn = fn
}

var uniqueErrors = make(map[TError]string)

// ListErrors
// returns list of all error messages with their keys
func ListErrors() map[TError]string {
	return maps.Clone(uniqueErrors)
}

type Error interface {
	WithCode(code int) Error
	WithInfo(reason string) Error

	Error() string
	ToHTTP(ctx context.Context) (statusCode int, localizedMessage string)
}

func New(key TError, message string) Error {
	if _, ok := uniqueErrors[key]; ok {
		panic(fmt.Sprintf("duplicate error %v", key))
	}
	uniqueErrors[key] = message

	return &er{
		code:    http.StatusInternalServerError,
		message: message,
		errorInfo: &Info{
			Domain: string(key),
			Reason: message,
		},
	}
}

type er struct {
	langFn    func(ctx context.Context, key string) error
	code      int
	message   string
	errorInfo *Info
}

type Info struct {
	Domain   string            // e.g., "user_not_found"
	Reason   string            // optional custom description
	Metadata map[string]string // optional
}

func (e *er) WithCode(code int) Error {
	e.code = code
	return e
}

func (e *er) WithInfo(reason string) Error {
	e.errorInfo.Reason = reason
	return e
}

func (e *er) Error() string {
	return fmt.Sprintf("%s: %s", e.errorInfo.Domain, e.errorInfo.Reason)
}

func (e *er) ToHTTP(ctx context.Context) (int, string) {
	if langFn == nil {
		return e.code, e.message
	}

	locale, localizedMsg := langFn(ctx, e.errorInfo.Domain)
	if locale == "" {
		return e.code, e.message
	}

	if len(localizedMsg) == 0 {
		locale = DefaultLocale
		localizedMsg = e.message
	}

	return e.code, localizedMsg
}
