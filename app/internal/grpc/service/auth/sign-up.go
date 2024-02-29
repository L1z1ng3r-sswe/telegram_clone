package auth_service_grpc

import (
	"time"

	models_grpc "github.com/L1z1ng3r-sswe/telegram_clone/app/internal/grpc/domain/models"
	tokens_grpc "github.com/L1z1ng3r-sswe/telegram_clone/app/internal/grpc/utils/jwt"
	password_hash_grpc "github.com/L1z1ng3r-sswe/telegram_clone/app/internal/grpc/utils/password"
	validation_grpc "github.com/L1z1ng3r-sswe/telegram_clone/app/internal/grpc/utils/validation"
	"google.golang.org/grpc/codes"
)

func (s *Service) SignUp(user models_grpc.UserSignUp, accessTokenExp time.Duration, refreshTokenExp time.Duration, secretKey string) (models_grpc.Tokens, models_grpc.UserSignUp, error, string, string, codes.Code, string) {
	// validation
	err, errKey, errMsg, code, fileInfo := validation_grpc.ValidationSignUp(user.Email, user.Password)
	if err != nil {
		return models_grpc.Tokens{}, models_grpc.UserSignUp{}, err, errKey, errMsg, code, fileInfo
	}

	// hash password
	user.Password, err, errKey, errMsg, code, fileInfo = password_hash_grpc.PasswordHasher(user.Password)
	if err != nil {
		return models_grpc.Tokens{}, models_grpc.UserSignUp{}, err, errKey, errMsg, code, fileInfo
	}

	// save new user in db
	userDB, err, errKey, errMsg, code, fileInfo := s.repo.SignUp(user)
	if err != nil {
		return models_grpc.Tokens{}, models_grpc.UserSignUp{}, err, errKey, errMsg, code, fileInfo
	}

	// create access-token
	accessToken, err, errKey, errMsg, code, fileInfo := tokens_grpc.CreateAccessToken(userDB.Id, accessTokenExp, secretKey)
	if err != nil {
		return models_grpc.Tokens{}, models_grpc.UserSignUp{}, err, errKey, errMsg, code, fileInfo
	}

	// create refresh-token
	refreshToken, err, errKey, errMsg, code, fileInfo := tokens_grpc.CreateRefreshToken(userDB.Id, refreshTokenExp, secretKey)
	if err != nil {
		return models_grpc.Tokens{}, models_grpc.UserSignUp{}, err, errKey, errMsg, code, fileInfo
	}

	return models_grpc.Tokens{AccessToken: accessToken, RefreshToken: refreshToken}, userDB, nil, "", "", codes.OK, ""
}
