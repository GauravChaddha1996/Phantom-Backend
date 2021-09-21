package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gomodule/redigo/redis"
	"log"
	"phantom/apis/filter"
	"phantom/apis/home"
	"phantom/config"
	"phantom/dataLayer"
	"phantom/ginRouter"
	"time"
)

func main() {

	// Read config struct
	envConfig := config.ReadEnvConfig()

	// Open sql database connection
	sqlDb := openSqlDB(envConfig)
	defer func(db *sql.DB) {
		dbCloseErr := db.Close()
		if dbCloseErr != nil {
			log.Fatal(dbCloseErr)
		}
	}(sqlDb)

	// Open redis cache pool
	redisCachePool := openRedisCachePool(envConfig)
	defer func(pool *redis.Pool) {
		redisCloseErr := pool.Close()
		if redisCloseErr != nil {
			log.Fatal(redisCloseErr)
		}
	}(redisCachePool)

	// Pre-populate redis cache pool
	cachePopulateErr := dataLayer.PopulateCacheLayer(sqlDb, redisCachePool)
	if cachePopulateErr != nil {
		log.Fatal(cachePopulateErr)
		return
	}

	// Initialize ginRouter
	router, routerInitErr := ginRouter.Initialize(redisCachePool, sqlDb)
	if routerInitErr != nil {
		log.Fatal(routerInitErr)
		return
	}

	// Register routes in router
	router.GET("/home", home.ApiHandler)
	router.GET("/filter", filter.ApiHandler)

	// Start router
	err := router.Run()
	if err != nil {
		return
	}
}

func openSqlDB(envConfig *config.EnvConfig) *sql.DB {
	dbConfig := envConfig.Database
	dataSourceName := fmt.Sprintf("%s:@%s(%s:%s)/%s?parseTime=true",
		dbConfig.Username, dbConfig.Network, dbConfig.Host, dbConfig.Port, dbConfig.Name)
	db, dbOpenErr := sql.Open(dbConfig.Driver, dataSourceName)
	if dbOpenErr != nil {
		panic(dbOpenErr)
	}
	return db
}

func openRedisCachePool(envConfig *config.EnvConfig) *redis.Pool {
	cacheConfig := envConfig.Cache
	address := fmt.Sprintf("%s:%d", cacheConfig.Host, cacheConfig.Port)
	return &redis.Pool{
		Dial: func() (redis.Conn, error) {
			return redis.Dial(cacheConfig.Network, address)
		},
		MaxIdle:     cacheConfig.MaxIdle,
		IdleTimeout: time.Duration(cacheConfig.IdleTimeout) * time.Second,
	}
}
