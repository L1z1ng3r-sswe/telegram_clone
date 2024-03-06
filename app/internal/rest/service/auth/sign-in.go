package auth_service_rest

import (
	"net/http"
	"time"

	models_rest "github.com/L1z1ng3r-sswe/telegram_clone/app/internal/rest/domain/models"
	tokens_rest "github.com/L1z1ng3r-sswe/telegram_clone/app/internal/rest/utils/jwt"
	password_hash_rest "github.com/L1z1ng3r-sswe/telegram_clone/app/internal/rest/utils/password"
	validation_rest "github.com/L1z1ng3r-sswe/telegram_clone/app/internal/rest/utils/validation"
)

func (s *service) SignIn(user models_rest.UserSignIn, accessTokenExp time.Duration, refreshTokenExp time.Duration, secretKey string) (models_rest.Tokens, models_rest.UserDB, error, string, string, int, string) {
	// validation
	err, errKey, errMsg, code, fileInfo := validation_rest.ValidationSignIn(user.Email, user.Password)
	if err != nil {
		return models_rest.Tokens{}, models_rest.UserDB{}, err, errKey, errMsg, code, fileInfo
	}

	// ger the user from db from db by email
	userDB, err, errKey, errMsg, code, fileInfo := s.repo.SignIn(user)
	if err != nil {
		return models_rest.Tokens{}, models_rest.UserDB{}, err, errKey, errMsg, code, fileInfo
	}

	// compare passwords
	if err, errKey, errMsg, code, fileInfo := password_hash_rest.ComparePasswords(userDB.Password, user.Password); err != nil {
		return models_rest.Tokens{}, models_rest.UserDB{}, err, errKey, errMsg, code, fileInfo
	}

	// create access-token
	accessToken, err, errKey, errMsg, code, fileInfo := tokens_rest.CreateAccessToken(userDB.Id, accessTokenExp, secretKey)
	if err != nil {
		return models_rest.Tokens{}, models_rest.UserDB{}, err, errKey, errMsg, code, fileInfo
	}

	// create refresh-token
	refreshToken, err, errKey, errMsg, code, fileInfo := tokens_rest.CreateRefreshToken(userDB.Id, refreshTokenExp, secretKey)
	if err != nil {
		return models_rest.Tokens{}, models_rest.UserDB{}, err, errKey, errMsg, code, fileInfo
	}

	return models_rest.Tokens{AccessToken: accessToken, RefreshToken: refreshToken}, userDB, nil, "", "", http.StatusOK, ""

}
