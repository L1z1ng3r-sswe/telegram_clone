package chat_rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (server *server) GetAllConnections(ctx *gin.Context) {
	// server.log.Inf("check", "all streaming clients", server.Hub.Clients)
	ctx.JSON(http.StatusOK, gin.H{"all streaming clients": server.Hub.Clients})
}
