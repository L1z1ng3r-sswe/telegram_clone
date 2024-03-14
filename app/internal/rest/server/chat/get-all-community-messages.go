package chat_rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (server *server) GetAllCommunityMessages(ctx *gin.Context) {

	communityId := ctx.Param("community-id")

	// save to the db
	messages, err := server.service.GetAllCommunityMessages(communityId)
	if err != nil {
		server.log.Err(err.ErrKey, err.ErrMsg, err.FileInfo)
		ctx.AbortWithStatusJSON(err.Code, gin.H{err.ErrKey: err.ErrMsg})
		return
	}

	server.log.Inf("all messages getted", "messages", messages)
	ctx.JSON(http.StatusOK, gin.H{"messages": messages})
}
