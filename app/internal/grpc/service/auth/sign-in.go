package auth_service_grpc

import (
	"time"

	models_grpc "github.com/L1z1ng3r-sswe/telegram_clone/app/internal/grpc/domain/models"
	tokens_grpc "github.com/L1z1ng3r-sswe/telegram_clone/app/internal/grpc/utils/jwt"
	password_hash_grpc "github.com/L1z1ng3r-sswe/telegram_clone/app/internal/grpc/utils/password"
	validation_grpc "github.com/L1z1ng3r-sswe/telegram_clone/app/internal/grpc/utils/validation"
	"google.golang.org/grpc/codes"
)

func (s *Service) SignIn(user models_grpc.UserSignIn, accessTokenExp time.Duration, refreshTokenExp time.Duration, secretKey string) (models_grpc.Tokens, models_grpc.UserSignIn, error, string, string, codes.Code, string) {
	// validation
	err, errKey, errMsg, code, fileInfo := validation_grpc.ValidationSignIn(user.Email, user.Password)
	if err != nil {
		return models_grpc.Tokens{}, models_grpc.UserSignIn{}, err, errKey, errMsg, code, fileInfo
	}

	// ger the user from db from db by email
	userDB, err, errKey, errMsg, code, fileInfo := s.repo.SignIn(user)
	if err != nil {
		return models_grpc.Tokens{}, models_grpc.UserSignIn{}, err, errKey, errMsg, code, fileInfo
	}

	// compare passwords
	if err, errKey, errMsg, code, fileInfo := password_hash_grpc.ComparePasswords(userDB.Password, user.Password); err != nil {
		return models_grpc.Tokens{}, models_grpc.UserSignIn{}, err, errKey, errMsg, code, fileInfo
	}

	// create access-token
	accessToken, err, errKey, errMsg, code, fileInfo := tokens_grpc.CreateAccessToken(userDB.Id, accessTokenExp, secretKey)
	if err != nil {
		return models_grpc.Tokens{}, models_grpc.UserSignIn{}, err, errKey, errMsg, code, fileInfo
	}

	// create refresh-token
	refreshToken, err, errKey, errMsg, code, fileInfo := tokens_grpc.CreateRefreshToken(userDB.Id, refreshTokenExp, secretKey)
	if err != nil {
		return models_grpc.Tokens{}, models_grpc.UserSignIn{}, err, errKey, errMsg, code, fileInfo
	}

	return models_grpc.Tokens{AccessToken: accessToken, RefreshToken: refreshToken}, userDB, nil, "", "", codes.OK, ""

}
