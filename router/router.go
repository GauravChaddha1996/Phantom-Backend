package router

import (
	"github.com/gin-gonic/gin"
	"phantom/config"
)

func Initialize(*config.Config) (*gin.Engine, error) {
	gin.SetMode(gin.ReleaseMode)

	router := gin.New()
	router.Use(gin.Recovery())
	router.Use(RequestLoggerMiddleware())

	return router, nil
}
