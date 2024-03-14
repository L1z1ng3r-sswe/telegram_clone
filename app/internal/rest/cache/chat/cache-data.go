package chat_redis_rest

import (
	"encoding/json"
	"net/http"
	"time"

	models_rest "github.com/L1z1ng3r-sswe/telegram_clone/app/internal/rest/domain/models"
)

func (cache *Cache) CacheData(key string, val interface{}, exp time.Duration) *models_rest.Response {
	jsonData, err := json.Marshal(val)
	if err != nil {
		return &models_rest.Response{err, "Internal Server Error", err.Error(), http.StatusInternalServerError, getFileInfo("cache-data.go")}
	}

	if err := cache.rdb.Set(cache.ctx, key, jsonData, exp).Err(); err != nil {
		return &models_rest.Response{err, "Internal Server Error", err.Error(), http.StatusInternalServerError, getFileInfo("cache-data.go")}
	}

	return nil
}
