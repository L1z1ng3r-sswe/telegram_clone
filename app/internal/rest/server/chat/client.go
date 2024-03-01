package chat_rest

import (
	models_rest "github.com/L1z1ng3r-sswe/telegram_clone/app/internal/rest/domain/models"
	"github.com/gorilla/websocket"
)

type Client struct {
	Id         int64
	Conn       *websocket.Conn
	User       models_rest.UserDB
	BIMessages chan models_rest.BIMessage
}

func (cl *Client) listen() {
	defer cl.Conn.Close()
	for {
		select {
		case msg := <-cl.BIMessages:
			if err := cl.Conn.WriteJSON(msg); err != nil {
				return
			}
		}
	}
}
