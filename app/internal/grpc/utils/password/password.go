package password_hash_grpc

import (
	"runtime"
	"strconv"

	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
)

func PasswordHasher(password string) (string, error, string, string, codes.Code, string) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err, "Internal Server Error", err.Error(), codes.Internal, getFileInfo("password.go")
	}

	return string(hashedPassword), nil, "", "", codes.OK, ""
}

func ComparePasswords(hashedPassword, password string) (error, string, string, codes.Code, string) {
	if err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password)); err != nil {
		return err, "Bad Request", "Wrong password", codes.InvalidArgument, getFileInfo("password.go")
	}

	return nil, "", "", codes.OK, ""
}

func getFileInfo(fileName string) string {
	_, _, line, _ := runtime.Caller(1)
	return "internal/grpc/utils/password/" + fileName + " line: " + strconv.Itoa(line)
}
