package chat_rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (server *server) GetAllBIMessages(ctx *gin.Context) {
	biChatId := ctx.Param("bi-chat-id")

	// save in db
	messages, err, errKey, errMsg, code, fileInfo := server.service.GetAllBIMessages(biChatId)
	if err != nil {
		server.log.Err(errKey, errMsg, fileInfo)
		ctx.AbortWithStatusJSON(code, gin.H{errKey: errMsg})
		return
	}

	server.log.Inf("all messages getted", "messages", messages)
	ctx.JSON(http.StatusOK, gin.H{"messages": messages})
}
