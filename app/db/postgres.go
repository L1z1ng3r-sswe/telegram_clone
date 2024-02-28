package db

import "github.com/jmoiron/sqlx"

func NewPostgres(postgresPath string) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", postgresPath)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
