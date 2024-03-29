package auth_rest

import (
	"net/http"

	tokens_rest "github.com/L1z1ng3r-sswe/telegram_clone/app/internal/rest/utils/jwt"
	"github.com/gin-gonic/gin"
)

func (server *server) IsAuthMW() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		accessToken := ctx.GetHeader("Authorization")
		if accessToken == "" {
			server.log.Err("Unauthorized", "Invalid token", "")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"Unauthorized": "token is required"})
			return
		}

		userId, err := tokens_rest.IsTokenValid(accessToken, server.secretKey)
		if err != nil {
			server.log.Err(err.ErrKey, err.ErrMsg, err.FileInfo)
			ctx.AbortWithStatusJSON(err.Code, gin.H{err.ErrKey: err.ErrMsg})
			return
		}

		ctx.Set("user_id", userId)
		ctx.Next()
	}
}
