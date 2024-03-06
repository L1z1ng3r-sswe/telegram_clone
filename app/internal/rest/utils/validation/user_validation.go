package validation_rest

import (
	"errors"
	"fmt"
	"net/http"
	"regexp"
)

func ValidationSignUp(email, password string) (error, string, string, int, string) {
	err, errKey, errMsg, code, fileInfo := isEmailValidSignUp(email)
	if err != nil {
		return err, errKey, errMsg, code, fileInfo
	}

	err, errKey, errMsg, code, fileInfo = isPasswordValidSignUp(password)
	if err != nil {
		return err, errKey, errMsg, code, fileInfo
	}

	return nil, "", "", http.StatusOK, ""
}

func isEmailValidSignUp(email string) (error, string, string, int, string) {
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	if !emailRegex.MatchString(email) {
		return errors.New("Invalid email format"), "Bad Request", "Invalid email format", http.StatusBadRequest, getFileInfo("user_validation.go")
	}

	return nil, "", "", http.StatusOK, ""
}

func isPasswordValidSignUp(password string) (error, string, string, int, string) {
	if len(password) < 8 {
		return errors.New("Password is too short"), "Bad Request", "Password is too short", http.StatusBadRequest, getFileInfo("user_validation.go")
	}

	return nil, "", "", http.StatusOK, ""
}

//! ======================================== sign-in =================================

func ValidationSignIn(email, password string) (error, string, string, int, string) {
	err, errKey, errMsg, code, fileInfo := isEmailValidSignIn(email)
	if err != nil {
		return err, errKey, errMsg, code, fileInfo
	}

	err, errKey, errMsg, code, fileInfo = isPasswordValidSignIn(password)
	if err != nil {
		return err, errKey, errMsg, code, fileInfo
	}

	return nil, "", "", http.StatusOK, ""
}

func isEmailValidSignIn(email string) (error, string, string, int, string) {
	fmt.Println(email)
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	if !emailRegex.MatchString(email) {
		return errors.New("Invalid email format"), "Bad Request", "Invalid email format", http.StatusBadRequest, getFileInfo("user_validation.go")
	}

	return nil, "", "", http.StatusOK, ""
}

func isPasswordValidSignIn(password string) (error, string, string, int, string) {
	if len(password) < 8 {
		return errors.New("Password is too short"), "Bad Request", "Password is too short", http.StatusBadRequest, getFileInfo("user_validation.go")
	}

	return nil, "", "", http.StatusOK, ""
}
