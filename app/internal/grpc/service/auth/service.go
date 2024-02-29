package auth_service_grpc

import (
	"runtime"
	"strconv"

	models_grpc "github.com/L1z1ng3r-sswe/telegram_clone/app/internal/grpc/domain/models"
	"google.golang.org/grpc/codes"
)

type repo interface {
	SignUp(user models_grpc.UserSignUp) (models_grpc.UserSignUp, error, string, string, codes.Code, string)
	SignIn(user models_grpc.UserSignIn) (models_grpc.UserSignIn, error, string, string, codes.Code, string)
}

type Service struct {
	repo repo
}

func New(repo repo) Service {
	return Service{
		repo: repo,
	}
}

func getFileInfo(fileName string) string {
	_, _, line, _ := runtime.Caller(1)
	return "internal/grpc/service/auth/" + fileName + " line: " + strconv.Itoa(line)
}
