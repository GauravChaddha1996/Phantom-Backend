package router

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/gomodule/redigo/redis"
)

func Initialize(redisPool *redis.Pool, sqlDb *sql.DB) (*gin.Engine, error) {
	gin.SetMode(gin.ReleaseMode)

	router := gin.New()
	router.Use(gin.Recovery())
	router.Use(LoggerMiddleware())
	router.Use(TimeoutMiddleware())
	router.Use(DependencyInjectionMiddleware(redisPool, sqlDb))
	return router, nil
}
