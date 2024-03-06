package chat_rest

import (
	"net/http"

	models_rest "github.com/L1z1ng3r-sswe/telegram_clone/app/internal/rest/domain/models"
	"github.com/gin-gonic/gin"
)

func (server *server) CreateBIChat(ctx *gin.Context) {
	var chat models_rest.BIChat

	if err := ctx.ShouldBindJSON(&chat); err != nil {
		server.log.Err("Bad Request", err.Error(), "")
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Bad Request": err.Error()})
		return
	}

	// verify that clients exist
	if _, exist := server.Hub.Clients[chat.FirstUserId]; !exist {
		server.log.Err("Bad Request", "Client doest not exist", "")
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Bad Request": "Client does not exist"})
		return
	} else if _, exist := server.Hub.Clients[chat.SecondUserId]; !exist {
		server.log.Err("Bad Request", "Client doest not exist", "")
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Bad Request": "Client does not exist"})
		return
	}

	chatDB, err, errKey, errMsg, code, fileInfo := server.service.CreateBIChat(chat)
	if err != nil {
		server.log.Err(errKey, errMsg, fileInfo)
		ctx.AbortWithStatusJSON(code, gin.H{errKey: errMsg})
		return
	}

	server.log.Inf("New chat", "chat", chatDB)
	ctx.JSON(http.StatusOK, gin.H{"New chat": chatDB})
}
