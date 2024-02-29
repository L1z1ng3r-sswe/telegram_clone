package models_rest

import "time"

type Message struct {
	Id         int64     `db:"id"`
	ChatId     int64     `db:"chat_id"`
	SenderId   int64     `db:"sender_id"`
	ReceiverId int64     `db:"receiver_id"`
	Content    string    `db:"content"`
	TimeStamp  time.Time `db:"timestamp"`
}

type Chat struct {
	Id           int64 `db:"id"`
	FirstUserId  int64 `db:"first_user_id"`
	SecondUserId int64 `db:"second_user_id"`
	Messages     chan Message
}
