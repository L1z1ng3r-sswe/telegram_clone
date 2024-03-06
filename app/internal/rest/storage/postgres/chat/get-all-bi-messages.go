package chat_postgres_rest

import (
	"net/http"

	models_rest "github.com/L1z1ng3r-sswe/telegram_clone/app/internal/rest/domain/models"
)

func (repo *repo) GetAllBIMessages(biChatId string) ([]models_rest.Message, error, string, string, int, string) {
	var messages []models_rest.Message

	stmt, err := repo.db.Preparex("SELECT * FROM bi_chat_messages WHERE bi_chat_id = $1")
	if err != nil {
		return nil, err, err.Error(), "Internal Server Error", http.StatusInternalServerError, getFileInfo("get-all-bi-messages.go")
	}
	defer stmt.Close()

	err = stmt.Select(&messages, biChatId)
	if err != nil {
		return nil, err, err.Error(), "Internal Server Error", http.StatusInternalServerError, getFileInfo("get-all-bi-messages.go")

	}
	return messages, nil, "", "", http.StatusOK, ""
}
