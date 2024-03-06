package models_rest

import "github.com/gorilla/websocket"

type Message struct {
	Type     string `json:"type"`
	Id       int64  `json:"id" db:"id"`
	Content  string `json:"content" db:"content"`
	SenderId int64  `json:"sender_id" db:"sender_id"`

	// if message bi-dirictional
	ReceiverId int64 `json:"receiver_id" db:"receiver_id"`
	BIChatId   int64 `json:"bi_chat_id" db:"bi_chat_id"`

	// if message from community:
	CommunityId int64 `json:"community_id" db:"community_id"`
}

type Client struct {
	UserId       int64           `json:"userId"`
	Conn         *websocket.Conn `json:"-"`
	BIMessages   chan Message    `json:"-"`
	CommMessages chan Message    `json:"-"`
}

type BIChat struct {
	Id           string `json:"id" db:"id"`
	FirstUserId  int64  `json:"first_user_id" db:"first_user_id"`
	SecondUserId int64  `json:"second_user_id" db:"second_user_id"`
}

type Community struct {
	Id      int64             `json:"id" db:"id"`
	Name    string            `json:"name" db:"name"`
	OwnerId int64             `json:"owner_id" db:"owner_id"`
	Clients map[int64]*Client `json:"-" db:"-"`
}

type CommunityMember struct {
	Id          string `json:"id" db:"id"`
	UserId      int64  `json:"user_id" db:"user_id"`
	CommunityId int64  `json:"community_id" db:"community_id"`
}
