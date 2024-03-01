package chat_service_rest

import models_rest "github.com/L1z1ng3r-sswe/telegram_clone/app/internal/rest/domain/models"

type repo interface {
	SessionCreation(user models_rest.UserSignIn) (models_rest.UserDB, error, string, string, int, string)
}

type service struct {
	repo repo
}

func New(repo repo) *service {
	return &service{
		repo: repo,
	}
}
