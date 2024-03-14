package validation_rest

import (
	"errors"
	"fmt"
	"net/http"
	"regexp"

	models_rest "github.com/L1z1ng3r-sswe/telegram_clone/app/internal/rest/domain/models"
)

func ValidationSignUp(email, password string) *models_rest.Response {
	err := isEmailValidSignUp(email)
	if err != nil {
		return err
	}

	err = isPasswordValidSignUp(password)
	if err != nil {
		return err
	}

	return nil
}

func isEmailValidSignUp(email string) *models_rest.Response {
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	if !emailRegex.MatchString(email) {
		return &models_rest.Response{errors.New("Invalid email format"), "Bad Request", "Invalid email format", http.StatusBadRequest, getFileInfo("user_validation.go")}
	}

	return nil
}

func isPasswordValidSignUp(password string) *models_rest.Response {
	if len(password) < 8 {
		return &models_rest.Response{errors.New("Password is too short"), "Bad Request", "Password is too short", http.StatusBadRequest, getFileInfo("user_validation.go")}
	}

	return nil
}

//! ======================================== sign-in =================================

func ValidationSignIn(email, password string) *models_rest.Response {
	err := isEmailValidSignIn(email)
	if err != nil {
		return err
	}

	err = isPasswordValidSignIn(password)
	if err != nil {
		return err
	}

	return nil
}

func isEmailValidSignIn(email string) *models_rest.Response {
	fmt.Println(email)
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	if !emailRegex.MatchString(email) {
		return &models_rest.Response{errors.New("Invalid email format"), "Bad Request", "Invalid email format", http.StatusBadRequest, getFileInfo("user_validation.go")}
	}

	return nil
}

func isPasswordValidSignIn(password string) *models_rest.Response {
	if len(password) < 8 {
		return &models_rest.Response{errors.New("Password is too short"), "Bad Request", "Password is too short", http.StatusBadRequest, getFileInfo("user_validation.go")}
	}

	return nil
}
