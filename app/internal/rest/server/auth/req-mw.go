package auth_rest

import (
	"time"

	"github.com/gin-gonic/gin"
)

func (server *server) RequesMW() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		startPoint := time.Now()
		ctx.Next()

		server.log.Request(ctx.Request.Method, ctx.Request.URL.Path, ctx.Writer.Status(), time.Since(startPoint))
	}
}
