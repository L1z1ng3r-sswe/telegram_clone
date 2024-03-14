package auth_postgres_rest

import (
	"net/http"

	models_rest "github.com/L1z1ng3r-sswe/telegram_clone/app/internal/rest/domain/models"
)

func (r *repo) SignUp(user models_rest.UserSignUp) (models_rest.UserSignUp, *models_rest.Response) {
	stmt, err := r.db.Preparex(`INSERT INTO users (email, password) VALUES($1, $2) RETURNING id`)
	defer stmt.Close()
	if err != nil {
		return models_rest.UserSignUp{}, &models_rest.Response{err, "Internal Server Error", err.Error(), http.StatusInternalServerError, getFileInfo("auth.go")}
	}

	if err := stmt.Get(&user, user.Email, user.Password); err != nil {
		if isUniqueConstraintViolation(err) {
			return models_rest.UserSignUp{}, &models_rest.Response{err, "Bad Request", "User already exists", http.StatusBadRequest, getFileInfo("auth.go")}
		}

		return models_rest.UserSignUp{}, &models_rest.Response{err, "Bad Request", err.Error(), http.StatusInternalServerError, getFileInfo("auth.go")}
	}

	return user, nil
}
