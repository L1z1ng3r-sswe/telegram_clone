package chat_postgres_rest

import (
	"net/http"

	models_rest "github.com/L1z1ng3r-sswe/telegram_clone/app/internal/rest/domain/models"
)

func (repo *repo) CreateBIChat(chat models_rest.BIChat) (models_rest.BIChat, error, string, string, int, string) {
	stmt, err := repo.db.Preparex("INSERT INTO bi_chats (id, first_user_id, second_user_id) VALUES ($1, $2, $3) RETURNING *")
	if err != nil {
		return models_rest.BIChat{}, err, "Internal Server Error", err.Error(), http.StatusInternalServerError, getFileInfo("create-bi-chat.go")
	}
	defer stmt.Close()

	if err := stmt.Get(&chat, chat.Id, chat.FirstUserId, chat.SecondUserId); err != nil {
		if isUniqueConstraintViolation(err) {
			return models_rest.BIChat{}, err, "Bad Request", "Chat is already exist", http.StatusBadRequest, getFileInfo("create-bi-chat.go")
		}
		return models_rest.BIChat{}, err, "Bad Request", err.Error(), http.StatusBadRequest, getFileInfo("create-bi-chat.go")
	}

	return chat, nil, "", "", http.StatusOK, ""

}
