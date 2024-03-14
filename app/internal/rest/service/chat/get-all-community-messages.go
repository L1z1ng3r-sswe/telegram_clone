package chat_service_rest

import (
	"errors"
	"fmt"
	"net/http"

	models_rest "github.com/L1z1ng3r-sswe/telegram_clone/app/internal/rest/domain/models"
)

func (service *service) GetAllCommunityMessages(communityId string) ([]models_rest.Message, *models_rest.Response) {
	// verify
	if communityId == "" {
		return nil, &models_rest.Response{errors.New("community_id is required"), "Bad Request", "community_id is required", http.StatusBadRequest, getFileInfo("get-all-community-messages.go")}
	}

	// retrieve from the cache
	cacheKey := fmt.Sprintf("community-message:%s", communityId)
	messages, err := service.cache.GetAllCommunityMessages(cacheKey)
	if err == nil {
		fmt.Println("from the cache")
		return messages, nil
	} else if err.ErrKey != "Not Found" {
		return nil, err
	}

	// retrieve from the db
	messages, err = service.repo.GetAllCommunityMessages(communityId)
	if err != nil {
		return nil, err
	}

	fmt.Println("store to the cache")

	err = service.cache.CacheData(cacheKey, messages, 0)
	if err != nil {
		return nil, err
	}
	// store to the cache
	fmt.Println("from db")
	return messages, nil
}
