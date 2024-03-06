package chat_postgres_rest

import (
	"net/http"

	models_rest "github.com/L1z1ng3r-sswe/telegram_clone/app/internal/rest/domain/models"
)

func (repo *repo) JoinCommunity(communityMember models_rest.CommunityMember) (models_rest.CommunityMember, error, string, string, int, string) {
	stmt, err := repo.db.Preparex("INSERT INTO community_members (id, user_id, community_id) VALUES ($1, $2, $3) RETURNING *")
	if err != nil {
		return models_rest.CommunityMember{}, err, "Internal Server Error", err.Error(), http.StatusInternalServerError, getFileInfo("join-community.go")
	}
	defer stmt.Close()

	if err := stmt.Get(&communityMember, communityMember.Id, communityMember.UserId, communityMember.CommunityId); err != nil {
		if isUniqueConstraintViolation(err) {
			return models_rest.CommunityMember{}, err, "Bad Request", "Client is already in community", http.StatusBadRequest, getFileInfo("create-community.go")
		}
		return models_rest.CommunityMember{}, err, "Bad Request", err.Error(), http.StatusBadRequest, getFileInfo("join-community.go")
	}

	return communityMember, nil, "", "", http.StatusOK, ""
}
