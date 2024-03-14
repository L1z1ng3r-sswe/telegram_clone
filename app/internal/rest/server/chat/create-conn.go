package chat_rest

import (
	"net/http"

	models_rest "github.com/L1z1ng3r-sswe/telegram_clone/app/internal/rest/domain/models"
	tokens_rest "github.com/L1z1ng3r-sswe/telegram_clone/app/internal/rest/utils/jwt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

func (server *server) CreateConn(ctx *gin.Context) {
	conn, err := server.upgrader.Upgrade(ctx.Writer, ctx.Request, nil)

	if err != nil {
		server.log.Err("Internal Server Error", err.Error(), "")
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"Internal Server Error": err.Error()})
		return
	}

	queryParams := ctx.Request.URL.Query()
	accessToken := queryParams.Get("access_token")

	if accessToken == "" {
		server.log.Err("Unauthorized", "Invalid Token", "")
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"Unauthorized": "Invalid Token"})
		return
	}

	userId, er := tokens_rest.IsTokenValid(accessToken, server.secretKey)
	if err != nil {
		server.log.Err(er.ErrKey, er.ErrMsg, er.FileInfo)
		ctx.AbortWithStatusJSON(er.Code, gin.H{er.ErrKey: er.ErrMsg})
		return
	}

	cl := &models_rest.Client{
		UserId:       userId,
		Conn:         conn,
		BIMessages:   make(chan models_rest.Message, 5),
		CommMessages: make(chan models_rest.Message, 5),
	}

	server.Hub.Clients[cl.UserId] = cl

	server.log.Inf("New connection", "user_id", cl.UserId)

	go server.handleMessages(cl)
}

func (server *server) handleMessages(cl *models_rest.Client) {
	// receive messages from users, channels-broadcast, communities-broadcast
	go server.receiveMessages(cl)

	for {
		var msg models_rest.Message

		// read message from postman-send
		if err := cl.Conn.ReadJSON(&msg); err != nil {
			if _, ok := err.(*websocket.CloseError); ok {
				delete(server.Hub.Clients, cl.UserId)
				server.log.Inf("Client closed the connection", "id", cl.UserId)
				return
			}
		}

		server.handleSendedMessage(cl, msg)
	}

}

func (server *server) receiveMessages(cl *models_rest.Client) {
	for {
		select {
		// from users
		case msg := <-cl.BIMessages:
			server.log.Inf("Received bi-message", "msg", msg)

			// from communities
		case msg := <-cl.CommMessages:
			cl.Conn.WriteJSON(msg)
		}
	}
}

func (server *server) handleSendedMessage(cl *models_rest.Client, msg models_rest.Message) {

	// Type assertion to determine the type of message
	switch msg.Type {
	//!  if bi-directional message
	case "bi-message":
		if cl.UserId == msg.SenderId {
			receiverCl, exist := server.Hub.Clients[msg.ReceiverId]
			if !exist {
				return
			}

			// save in db
			msgDB, err := server.service.CreateBIMessage(msg)
			if err != nil {
				server.log.Err(err.ErrKey, err.ErrMsg, err.FileInfo)
			}

			// send to postman
			receiverCl.Conn.WriteJSON(msgDB)

			// send to broadcast
			receiverCl.BIMessages <- msgDB
		}

		server.log.Inf("Received a new message from frontend", "msg", msg)

	//!  if community message
	case "community-message":
		if cl.UserId == msg.SenderId {

			// save in db
			msgDB, err := server.service.CreateCommunityMessage(msg)
			if err != nil {
				server.log.Err(err.ErrKey, err.ErrMsg, err.FileInfo)
			}

			server.Hub.BroadcastComm <- msgDB
		}

		server.log.Inf("Received a new message from frontend community", "msg", msg)

	//!  handle error
	default:
		server.log.Err("Bad Request", "Unkown message type. msg", "")
	}

}
