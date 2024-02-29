package tokens_grpc

import (
	"runtime"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"
	"google.golang.org/grpc/codes"
)

func CreateAccessToken(userId int64, exp time.Duration, secretKey string) (string, error, string, string, codes.Code, string) {
	claims := jwt.MapClaims{"sub": userId, "exp": time.Now().Add(exp).Unix()}
	tokenString, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(secretKey))
	if err != nil {
		return "", err, err.Error(), "Internal Server Error", codes.Internal, getFileInfo("tokens.go")
	}
	return tokenString, nil, "", "", codes.OK, ""
}

func CreateRefreshToken(userId int64, exp time.Duration, secretKey string) (string, error, string, string, codes.Code, string) {
	claims := jwt.MapClaims{"sub": userId, "exp": time.Now().Add(exp).Unix()}
	tokenString, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(secretKey))
	if err != nil {
		return "", err, err.Error(), "Internal Server Error", codes.Internal, getFileInfo("tokens.go")
	}
	return tokenString, nil, "", "", codes.OK, ""
}

func getFileInfo(fileName string) string {
	_, _, line, _ := runtime.Caller(1)
	return "internal/grpc/utils/jwt/" + fileName + " line: " + strconv.Itoa(line)
}
