package chat_rest

import models_rest "github.com/L1z1ng3r-sswe/telegram_clone/app/internal/rest/domain/models"

type Hub struct {
	BroadcastBIMessages chan models_rest.BIMessage
	BIChats             map[int64]models_rest.BIChats
	Clients             map[int64]*Client
}

func (h *Hub) Run() {
	for {
		select {
		// from bi-chats:
		case msg := <-h.BroadcastBIMessages:
			for {
				chat := h.BIChats[msg.ChatId]

				// get chats with them
				firstClient := h.Clients[chat.FirstUserId]
				secondClient := h.Clients[chat.SecondUserId]

				// send message to clients
				firstClient.BIMessages <- msg
				secondClient.BIMessages <- msg
			}
		}
	}
}
