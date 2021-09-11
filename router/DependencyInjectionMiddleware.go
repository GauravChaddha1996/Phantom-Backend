package router

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/gomodule/redigo/redis"
)

const REDIS_POOL = "redis_pool"
const SQL_DB = "sql_db"

func DependencyInjectionMiddleware(redisPool *redis.Pool, sqlDB *sql.DB) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		ctx.Set(REDIS_POOL, redisPool)
		ctx.Set(SQL_DB, sqlDB)
		ctx.Next()
	}
}
