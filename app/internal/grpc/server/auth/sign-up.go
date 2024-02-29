package auth_grpc

import (
	"context"

	auth "github.com/L1z1ng3r-sswe/telegram_clone-proto_contract/gen/go/auth"
	models_grpc "github.com/L1z1ng3r-sswe/telegram_clone/app/internal/grpc/domain/models"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *userServerAPI) SignUp(ctx context.Context, req *auth.SignUpRequest) (*auth.SignUpResponse, error) {
	if req.GetEmail() == "" {
		return nil, status.Error(codes.InvalidArgument, "email is required")
	}

	if req.GetPassword() == "" {
		return nil, status.Error(codes.InvalidArgument, "password is required")
	}

	user := models_grpc.UserSignUp{
		Email:    req.Email,
		Password: req.Password,
	}

	tokens, userDB, err, errKey, errMsg, code, fileInfo := server.service.SignUp(user, server.accessTokenExp, server.refreshTokenExp, server.secretKey)
	if err != nil {
		server.log.Err(errKey, errMsg, fileInfo)
		return nil, status.Error(code, errMsg)
	}

	server.log.Inf("signed up a new user", "id", userDB.Id)
	return &auth.SignUpResponse{AccessToken: tokens.AccessToken, RefreshToken: tokens.RefreshToken}, nil
}
