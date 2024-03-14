package validation_rest

import (
	"errors"
	"net/http"

	models_rest "github.com/L1z1ng3r-sswe/telegram_clone/app/internal/rest/domain/models"
)

func ValidationCommunityMessage(msg models_rest.Message) *models_rest.Response {

	if msg.Type != "community-message" {
		return &models_rest.Response{errors.New("Invalid message type"), "Bad Request", "Invalid message type", http.StatusBadRequest, getFileInfo("message-validation.go")}
	} else if len(msg.Content) == 0 {
		return &models_rest.Response{errors.New("Content is too short"), "Bad Request", "Content is too short", http.StatusBadRequest, getFileInfo("message-validation.go")}
	}

	return nil
}

func ValidationBIMessage(msg models_rest.Message) *models_rest.Response {

	if msg.Type != "bi-message" {
		return &models_rest.Response{errors.New("Invalid message type"), "Bad Request", "Invalid message type", http.StatusBadRequest, getFileInfo("message-validation.go")}
	} else if len(msg.Content) == 0 {
		return &models_rest.Response{errors.New("Content is too short"), "Bad Request", "Content is too short", http.StatusBadRequest, getFileInfo("message-validation.go")}
	}

	return nil
}
