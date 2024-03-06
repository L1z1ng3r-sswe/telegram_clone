package chat_rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (server *server) GetAllCommunity(ctx *gin.Context) {
	// server.log.Inf("check", "all streaming communities", server.Hub.Communities)
	ctx.JSON(http.StatusOK, gin.H{"all streaming communities": server.Hub.Communities})
}
