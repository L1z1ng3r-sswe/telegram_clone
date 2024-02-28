package auth_postgres

import (
	"database/sql"

	"github.com/L1z1ng3r-sswe/telegram_clone/app/internal/domain/models"
	"google.golang.org/grpc/codes"
)

func (r *repo) SignIn(user models.UserSignIn) (models.UserSignIn, error, string, string, codes.Code, string) {
	stmt, err := r.db.Preparex(`SELECT id, password FROM users WHERE email=$1 LIMIT 1`)
	defer stmt.Close()
	if err != nil {
		return models.UserSignIn{}, err, "Internal Server Error", "", codes.Internal, getFileInfo("auth.go")
	}

	if err := stmt.Get(&user, user.Email); err != nil {
		if err == sql.ErrNoRows {
			return models.UserSignIn{}, err, "Bad Request", "Wrong email", codes.InvalidArgument, getFileInfo("auth.go")
		}

		return models.UserSignIn{}, err, "Internal Server Error", err.Error(), codes.Internal, getFileInfo("auth.go")
	}

	return user, nil, "", "", codes.OK, ""
}
