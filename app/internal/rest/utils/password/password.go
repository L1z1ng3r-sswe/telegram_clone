package password_hash_rest

import (
	"net/http"
	"runtime"
	"strconv"

	models_rest "github.com/L1z1ng3r-sswe/telegram_clone/app/internal/rest/domain/models"
	"golang.org/x/crypto/bcrypt"
)

func PasswordHasher(password string) (string, *models_rest.Response) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", &models_rest.Response{err, "Internal Server Error", err.Error(), http.StatusInternalServerError, getFileInfo("password.go")}
	}

	return string(hashedPassword), nil
}

func ComparePasswords(hashedPassword, password string) *models_rest.Response {
	if err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password)); err != nil {
		return &models_rest.Response{err, "Bad Request", "Wrong password", http.StatusBadRequest, getFileInfo("password.go")}
	}

	return nil
}

func getFileInfo(fileName string) string {
	_, _, line, _ := runtime.Caller(1)
	return "internal/rest/utils/password/" + fileName + " line: " + strconv.Itoa(line)
}
