package password_hash_rest

import (
	"net/http"
	"runtime"
	"strconv"

	"golang.org/x/crypto/bcrypt"
)

func PasswordHasher(password string) (string, error, string, string, int, string) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err, "Internal Server Error", err.Error(), http.StatusInternalServerError, getFileInfo("password.go")
	}

	return string(hashedPassword), nil, "", "", http.StatusOK, ""
}

func ComparePasswords(hashedPassword, password string) (error, string, string, int, string) {
	if err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password)); err != nil {
		return err, "Bad Request", "Wrong password", http.StatusBadRequest, getFileInfo("password.go")
	}

	return nil, "", "", http.StatusOK, ""
}

func getFileInfo(fileName string) string {
	_, _, line, _ := runtime.Caller(1)
	return "internal/rest/utils/password/" + fileName + " line: " + strconv.Itoa(line)
}
