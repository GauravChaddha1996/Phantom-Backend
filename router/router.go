package router

import (
	"github.com/gin-gonic/gin"
)

func Initialize() (*gin.Engine, error) {
	gin.SetMode(gin.ReleaseMode)

	router := gin.New()
	router.Use(gin.Recovery())
	router.Use(LoggerMiddleware())
	router.Use(TimeoutMiddleware())
	return router, nil
}
