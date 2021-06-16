package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func RequestLoggerMiddleware() gin.HandlerFunc {
	return gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		var statusColor, methodColor, resetColor string
		statusColor = param.StatusCodeColor()
		methodColor = param.MethodColor()
		resetColor = param.ResetColor()
		return fmt.Sprintf("[GIN] At %v || %s%3d%s %s%s%s \"%s\" || Served in %s",
			param.TimeStamp.Format("2006/01/02 - 15:04:05"),
			statusColor, param.StatusCode, resetColor,
			methodColor, param.Method, resetColor,
			param.Path,
			param.Latency.String(),
		)
	})
}
