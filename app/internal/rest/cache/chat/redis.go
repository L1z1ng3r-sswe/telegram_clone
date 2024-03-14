package chat_redis_rest

import (
	"context"
	"runtime"
	"strconv"

	"github.com/redis/go-redis/v9"
)

type Cache struct {
	rdb *redis.Client
	ctx context.Context
}

func New() *Cache {
	return &Cache{
		rdb: redis.NewClient(&redis.Options{
			Addr:     "localhost:6379",
			Password: "",
			DB:       0,
		}),
		ctx: context.Background(),
	}
}

func getFileInfo(fileName string) string {
	_, _, line, _ := runtime.Caller(1)
	return "internal/rest/cache/chat/" + fileName + " line: " + strconv.Itoa(line)
}
