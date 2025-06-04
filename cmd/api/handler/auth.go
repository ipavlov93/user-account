package handler

import (
	"net/http"

	"event-calendar/facade"
	"event-calendar/internal/logger"

	"go.uber.org/zap"
)

type AuthController struct {
	service facade.AuthService
	logger  logger.Logger
}

// NewAuthController set default logger. Use WithLogger() to set custom logger.
func NewAuthController(service facade.AuthService) *AuthController {
	return &AuthController{
		service: service,
	}
}

// WithLogger sets the logger and returns the *AuthMiddleware
func (c *AuthController) WithLogger(logger logger.Logger) *AuthController {
	if logger != nil {
		c.logger = logger
	}
	return c
}

func (c *AuthController) LoginHandler(rw http.ResponseWriter, r *http.Request) {
	claims, ok := getTokenClaims(r.Context())
	if claims == nil || !ok {
		c.logger.Error("request context contains not valid claims. %T are expected")
		http.Error(rw,
			"empty claims",
			http.StatusInternalServerError)
		return
	}

	err := c.service.Login(r.Context(), claims)
	if err != nil {
		c.logger.Error("Login()", zap.Error(err))
		http.Error(rw,
			http.StatusText(http.StatusInternalServerError),
			http.StatusInternalServerError)
	}

	rw.WriteHeader(http.StatusOK)
}
