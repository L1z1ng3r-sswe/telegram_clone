package chat_redis_rest

import (
	"encoding/json"
	"net/http"

	models_rest "github.com/L1z1ng3r-sswe/telegram_clone/app/internal/rest/domain/models"
	"github.com/redis/go-redis/v9"
)

func (cache *Cache) GetAllBIMessages(key string) ([]models_rest.Message, *models_rest.Response) {
	val, err := cache.rdb.Get(cache.ctx, key).Result()
	if err == redis.Nil {
		return nil, &models_rest.Response{err, "Not Found", err.Error(), http.StatusNotFound, getFileInfo("cache-data.go")}
	} else if err != nil {
		return nil, &models_rest.Response{err, "Internal Server Error", err.Error(), http.StatusInternalServerError, getFileInfo("cache-data.go")}
	}

	var messages []models_rest.Message
	if err := json.Unmarshal([]byte(val), &messages); err != nil {
		return nil, &models_rest.Response{err, "Internal Server Error", err.Error(), http.StatusInternalServerError, getFileInfo("cache-data.go")}
	}

	return messages, nil
}
