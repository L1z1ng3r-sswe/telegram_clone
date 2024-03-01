package chat_rest

import (
	"runtime"
	"strconv"
	"time"

	logger "github.com/L1z1ng3r-sswe/telegram_clone/app/internal/lib/zerolog"
	models_rest "github.com/L1z1ng3r-sswe/telegram_clone/app/internal/rest/domain/models"
)

type service interface {
	SessionCreation(user models_rest.UserSignIn, accessTokenExp time.Duration, refreshTokenExp time.Duration, secretKey string) (models_rest.Tokens, models_rest.UserDB, error, string, string, int, string)
}

type server struct {
	service         service
	log             *logger.Logger
	accessTokenExp  time.Duration
	refreshTokenExp time.Duration
	secretKey       string
	hub             *Hub
}

func New(log *logger.Logger, service service, accessTokenExp time.Duration, refreshTokenExp time.Duration, secretKey string) *server {
	return &server{
		service: service,
		hub: &Hub{
			BroadcastBIMessages: make(chan models_rest.BIMessage, 10),
			BIChats:             make(map[int64]models_rest.BIChats),
			Clients:             make(map[int64]*Client),
		},
	}
}

func getFileInfo(fileName string) string {
	_, _, line, _ := runtime.Caller(1)
	return "internal/rest/server/chat" + fileName + " line: " + strconv.Itoa(line)
}
