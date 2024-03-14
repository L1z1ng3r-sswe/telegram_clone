package auth_service_rest

import (
	"time"

	models_rest "github.com/L1z1ng3r-sswe/telegram_clone/app/internal/rest/domain/models"
	tokens_rest "github.com/L1z1ng3r-sswe/telegram_clone/app/internal/rest/utils/jwt"
	password_hash_rest "github.com/L1z1ng3r-sswe/telegram_clone/app/internal/rest/utils/password"
	validation_rest "github.com/L1z1ng3r-sswe/telegram_clone/app/internal/rest/utils/validation"
)

func (s *service) SignIn(user models_rest.UserSignIn, accessTokenExp time.Duration, refreshTokenExp time.Duration, secretKey string) (models_rest.Tokens, models_rest.UserDB, *models_rest.Response) {
	// validation
	err := validation_rest.ValidationSignIn(user.Email, user.Password)
	if err != nil {
		return models_rest.Tokens{}, models_rest.UserDB{}, err
	}

	// ger the user from db from db by email
	userDB, err := s.repo.SignIn(user)
	if err != nil {
		return models_rest.Tokens{}, models_rest.UserDB{}, err
	}

	// compare passwords
	if err := password_hash_rest.ComparePasswords(userDB.Password, user.Password); err != nil {
		return models_rest.Tokens{}, models_rest.UserDB{}, err
	}

	// create access-token
	accessToken, err := tokens_rest.CreateAccessToken(userDB.Id, accessTokenExp, secretKey)
	if err != nil {
		return models_rest.Tokens{}, models_rest.UserDB{}, err
	}

	// create refresh-token
	refreshToken, err := tokens_rest.CreateRefreshToken(userDB.Id, refreshTokenExp, secretKey)
	if err != nil {
		return models_rest.Tokens{}, models_rest.UserDB{}, err
	}

	return models_rest.Tokens{AccessToken: accessToken, RefreshToken: refreshToken}, userDB, nil

}
