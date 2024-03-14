package chat_service_rest

import (
	"strconv"

	models_rest "github.com/L1z1ng3r-sswe/telegram_clone/app/internal/rest/domain/models"
)

func (service *service) JoinCommunity(communityMember models_rest.CommunityMember) (models_rest.CommunityMember, *models_rest.Response) {
	// create a unique id for the new member
	communityMember.Id = strconv.FormatInt(communityMember.UserId, 10) + strconv.FormatInt(communityMember.CommunityId, 10)

	communityMemberDB, err := service.repo.JoinCommunity(communityMember)
	if err != nil {
		return models_rest.CommunityMember{}, err
	}

	return communityMemberDB, nil
}
