package chat_rest

import (
	"net/http"

	models_rest "github.com/L1z1ng3r-sswe/telegram_clone/app/internal/rest/domain/models"
	"github.com/gin-gonic/gin"
)

func (server *server) JoinCommunity(ctx *gin.Context) {
	var communityMember models_rest.CommunityMember

	if err := ctx.ShouldBindJSON(&communityMember); err != nil {
		server.log.Err("Bad Request", err.Error(), "")
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Bad Request": err.Error()})
		return
	}

	// get a user-id from access-token
	userId, exist := ctx.Get("user_id")
	if !exist {
		server.log.Err("Unauthorized", "user_Id not found", "")
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"Unauthorized": "user id not found"})
		return
	}
	communityMember.UserId = userId.(int64)

	// verify that the client is already in stream
	cl, exist := server.Hub.Clients[communityMember.UserId]
	if !exist {
		server.log.Err("Bad Request", "Client doest not exist", "")
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Bad Request": "Client does not exist"})
		return
	}

	// verify that the community is already in stream
	community, exist := server.Hub.Communities[communityMember.CommunityId]
	if !exist {
		server.log.Err("Bad Request", "Community does not exist", "")
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Bad Request": "Client does not exist"})
		return
	}

	// save in a db a new member
	communityMemberDB, err, errKey, errMsg, code, fileInfo := server.service.JoinCommunity(communityMember)
	if err != nil {
		server.log.Err(errKey, errMsg, fileInfo)
		ctx.AbortWithStatusJSON(code, gin.H{errKey: errMsg})
		return
	}

	// add the client to the community
	community.Clients[cl.UserId] = cl

	server.log.Inf("User Joined a community", "comm-member", communityMemberDB)
	ctx.JSON(http.StatusOK, gin.H{"new member": communityMemberDB})
}
