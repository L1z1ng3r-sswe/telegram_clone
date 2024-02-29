package auth_postgres_grpc

import (
	models_grpc "github.com/L1z1ng3r-sswe/telegram_clone/app/internal/grpc/domain/models"
	"google.golang.org/grpc/codes"
)

func (r *repo) SignUp(user models_grpc.UserSignUp) (models_grpc.UserSignUp, error, string, string, codes.Code, string) {
	stmt, err := r.db.Preparex(`INSERT INTO users (email, password) VALUES($1, $2) RETURNING id`)
	defer stmt.Close()
	if err != nil {
		return models_grpc.UserSignUp{}, err, "Internal Server Error", err.Error(), codes.Internal, getFileInfo("auth.go")
	}

	if err := stmt.Get(&user, user.Email, user.Password); err != nil {
		if isUniqueConstraintViolation(err) {
			return models_grpc.UserSignUp{}, err, "Bad Request", "User already exists", codes.AlreadyExists, getFileInfo("auth.go")
		}

		return models_grpc.UserSignUp{}, err, "Bad Request", err.Error(), codes.InvalidArgument, getFileInfo("auth.go")
	}

	return user, nil, "", "", codes.OK, ""
}
