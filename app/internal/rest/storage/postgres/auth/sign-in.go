package auth_postgres_rest

import (
	"database/sql"
	"net/http"

	models_rest "github.com/L1z1ng3r-sswe/telegram_clone/app/internal/rest/domain/models"
)

func (r *repo) SignIn(user models_rest.UserSignIn) (models_rest.UserDB, *models_rest.Response) {
	stmt, err := r.db.Preparex(`SELECT id, email, password  FROM users WHERE email=$1 LIMIT 1`)
	defer stmt.Close()
	if err != nil {
		return models_rest.UserDB{}, &models_rest.Response{err, "Internal Server Error", "", http.StatusInternalServerError, getFileInfo("auth.go")}
	}

	var userDB models_rest.UserDB
	if err := stmt.Get(&userDB, user.Email); err != nil {
		if err == sql.ErrNoRows {
			return models_rest.UserDB{}, &models_rest.Response{err, "Bad Request", "Wrong email", http.StatusBadRequest, getFileInfo("auth.go")}
		}

		return models_rest.UserDB{}, &models_rest.Response{err, "Internal Server Error", err.Error(), http.StatusInternalServerError, getFileInfo("auth.go")}
	}

	return userDB, nil
}
