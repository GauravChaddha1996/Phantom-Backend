package ginRouter

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func TimeoutMiddleware() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		newCtx, cancel := context.WithTimeout(ctx.Request.Context(), 1*time.Second)
		defer func() {
			if newCtx.Err() == context.DeadlineExceeded {
				ctx.AbortWithStatus(http.StatusGatewayTimeout)
			}
			cancel()
		}()

		ctx.Request = ctx.Request.WithContext(newCtx)
		ctx.Next()
	}
}
