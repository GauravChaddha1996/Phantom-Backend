package apiCommons

import "github.com/gin-gonic/gin"

type ApiErrorLogData struct {
	EndPoint  string
	UserId    string
	Timestamp string
	Message   string
	Error     error
	Data      map[string]string
}

func NewApiErrorLogData(ctx *gin.Context, message string, err error) ApiErrorLogData {
	return ApiErrorLogData{
		EndPoint:  ctx.Request.RequestURI,
		UserId:    "",
		Timestamp: "",
		Message:   message,
		Error:     err,
		Data:      make(map[string]string, 0),
	}
}
