package auth_grpc

import (
	"runtime"
	"strconv"
	"time"

	auth "github.com/L1z1ng3r-sswe/telegram_clone-proto_contract/gen/go/auth"
	models_grpc "github.com/L1z1ng3r-sswe/telegram_clone/app/internal/grpc/domain/models"
	logger "github.com/L1z1ng3r-sswe/telegram_clone/app/internal/lib/zerolog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

type service interface {
	SignUp(models_grpc.UserSignUp, time.Duration, time.Duration, string) (models_grpc.Tokens, models_grpc.UserSignUp, error, string, string, codes.Code, string)
	SignIn(models_grpc.UserSignIn, time.Duration, time.Duration, string) (models_grpc.Tokens, models_grpc.UserSignIn, error, string, string, codes.Code, string)
}

type userServerAPI struct {
	auth.UnimplementedAuthServer
	service         service
	log             *logger.Logger
	accessTokenExp  time.Duration
	refreshTokenExp time.Duration
	secretKey       string
}

func Register(log *logger.Logger, gRPC *grpc.Server, service service, accessTokenExp time.Duration, refreshTokenExp time.Duration, secretKey string) {
	auth.RegisterAuthServer(gRPC, &userServerAPI{log: log, service: service, accessTokenExp: accessTokenExp, refreshTokenExp: refreshTokenExp, secretKey: secretKey})
}

func getFileInfo(fileName string) string {
	_, _, line, _ := runtime.Caller(1)
	return "internal/grpc/server/auth" + fileName + " line: " + strconv.Itoa(line)
}
