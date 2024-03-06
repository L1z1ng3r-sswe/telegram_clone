package chat_rest

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (server *server) GetAllMembersOfCommunity(ctx *gin.Context) {
	communityIdStr := ctx.Query("community-id")

	if communityIdStr == "" {
		server.log.Err("Bad Request", "Community ID is required", "")
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Bad Request": "Community ID is required"})
		return
	}

	communityId, err := strconv.ParseInt(communityIdStr, 10, 64)
	if err != nil {
		server.log.Err("Bad Request", "Invalid community ID", "")
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Bad Request": "Invalid community ID"})
		return
	}

	community, exist := server.Hub.Communities[communityId]
	if !exist {
		server.log.Err("Not Found", "Community not found", "")
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"Not Found": "Community not found"})
		return
	}

	// server.log.Inf("check", "all members of the community", community.Clients)
	ctx.JSON(http.StatusOK, gin.H{"all streaming members of the community": community.Clients})
}
