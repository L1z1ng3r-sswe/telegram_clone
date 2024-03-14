package chat_postgres_rest

import (
	"database/sql"
	"net/http"

	models_rest "github.com/L1z1ng3r-sswe/telegram_clone/app/internal/rest/domain/models"
)

func (repo *repo) GetAllCommunityMessages(communityId string) ([]models_rest.Message, *models_rest.Response) {
	stmt, err := repo.db.Preparex(`SELECT * FROM community_messages WHERE community_id = $1`)
	if err != nil {
		return nil, &models_rest.Response{err, "Internal Server Error", err.Error(), http.StatusInternalServerError, getFileInfo("get-all-community-messages.go")}
	}
	defer stmt.Close()

	var messages []models_rest.Message
	if err := stmt.Select(&messages, communityId); err != nil {
		if err == sql.ErrNoRows {
			return nil, &models_rest.Response{err, "Not Found", err.Error(), http.StatusNotFound, getFileInfo("get-all-community-messages.go")}
		}
		return nil, &models_rest.Response{err, "Internal Server Error", err.Error(), http.StatusInternalServerError, getFileInfo("get-all-community-messages.go")}
	}

	return messages, nil
}
