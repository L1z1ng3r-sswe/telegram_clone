package auth_postgres_grpc

import (
	"runtime"
	"strconv"

	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	_ "github.com/lib/pq"
)

type repo struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) *repo {
	return &repo{
		db: db,
	}
}

func getFileInfo(fileName string) string {
	_, _, line, _ := runtime.Caller(1)
	return "internal/grpc/storage/postgres/auth/" + fileName + " line: " + strconv.Itoa(line)
}

func isUniqueConstraintViolation(err error) bool {
	if pgErr, ok := err.(*pq.Error); ok {
		return pgErr.Code == "23505"
	}
	return false
}
