package auth_service_rest

import (
	"net/http"
	"time"

	models_rest "github.com/L1z1ng3r-sswe/telegram_clone/app/internal/rest/domain/models"
	tokens_rest "github.com/L1z1ng3r-sswe/telegram_clone/app/internal/rest/utils/jwt"
	password_hash_rest "github.com/L1z1ng3r-sswe/telegram_clone/app/internal/rest/utils/password"
	validation_rest "github.com/L1z1ng3r-sswe/telegram_clone/app/internal/rest/utils/validation"
)

func (s *Service) SignUp(user models_rest.UserSignUp, accessTokenExp time.Duration, refreshTokenExp time.Duration, secretKey string) (models_rest.Tokens, models_rest.UserSignUp, error, string, string, int, string) {
	// validation
	err, errKey, errMsg, code, fileInfo := validation_rest.ValidationSignUp(user.Email, user.Password)
	if err != nil {
		return models_rest.Tokens{}, models_rest.UserSignUp{}, err, errKey, errMsg, code, fileInfo
	}

	// hash password
	user.Password, err, errKey, errMsg, code, fileInfo = password_hash_rest.PasswordHasher(user.Password)
	if err != nil {
		return models_rest.Tokens{}, models_rest.UserSignUp{}, err, errKey, errMsg, code, fileInfo
	}

	// save new user in db
	userDB, err, errKey, errMsg, code, fileInfo := s.repo.SignUp(user)
	if err != nil {
		return models_rest.Tokens{}, models_rest.UserSignUp{}, err, errKey, errMsg, code, fileInfo
	}

	// create access-token
	accessToken, err, errKey, errMsg, code, fileInfo := tokens_rest.CreateAccessToken(userDB.Id, accessTokenExp, secretKey)
	if err != nil {
		return models_rest.Tokens{}, models_rest.UserSignUp{}, err, errKey, errMsg, code, fileInfo
	}

	// create refresh-token
	refreshToken, err, errKey, errMsg, code, fileInfo := tokens_rest.CreateRefreshToken(userDB.Id, refreshTokenExp, secretKey)
	if err != nil {
		return models_rest.Tokens{}, models_rest.UserSignUp{}, err, errKey, errMsg, code, fileInfo
	}

	return models_rest.Tokens{AccessToken: accessToken, RefreshToken: refreshToken}, userDB, nil, "", "", http.StatusOK, ""
}
