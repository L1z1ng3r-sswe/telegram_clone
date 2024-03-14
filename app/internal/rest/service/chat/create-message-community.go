package chat_service_rest

import (
	models_rest "github.com/L1z1ng3r-sswe/telegram_clone/app/internal/rest/domain/models"
	validation_rest "github.com/L1z1ng3r-sswe/telegram_clone/app/internal/rest/utils/validation"
)

func (service *service) CreateCommunityMessage(msg models_rest.Message) (models_rest.Message, *models_rest.Response) {
	// validate message type and content length
	if err := validation_rest.ValidationCommunityMessage(msg); err != nil {
		return models_rest.Message{}, err
	}

	msgDB, err := service.repo.CreateCommunityMessage(msg)
	if err != nil {
		return models_rest.Message{}, err
	}

	return msgDB, nil
}
