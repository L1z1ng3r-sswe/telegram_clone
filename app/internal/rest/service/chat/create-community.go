package chat_service_rest

import (
	"errors"
	"net/http"

	models_rest "github.com/L1z1ng3r-sswe/telegram_clone/app/internal/rest/domain/models"
)

func (service *service) CreateCommunity(community models_rest.Community) (models_rest.Community, error, string, string, int, string) {

	if len(community.Name) <= 4 {
		return models_rest.Community{}, errors.New("Community lenght is too short"), "Bad Request", "Community lenght is too short", http.StatusBadRequest, getFileInfo("create-community.go")
	}

	communityDB, err, errMsg, errKey, code, fileInfo := service.repo.CreateCommunity(community)
	if err != nil {
		return models_rest.Community{}, err, errKey, errMsg, code, fileInfo
	}

	return communityDB, nil, "", "", http.StatusOK, ""
}
