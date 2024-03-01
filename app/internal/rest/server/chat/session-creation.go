package chat_rest

import (
	"net/http"

	models_rest "github.com/L1z1ng3r-sswe/telegram_clone/app/internal/rest/domain/models"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func (server *server) SessionCreation(ctx *gin.Context) {
	var user models_rest.UserSignIn
	if err := ctx.ShouldBindJSON(&user); err != nil {
		server.log.Err("Internal Server Error", err.Error(), getFileInfo("session-creation.go"))
		ctx.JSON(http.StatusInternalServerError, gin.H{"Internal Server Error": err.Error()})
		return
	}

	conn, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		conn.Close()
		server.log.Err("Internal Server Error", err.Error(), getFileInfo("session-creation.go"))
		ctx.JSON(http.StatusInternalServerError, gin.H{"Internal Server Error": err.Error()})
		return
	}

	tokens, userDB, err, errKey, errMsg, code, fileInfo := server.service.SessionCreation(user, server.accessTokenExp, server.refreshTokenExp, server.secretKey)
	if err != nil {
		server.log.Err(errKey, errMsg, fileInfo)
		ctx.JSON(code, gin.H{errKey: errMsg})
		return
	}

	cl := Client{
		Id:         userDB.Id,
		Conn:       conn,
		BIMessages: make(chan models_rest.BIMessage, 5),
		User:       userDB,
	}

	go cl.listen()

	server.log.Inf("user signed in", "id", userDB.Id)
	ctx.JSON(http.StatusOK, gin.H{"access_token": tokens.AccessToken, "refresh_token": tokens.RefreshToken})
}
