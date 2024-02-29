package validation_grpc

import (
	"errors"
	"regexp"

	"google.golang.org/grpc/codes"
)

func ValidationSignUp(email, password string) (error, string, string, codes.Code, string) {
	err, errKey, errMsg, code, fileInfo := isEmailValidSignUp(email)
	if err != nil {
		return err, errKey, errMsg, code, fileInfo
	}

	err, errKey, errMsg, code, fileInfo = isPasswordValidSignUp(password)
	if err != nil {
		return err, errKey, errMsg, code, fileInfo
	}

	return nil, "", "", codes.OK, ""
}

func isEmailValidSignUp(email string) (error, string, string, codes.Code, string) {
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	if !emailRegex.MatchString(email) {
		return errors.New("Invalid email format"), "Bad Request", "Invalid email format", codes.InvalidArgument, getFileInfo("user_validation.go")
	}

	return nil, "", "", codes.OK, ""
}

func isPasswordValidSignUp(password string) (error, string, string, codes.Code, string) {
	if len(password) < 8 {
		return errors.New("Password is too short"), "Bad Request", "Password is too short", codes.InvalidArgument, getFileInfo("user_validation.go")
	}

	return nil, "", "", codes.OK, ""
}

//! ======================================== sign-in =================================

func ValidationSignIn(email, password string) (error, string, string, codes.Code, string) {
	err, errKey, errMsg, code, fileInfo := isEmailValidSignIn(email)
	if err != nil {
		return err, errKey, errMsg, code, fileInfo
	}

	err, errKey, errMsg, code, fileInfo = isPasswordValidSignIn(password)
	if err != nil {
		return err, errKey, errMsg, code, fileInfo
	}

	return nil, "", "", codes.OK, ""
}

func isEmailValidSignIn(email string) (error, string, string, codes.Code, string) {
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	if !emailRegex.MatchString(email) {
		return errors.New("Invalid email format"), "Bad Request", "Invalid email format", codes.InvalidArgument, getFileInfo("user_validation.go")
	}

	return nil, "", "", codes.OK, ""
}

func isPasswordValidSignIn(password string) (error, string, string, codes.Code, string) {
	if len(password) < 8 {
		return errors.New("Password is too short"), "Bad Request", "Password is too short", codes.InvalidArgument, getFileInfo("user_validation.go")
	}

	return nil, "", "", codes.OK, ""
}
