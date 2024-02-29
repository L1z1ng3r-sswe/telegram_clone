package auth_rest

import (
	"net/http"

	models_rest "github.com/L1z1ng3r-sswe/telegram_clone/app/internal/rest/domain/models"
	"github.com/gin-gonic/gin"
)

func (server *userServerAPI) SignUp(ctx *gin.Context) {
	var user models_rest.UserSignUp
	if err := ctx.ShouldBindJSON(&user); err != nil {
		server.log.Err("Internal Server Error", err.Error(), getFileInfo("sign-up.go"))
		ctx.JSON(http.StatusInternalServerError, gin.H{"Internal Server Error": err.Error()})
		return
	}

	tokens, userDB, err, errKey, errMsg, code, fileInfo := server.service.SignUp(user, server.accessTokenExp, server.refreshTokenExp, server.secretKey)
	if err != nil {
		server.log.Err(errKey, errMsg, fileInfo)
		ctx.JSON(code, gin.H{errKey: errMsg})
		return
	}

	server.log.Inf("signed up a new user", "id", userDB.Id)
	ctx.JSON(http.StatusOK, gin.H{"access_token": tokens.AccessToken, "refresh_token": tokens.RefreshToken})
}
