package auth_service_rest

import (
	"time"

	models_rest "github.com/L1z1ng3r-sswe/telegram_clone/app/internal/rest/domain/models"
	tokens_rest "github.com/L1z1ng3r-sswe/telegram_clone/app/internal/rest/utils/jwt"
	password_hash_rest "github.com/L1z1ng3r-sswe/telegram_clone/app/internal/rest/utils/password"
	validation_rest "github.com/L1z1ng3r-sswe/telegram_clone/app/internal/rest/utils/validation"
)

func (s *service) SignUp(user models_rest.UserSignUp, accessTokenExp time.Duration, refreshTokenExp time.Duration, secretKey string) (models_rest.Tokens, models_rest.UserSignUp, *models_rest.Response) {
	// validation
	err := validation_rest.ValidationSignUp(user.Email, user.Password)
	if err != nil {
		return models_rest.Tokens{}, models_rest.UserSignUp{}, err
	}

	// hash password
	user.Password, err = password_hash_rest.PasswordHasher(user.Password)
	if err != nil {
		return models_rest.Tokens{}, models_rest.UserSignUp{}, err
	}

	// save new user in db
	userDB, err := s.repo.SignUp(user)
	if err != nil {
		return models_rest.Tokens{}, models_rest.UserSignUp{}, err
	}

	// create access-token
	accessToken, err := tokens_rest.CreateAccessToken(userDB.Id, accessTokenExp, secretKey)
	if err != nil {
		return models_rest.Tokens{}, models_rest.UserSignUp{}, err
	}

	// create refresh-token
	refreshToken, err := tokens_rest.CreateRefreshToken(userDB.Id, refreshTokenExp, secretKey)
	if err != nil {
		return models_rest.Tokens{}, models_rest.UserSignUp{}, err
	}

	return models_rest.Tokens{AccessToken: accessToken, RefreshToken: refreshToken}, userDB, nil
}
