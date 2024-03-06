package chat_service_rest

import (
	"runtime"
	"strconv"

	models_rest "github.com/L1z1ng3r-sswe/telegram_clone/app/internal/rest/domain/models"
)

type repo interface {
	CreateCommunity(comm models_rest.Community) (models_rest.Community, error, string, string, int, string)
	JoinCommunity(communityMember models_rest.CommunityMember) (models_rest.CommunityMember, error, string, string, int, string)
	CreateCommunityMessage(msg models_rest.Message) (models_rest.Message, error, string, string, int, string)
	CreateBIChat(chat models_rest.BIChat) (models_rest.BIChat, error, string, string, int, string)
	CreateBIMessage(msg models_rest.Message) (models_rest.Message, error, string, string, int, string)
	GetAllBIMessages(biChatId string) ([]models_rest.Message, error, string, string, int, string)
}

type service struct {
	repo repo
}

func New(repo repo) *service {
	return &service{
		repo: repo,
	}
}

func getFileInfo(fileName string) string {
	_, _, line, _ := runtime.Caller(1)
	return "internal/rest/service/chat/" + fileName + " line: " + strconv.Itoa(line)
}
