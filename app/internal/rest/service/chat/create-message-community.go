package chat_service_rest

import (
	"net/http"

	models_rest "github.com/L1z1ng3r-sswe/telegram_clone/app/internal/rest/domain/models"
	validation_rest "github.com/L1z1ng3r-sswe/telegram_clone/app/internal/rest/utils/validation"
)

func (service *service) CreateCommunityMessage(msg models_rest.Message) (models_rest.Message, error, string, string, int, string) {
	// validate message type and content length
	if err, errKey, errMsg, code, fileInfo := validation_rest.ValidationCommunityMessage(msg); err != nil {
		return models_rest.Message{}, err, errKey, errMsg, code, fileInfo
	}

	msgDB, err, errKey, errMsg, code, fileInfo := service.repo.CreateCommunityMessage(msg)
	if err != nil {
		return models_rest.Message{}, err, errKey, errMsg, code, fileInfo
	}

	return msgDB, nil, "", "", http.StatusOK, ""
}
