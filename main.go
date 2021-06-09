package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gomodule/redigo/redis"
	"log"
	"phantom/dataLayer"
	"phantom/dataLayer/cacheDaos"
	"time"
)

func main() {
	db := openDB()
	defer db.Close()
	pool := openCachePool()
	defer pool.Close()

	err := dataLayer.PopulateCacheLayer(db, pool)
	if err != nil {
		log.Fatal(err)
		return
	}

	cacheDao := cacheDaos.CategoryIdToPropertyIdDao{Pool: pool}
	propertyIds, cacheReadErr := cacheDao.ReadPropertyIdsForCategoryId(1)
	if cacheReadErr != nil {
		log.Fatal(cacheReadErr)
		return
	}
	log.Println(propertyIds)
}

func openDB() *sql.DB {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/phantom")
	if err != nil {
		panic(err)
	}
	return db
}

func openCachePool() *redis.Pool {
	return &redis.Pool{
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", "localhost:6379")
		},
		MaxIdle:     10,
		IdleTimeout: 120 * time.Second,
	}
}
