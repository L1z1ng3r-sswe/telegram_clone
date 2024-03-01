package models_rest

type UserSignUp struct {
	Id       int64  `db:"id" json:"id"`
	Email    string `db:"email" json:"email"`
	Password string `db:"password" json:"password"`
}

type UserSignIn struct {
	Id       int64  `db:"id" json:"id"`
	Email    string `db:"email" json:"email"`
	Password string `db:"password" json:"password"`
}

type UserDB struct {
	Id       int64  `db:"id" json:"id"`
	Email    string `db:"email" json:"email"`
	Password string `db:"password" json:"password"`
}
