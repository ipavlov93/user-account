package handler

import (
	"log"
	"net/http"
	"os"

	"event-calendar/facade"
)

type AuthController struct {
	service facade.AuthService
	logger  *log.Logger
}

// NewAuthController set default logger. Use WithOption() to set custom logger.
func NewAuthController(service facade.AuthService) AuthController {
	return AuthController{
		service: service,
		logger:  log.New(os.Stdout, loggerPrefix, log.LstdFlags|log.Lshortfile),
	}
}

func (c AuthController) WithOption(logger *log.Logger) {
	if logger != nil {
		c.logger = logger
	}
}

func (c AuthController) LoginHandler(rw http.ResponseWriter, r *http.Request) {
	claims, ok := getTokenClaims(r.Context())
	if claims == nil || !ok {
		http.Error(rw,
			"empty claims",
			http.StatusInternalServerError)
		return
	}

	err := c.service.Login(r.Context(), claims)
	if err != nil {
		c.logger.Printf("Login(): %s", err)
		http.Error(rw,
			http.StatusText(http.StatusInternalServerError),
			http.StatusInternalServerError)
	}

	rw.WriteHeader(http.StatusOK)
}
