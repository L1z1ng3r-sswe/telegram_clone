package auth_rest

import (
	"fmt"
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

func (server *server) SignIn(ctx *gin.Context) {
	var user models_rest.UserSignIn
	if err := ctx.ShouldBindJSON(&user); err != nil {
		server.log.Err("Internal Server Error", err.Error(), "")
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"Internal Server Error": err.Error()})
		return
	}

	fmt.Println(user)

	tokens, userDB, err := server.service.SignIn(user, server.accessTokenExp, server.refreshTokenExp, server.secretKey)
	if err != nil {
		server.log.Err(err.ErrKey, err.ErrMsg, err.FileInfo)
		ctx.AbortWithStatusJSON(err.Code, gin.H{err.ErrKey: err.ErrMsg})
		return
	}

	server.log.Inf("user signed in", "id", userDB.Id)
	ctx.JSON(http.StatusOK, gin.H{"access_token": tokens.AccessToken, "refresh_token": tokens.RefreshToken})
}
