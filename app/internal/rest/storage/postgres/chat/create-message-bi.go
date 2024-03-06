package chat_postgres_rest

import (
	"net/http"

	models_rest "github.com/L1z1ng3r-sswe/telegram_clone/app/internal/rest/domain/models"
)

func (repo *repo) CreateBIMessage(msg models_rest.Message) (models_rest.Message, error, string, string, int, string) {
	stmt, err := repo.db.Preparex("INSERT INTO bi_chat_messages (sender_id, bi_chat_id, content) VALUES ($1, $2, $3) RETURNING *")
	if err != nil {
		return models_rest.Message{}, err, "Internal Server Error", err.Error(), http.StatusInternalServerError, getFileInfo("create-message-community.go")
	}
	defer stmt.Close()

	if err := stmt.Get(&msg, msg.SenderId, msg.BIChatId, msg.Content); err != nil {
		return models_rest.Message{}, err, "Internal Server Error", err.Error(), http.StatusInternalServerError, getFileInfo("create-message-community.go")
	}

	return msg, nil, "", "", http.StatusOK, ""
}
