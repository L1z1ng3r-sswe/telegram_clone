package models_rest

type UserSignUp struct {
	Id       int64  `db:"id"`
	Email    string `db:"email"`
	Password string `db:"password"`
}

type UserSignIn struct {
	Id       int64  `db:"id"`
	Email    string `db:"email"`
	Password string `db:"password"`
}