package apiCommons

import (
	"github.com/gin-gonic/gin"
	"time"
)

type ApiErrorLogData struct {
	EndPoint  string            `json:"end_point,omitempty"`
	UserId    string            `json:"user_id,omitempty"`
	Timestamp time.Time         `json:"timestamp"`
	Message   string            `json:"message,omitempty"`
	Error     string            `json:"error,omitempty"`
	Data      map[string]string `json:"data,omitempty"`
}

func NewApiErrorLogData(ctx *gin.Context, message string, err error) ApiErrorLogData {
	return ApiErrorLogData{
		EndPoint:  ctx.Request.RequestURI,
		UserId:    "",
		Timestamp: time.Now(),
		Message:   message,
		Error:     err.Error(),
		Data:      make(map[string]string, 0),
	}
}
