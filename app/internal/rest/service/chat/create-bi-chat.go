package chat_service_rest

import (
	"strconv"

	models_rest "github.com/L1z1ng3r-sswe/telegram_clone/app/internal/rest/domain/models"
)

func (service *service) CreateBIChat(chat models_rest.BIChat) (models_rest.BIChat, *models_rest.Response) {
	// create a unique id for the new chat
	chat.Id = strconv.FormatInt(chat.FirstUserId, 10) + strconv.FormatInt(chat.SecondUserId, 10)

	chatDB, err := service.repo.CreateBIChat(chat)
	if err != nil {
		return models_rest.BIChat{}, err
	}

	return chatDB, nil
}
