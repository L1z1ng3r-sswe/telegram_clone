package chat_rest

import (
	"net/http"

	models_rest "github.com/L1z1ng3r-sswe/telegram_clone/app/internal/rest/domain/models"
	"github.com/gin-gonic/gin"
)

func (server *server) CreateCommunity(ctx *gin.Context) {
	var community models_rest.Community

	if err := ctx.ShouldBindJSON(&community); err != nil {
		server.log.Err("Bad Request", err.Error(), "")
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Bad Request": err.Error()})
		return
	}

	// get user-id from access-token
	userId, exist := ctx.Get("user_id")
	if !exist {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"Unauthorized": "user id not found"})
		server.log.Err("Unauthorized", "user_Id not found", "")
		return
	}

	community.OwnerId = userId.(int64)

	// verify that the client is exist in hub
	cl, exist := server.Hub.Clients[community.OwnerId]
	if !exist {
		server.log.Err("Bad Request", "Client doest not exist", "")
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Bad Request": "Client does not exist"})
		return
	}

	// save a new community in the db
	communityDB, err := server.service.CreateCommunity(community)
	if err != nil {
		server.log.Err(err.ErrKey, err.ErrMsg, err.FileInfo)
		ctx.AbortWithStatusJSON(err.Code, gin.H{err.ErrKey: err.ErrMsg})
		return
	}

	// create the community's clients
	communityDB.Clients = make(map[int64]*models_rest.Client)

	// add the first one (owner)
	communityDB.Clients[cl.UserId] = cl

	// add new community to all communties
	server.Hub.Communities[communityDB.Id] = communityDB
	server.log.Inf("New community created", "community", server.Hub.Communities)
	ctx.JSON(http.StatusOK, gin.H{"New community": server.Hub.Communities[communityDB.Id]})
}
