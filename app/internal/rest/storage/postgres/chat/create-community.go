package chat_postgres_rest

import (
	"fmt"
	"net/http"

	models_rest "github.com/L1z1ng3r-sswe/telegram_clone/app/internal/rest/domain/models"
)

func (repo *repo) CreateCommunity(community models_rest.Community) (models_rest.Community, *models_rest.Response) {
	stmt, err := repo.db.Preparex("INSERT INTO communities (owner_id, name) VALUES ($1, $2) RETURNING *")
	if err != nil {
		return models_rest.Community{}, &models_rest.Response{err, "Internal Server Error", err.Error(), http.StatusInternalServerError, getFileInfo("create-community.go")}
	}
	defer stmt.Close()

	if err := stmt.Get(&community, community.OwnerId, community.Name); err != nil {
		if isUniqueConstraintViolation(err) {
			return models_rest.Community{}, &models_rest.Response{err, "Bad Request", "Community with this name is already exist", http.StatusBadRequest, getFileInfo("create-community.go")}
		}
		return models_rest.Community{}, &models_rest.Response{err, "Internal Server Error", err.Error(), http.StatusInternalServerError, getFileInfo("create-community.go")}
	}

	fmt.Println("db_community: ", community)
	return community, nil

}
