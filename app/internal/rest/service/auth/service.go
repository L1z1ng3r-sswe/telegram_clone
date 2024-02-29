package auth_service_rest

import (
	"runtime"
	"strconv"

	models_rest "github.com/L1z1ng3r-sswe/telegram_clone/app/internal/rest/domain/models"
)

type repo interface {
	SignUp(user models_rest.UserSignUp) (models_rest.UserSignUp, error, string, string, int, string)
	SignIn(user models_rest.UserSignIn) (models_rest.UserSignIn, error, string, string, int, string)
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
	return "internal/rest/service/auth/" + fileName + " line: " + strconv.Itoa(line)
}
