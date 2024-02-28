package auth_service

import (
	"time"

	"github.com/L1z1ng3r-sswe/telegram_clone/app/internal/domain/models"
	password_hash "github.com/L1z1ng3r-sswe/telegram_clone/app/internal/lib/bcrypt"
	tokens "github.com/L1z1ng3r-sswe/telegram_clone/app/internal/lib/jwt"
	"github.com/L1z1ng3r-sswe/telegram_clone/app/internal/lib/validation"
	"google.golang.org/grpc/codes"
)

func (s *Service) SignIn(user models.UserSignIn, accessTokenExp time.Duration, refreshTokenExp time.Duration, secretKey string) (models.Tokens, models.UserSignIn, error, string, string, codes.Code, string) {
	// validation
	err, errKey, errMsg, code, fileInfo := validation.ValidationSignIn(user.Email, user.Password)
	if err != nil {
		return models.Tokens{}, models.UserSignIn{}, err, errKey, errMsg, code, fileInfo
	}

	// ger the user from db from db by email
	userDB, err, errKey, errMsg, code, fileInfo := s.repo.SignIn(user)
	if err != nil {
		return models.Tokens{}, models.UserSignIn{}, err, errKey, errMsg, code, fileInfo
	}

	// compare passwords
	if err, errKey, errMsg, code, fileInfo := password_hash.ComparePasswords(userDB.Password, user.Password); err != nil {
		return models.Tokens{}, models.UserSignIn{}, err, errKey, errMsg, code, fileInfo
	}

	// create access-token
	accessToken, err, errKey, errMsg, code, fileInfo := tokens.CreateAccessToken(userDB.Id, accessTokenExp, secretKey)
	if err != nil {
		return models.Tokens{}, models.UserSignIn{}, err, errKey, errMsg, code, fileInfo
	}

	// create refresh-token
	refreshToken, err, errKey, errMsg, code, fileInfo := tokens.CreateRefreshToken(userDB.Id, refreshTokenExp, secretKey)
	if err != nil {
		return models.Tokens{}, models.UserSignIn{}, err, errKey, errMsg, code, fileInfo
	}

	return models.Tokens{AccessToken: accessToken, RefreshToken: refreshToken}, userDB, nil, "", "", codes.OK, ""

}
