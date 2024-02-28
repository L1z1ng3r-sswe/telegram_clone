package auth_service

import (
	"runtime"
	"strconv"

	"github.com/L1z1ng3r-sswe/telegram_clone/app/internal/domain/models"
	"google.golang.org/grpc/codes"
)

type repo interface {
	SignUp(user models.UserSignUp) (models.UserSignUp, error, string, string, codes.Code, string)
	SignIn(user models.UserSignIn) (models.UserSignIn, error, string, string, codes.Code, string)
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
	return "internal/service/auth/" + fileName + " line: " + strconv.Itoa(line)
}
