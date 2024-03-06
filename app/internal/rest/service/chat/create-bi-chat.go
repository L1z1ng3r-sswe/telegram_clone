package chat_service_rest

import (
	"net/http"
	"strconv"

	models_rest "github.com/L1z1ng3r-sswe/telegram_clone/app/internal/rest/domain/models"
)

func (service *service) CreateBIChat(chat models_rest.BIChat) (models_rest.BIChat, error, string, string, int, string) {
	// create a unique id for the new chat
	chat.Id = strconv.FormatInt(chat.FirstUserId, 10) + strconv.FormatInt(chat.SecondUserId, 10)

	chatDB, err, errKey, errMsg, code, fileInfo := service.repo.CreateBIChat(chat)
	if err != nil {
		return models_rest.BIChat{}, err, errKey, errMsg, code, fileInfo
	}

	return chatDB, nil, "", "", http.StatusOK, fileInfo
}
