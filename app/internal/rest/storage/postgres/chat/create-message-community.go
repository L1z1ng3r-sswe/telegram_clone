package chat_postgres_rest

import (
	"net/http"

	models_rest "github.com/L1z1ng3r-sswe/telegram_clone/app/internal/rest/domain/models"
)

func (repo *repo) CreateCommunityMessage(msg models_rest.Message) (models_rest.Message, *models_rest.Response) {
	stmt, err := repo.db.Preparex("INSERT INTO community_messages (sender_id, community_id, content) VALUES ($1, $2, $3) RETURNING *")
	if err != nil {
		return models_rest.Message{}, &models_rest.Response{err, "Internal Server Error", err.Error(), http.StatusInternalServerError, getFileInfo("create-message-community.go")}
	}
	defer stmt.Close()

	if err := stmt.Get(&msg, msg.SenderId, msg.CommunityId, msg.Content); err != nil {
		return models_rest.Message{}, &models_rest.Response{err, "Internal Server Error", err.Error(), http.StatusInternalServerError, getFileInfo("create-message-community.go")}
	}

	return msg, nil
}
