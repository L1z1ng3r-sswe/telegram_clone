package chat_service_rest

import (
	"runtime"
	"strconv"
	"time"

	models_rest "github.com/L1z1ng3r-sswe/telegram_clone/app/internal/rest/domain/models"
)

type repo interface {
	CreateCommunity(comm models_rest.Community) (models_rest.Community, *models_rest.Response)
	JoinCommunity(communityMember models_rest.CommunityMember) (models_rest.CommunityMember, *models_rest.Response)
	CreateCommunityMessage(msg models_rest.Message) (models_rest.Message, *models_rest.Response)
	CreateBIChat(chat models_rest.BIChat) (models_rest.BIChat, *models_rest.Response)
	CreateBIMessage(msg models_rest.Message) (models_rest.Message, *models_rest.Response)
	GetAllBIMessages(biChatId string) ([]models_rest.Message, *models_rest.Response)
	GetAllCommunityMessages(communityId string) ([]models_rest.Message, *models_rest.Response)
}

type cache interface {
	CacheData(key string, val interface{}, exp time.Duration) *models_rest.Response
	GetAllBIMessages(key string) ([]models_rest.Message, *models_rest.Response)
	GetAllCommunityMessages(communityId string) ([]models_rest.Message, *models_rest.Response)
}

type service struct {
	repo  repo
	cache cache
}

func New(repo repo, cache cache) *service {
	return &service{
		repo:  repo,
		cache: cache,
	}
}

func getFileInfo(fileName string) string {
	_, _, line, _ := runtime.Caller(1)
	return "internal/rest/service/chat/" + fileName + " line: " + strconv.Itoa(line)
}
