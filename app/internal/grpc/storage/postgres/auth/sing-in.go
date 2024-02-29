package auth_postgres_grpc

import (
	"database/sql"

	models_grpc "github.com/L1z1ng3r-sswe/telegram_clone/app/internal/grpc/domain/models"
	"google.golang.org/grpc/codes"
)

func (r *repo) SignIn(user models_grpc.UserSignIn) (models_grpc.UserSignIn, error, string, string, codes.Code, string) {
	stmt, err := r.db.Preparex(`SELECT id, password FROM users WHERE email=$1 LIMIT 1`)
	defer stmt.Close()
	if err != nil {
		return models_grpc.UserSignIn{}, err, "Internal Server Error", "", codes.Internal, getFileInfo("auth.go")
	}

	if err := stmt.Get(&user, user.Email); err != nil {
		if err == sql.ErrNoRows {
			return models_grpc.UserSignIn{}, err, "Bad Request", "Wrong email", codes.InvalidArgument, getFileInfo("auth.go")
		}

		return models_grpc.UserSignIn{}, err, "Internal Server Error", err.Error(), codes.Internal, getFileInfo("auth.go")
	}

	return user, nil, "", "", codes.OK, ""
}
