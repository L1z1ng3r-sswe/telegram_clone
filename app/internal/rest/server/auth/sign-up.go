package auth_rest

import (
	"net/http"

	models_rest "github.com/L1z1ng3r-sswe/telegram_clone/app/internal/rest/domain/models"
	"github.com/gin-gonic/gin"
)

func (server *server) SignUp(ctx *gin.Context) {
	var user models_rest.UserSignUp
	if err := ctx.ShouldBindJSON(&user); err != nil {
		server.log.Err("Internal Server Error", err.Error(), "")
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"Internal Server Error": err.Error()})
		return
	}

	tokens, userDB, err := server.service.SignUp(user, server.accessTokenExp, server.refreshTokenExp, server.secretKey)
	if err != nil {
		server.log.Err(err.ErrKey, err.ErrMsg, err.FileInfo)
		ctx.AbortWithStatusJSON(err.Code, gin.H{err.ErrKey: err.ErrMsg})
		return
	}

	server.log.Inf("signed up a new user", "id", userDB.Id)
	ctx.JSON(http.StatusOK, gin.H{"access_token": tokens.AccessToken, "refresh_token": tokens.RefreshToken})
}
