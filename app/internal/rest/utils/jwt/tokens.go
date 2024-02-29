package tokens_rest

import (
	"net/http"
	"runtime"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"
)

func CreateAccessToken(userId int64, exp time.Duration, secretKey string) (string, error, string, string, int, string) {
	claims := jwt.MapClaims{"sub": userId, "exp": time.Now().Add(exp).Unix()}
	tokenString, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(secretKey))
	if err != nil {
		return "", err, err.Error(), "Internal Server Error", http.StatusInternalServerError, getFileInfo("tokens.go")
	}
	return tokenString, nil, "", "", http.StatusOK, ""
}

func CreateRefreshToken(userId int64, exp time.Duration, secretKey string) (string, error, string, string, int, string) {
	claims := jwt.MapClaims{"sub": userId, "exp": time.Now().Add(exp).Unix()}
	tokenString, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(secretKey))
	if err != nil {
		return "", err, err.Error(), "Internal Server Error", http.StatusInternalServerError, getFileInfo("tokens.go")
	}
	return tokenString, nil, "", "", http.StatusOK, ""
}

func getFileInfo(fileName string) string {
	_, _, line, _ := runtime.Caller(1)
	return "internal/rest/utils/jwt/" + fileName + " line: " + strconv.Itoa(line)
}
