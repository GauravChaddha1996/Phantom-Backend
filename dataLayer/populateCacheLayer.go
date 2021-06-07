package dataLayer

import (
	"database/sql"
	"github.com/gomodule/redigo/redis"
	"phantom/dataLayer/cacheDaos"
	"phantom/dataLayer/databasDaos"
)

func PopulateCacheLayer(db *sql.DB, pool *redis.Pool) error {
	err := populateAllProductIdsCache(db, pool)
	if err != nil {
		return err
	}
	return nil
}

func populateAllProductIdsCache(db *sql.DB, pool *redis.Pool) error {
	databaseDao := databasDaos.ProductDao{DB: db}
	cacheDao := cacheDaos.AllProductIdsCacheDao{Pool: pool}
	cacheDelErr := cacheDao.DeleteWholeCache()
	if cacheDelErr != nil {
		return cacheDelErr
	}

	productsFromDb, dbErr := databaseDao.ReadAllProducts()
	if dbErr != nil {
		return dbErr
	}

	for _, product := range productsFromDb {
		cacheSetError := cacheDao.SetProductId(product)
		if cacheSetError != nil {
			return cacheSetError
		}
	}

	return nil
}
