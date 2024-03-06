package auth_rest

import (
	"time"

	logger "github.com/L1z1ng3r-sswe/telegram_clone/app/internal/lib/zerolog"
	models_rest "github.com/L1z1ng3r-sswe/telegram_clone/app/internal/rest/domain/models"
)

type service interface {
	SignUp(models_rest.UserSignUp, time.Duration, time.Duration, string) (models_rest.Tokens, models_rest.UserSignUp, error, string, string, int, string)
	SignIn(models_rest.UserSignIn, time.Duration, time.Duration, string) (models_rest.Tokens, models_rest.UserDB, error, string, string, int, string)
}

type server struct {
	service         service
	log             *logger.Logger
	accessTokenExp  time.Duration
	refreshTokenExp time.Duration
	secretKey       string
}

func New(log *logger.Logger, service service, accessTokenExp time.Duration, refreshTokenExp time.Duration, secretKey string) *server {
	return &server{log: log, service: service, accessTokenExp: accessTokenExp, refreshTokenExp: refreshTokenExp, secretKey: secretKey}
}
