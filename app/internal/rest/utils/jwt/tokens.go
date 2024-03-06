package tokens_rest

import (
	"errors"
	"net/http"
	"runtime"
	"strconv"
	"strings"
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

func IsTokenValid(tokenString string, secretKey string) (int64, error, string, string, int, string) {

	tokenSlice := strings.Split(tokenString, " ")

	if len(tokenSlice) != 2 || tokenSlice[0] != "Bearer" {
		return 0, errors.New("Invalid authorization header format"), "Unauthorized", "Invalid authorization header format", http.StatusUnauthorized, getFileInfo("tokens.go")
	}

	token, err := jwt.Parse(tokenSlice[1], func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("Unexpected signing method: " + token.Header["alg"].(string))
		}
		return []byte(secretKey), nil
	})

	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return 0, errors.New("Invalid token"), "Unauthorized", "Token is malformed", http.StatusUnauthorized, getFileInfo("tokens.go")
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return 0, errors.New("Invalid token"), "Unauthorized", "Token has expired", http.StatusUnauthorized, getFileInfo("tokens.go")
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return 0, errors.New("Invalid token"), "Unauthorized", "Token not yet valid", http.StatusUnauthorized, getFileInfo("tokens.go")
			} else {
				return 0, errors.New("Invalid token"), "Unauthorized", "Couldn't handle this token", http.StatusUnauthorized, getFileInfo("tokens.go")
			}
		}
		return 0, errors.New("Invalid token"), "Unauthorized", "Couldn't handle this token", http.StatusUnauthorized, getFileInfo("tokens.go")
	}

	if !token.Valid {
		return 0, errors.New("Invalid token"), "Unauthorized", "Invalid token", http.StatusUnauthorized, getFileInfo("tokens.go")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return 0, errors.New("Invalid token claims"), "Unauthorized", "Invalid token claims", http.StatusUnauthorized, getFileInfo("tokens.go")
	}

	userIdFloat, ok := claims["sub"].(float64)
	if !ok {
		return 0, errors.New("Invalid userId in token"), "Unauthorized", "Invalid userId in token", http.StatusUnauthorized, getFileInfo("tokens.go")
	}

	userId := int64(userIdFloat)
	return userId, nil, "", "", http.StatusOK, ""
}

func getFileInfo(fileName string) string {
	_, _, line, _ := runtime.Caller(1)
	return "internal/rest/utils/jwt/" + fileName + " line: " + strconv.Itoa(line)
}
