package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gomodule/redigo/redis"
	"log"
	"phantom/config"
	"phantom/dataLayer"
	"phantom/router"
	"time"
)

func main() {

	// Read config struct
	conf := config.ReadConfig()
	log.Println(conf)

	// Open sql database conection
	sqlDb := openSqlDB(conf)
	defer func(db *sql.DB) {
		_ = db.Close()
	}(sqlDb)

	// Open redis cache pool
	redisCachePool := openRedisCachePool(conf)
	defer func(pool *redis.Pool) {
		_ = pool.Close()
	}(redisCachePool)

	// Pre-populate redis cache pool
	cachePopulateErr := dataLayer.PopulateCacheLayer(sqlDb, redisCachePool)
	if cachePopulateErr != nil {
		log.Fatal(cachePopulateErr)
		return
	}

	// Initialize router
	_, routerInitErr := router.Initialize(conf)
	if routerInitErr != nil {
		log.Fatal(routerInitErr)
		return
	}
}

func openSqlDB(config *config.Config) *sql.DB {
	dbConfig := config.Database
	dataSourceName := fmt.Sprintf("%s:@%s(%s:%s)/%s", dbConfig.Username, dbConfig.Network, dbConfig.Host, dbConfig.Port, dbConfig.Name)
	db, dbOpenErr := sql.Open(dbConfig.Driver, dataSourceName)
	if dbOpenErr != nil {
		panic(dbOpenErr)
	}
	return db
}

func openRedisCachePool(config *config.Config) *redis.Pool {
	cacheConfig := config.Cache
	address := fmt.Sprintf("%s:%d", cacheConfig.Host, cacheConfig.Port)
	return &redis.Pool{
		Dial: func() (redis.Conn, error) {
			return redis.Dial(cacheConfig.Network, address)
		},
		MaxIdle:     cacheConfig.MaxIdle,
		IdleTimeout: time.Duration(cacheConfig.IdleTimeout) * time.Second,
	}
}
