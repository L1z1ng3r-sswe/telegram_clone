package chat_rest

import (
	"net/http"

	models_rest "github.com/L1z1ng3r-sswe/telegram_clone/app/internal/rest/domain/models"
	"github.com/gin-gonic/gin"
)

func (server *server) SendBIMessage(ctx *gin.Context) {
	var message models_rest.BIMessage
	if err := ctx.ShouldBindJSON(&message); err != nil {
		server.log.Err("Bad Request", err.Error(), getFileInfo("send-bi-message.go"))
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Bad Request": err.Error()})
		return
	}

	// verify and handle in service and save in db

	server.hub.BroadcastBIMessages <- message
}
