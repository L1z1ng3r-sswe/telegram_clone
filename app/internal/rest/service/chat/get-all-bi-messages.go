package chat_service_rest

import (
	"errors"
	"net/http"

	models_rest "github.com/L1z1ng3r-sswe/telegram_clone/app/internal/rest/domain/models"
)

func (service *service) GetAllBIMessages(biChatId string) ([]models_rest.Message, error, string, string, int, string) {

	if biChatId == "" {
		return nil, errors.New("bi_chat_id is required"), "Bad Request", "bi_chat_id is required", http.StatusBadRequest, getFileInfo("get-all-bi-messages.go")
	}

	messages, err, errKey, errMsg, code, fileInfo := service.repo.GetAllBIMessages(biChatId)
	if err != nil {
		return nil, err, errKey, errMsg, code, fileInfo
	}

	return messages, nil, "", "", http.StatusOK, ""
}
