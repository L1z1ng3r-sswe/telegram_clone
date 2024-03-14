package chat_service_rest

import (
	"errors"
	"fmt"
	"net/http"

	models_rest "github.com/L1z1ng3r-sswe/telegram_clone/app/internal/rest/domain/models"
)

func (service *service) GetAllBIMessages(biChatId string) ([]models_rest.Message, *models_rest.Response) {
	// verify
	if biChatId == "" {
		return nil, &models_rest.Response{errors.New("bi_chat_id is required"), "Bad Request", "bi_chat_id is required", http.StatusBadRequest, getFileInfo("get-all-bi-messages.go")}
	}

	// retrieve from the cache
	cacheKey := fmt.Sprintf("bi-messages:%d", biChatId)
	messages, err := service.cache.GetAllBIMessages(cacheKey)
	if err == nil {
		return messages, err
	} else if err.ErrKey != "Not Found" {
		return nil, err
	}

	// retrieve from the db
	messages, err = service.repo.GetAllBIMessages(biChatId)
	if err != nil {
		return nil, err
	}

	// store to the cache
	if err := service.cache.CacheData(cacheKey, messages, 0); err != nil {
		return nil, err
	}

	return messages, nil
}
