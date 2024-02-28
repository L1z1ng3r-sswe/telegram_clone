package auth_service

import (
	"time"

	"github.com/L1z1ng3r-sswe/telegram_clone/app/internal/domain/models"
	password_hash "github.com/L1z1ng3r-sswe/telegram_clone/app/internal/lib/bcrypt"
	tokens "github.com/L1z1ng3r-sswe/telegram_clone/app/internal/lib/jwt"
	"github.com/L1z1ng3r-sswe/telegram_clone/app/internal/lib/validation"
	"google.golang.org/grpc/codes"
)

func (s *Service) SignUp(user models.UserSignUp, accessTokenExp time.Duration, refreshTokenExp time.Duration, secretKey string) (models.Tokens, models.UserSignUp, error, string, string, codes.Code, string) {
	// validation
	err, errKey, errMsg, code, fileInfo := validation.ValidationSignUp(user.Email, user.Password)
	if err != nil {
		return models.Tokens{}, models.UserSignUp{}, err, errKey, errMsg, code, fileInfo
	}

	// hash password
	user.Password, err, errKey, errMsg, code, fileInfo = password_hash.PasswordHasher(user.Password)
	if err != nil {
		return models.Tokens{}, models.UserSignUp{}, err, errKey, errMsg, code, fileInfo
	}

	// save new user in db
	userDB, err, errKey, errMsg, code, fileInfo := s.repo.SignUp(user)
	if err != nil {
		return models.Tokens{}, models.UserSignUp{}, err, errKey, errMsg, code, fileInfo
	}

	// create access-token
	accessToken, err, errKey, errMsg, code, fileInfo := tokens.CreateAccessToken(userDB.Id, accessTokenExp, secretKey)
	if err != nil {
		return models.Tokens{}, models.UserSignUp{}, err, errKey, errMsg, code, fileInfo
	}

	// create refresh-token
	refreshToken, err, errKey, errMsg, code, fileInfo := tokens.CreateRefreshToken(userDB.Id, refreshTokenExp, secretKey)
	if err != nil {
		return models.Tokens{}, models.UserSignUp{}, err, errKey, errMsg, code, fileInfo
	}

	return models.Tokens{AccessToken: accessToken, RefreshToken: refreshToken}, userDB, nil, "", "", codes.OK, ""
}
