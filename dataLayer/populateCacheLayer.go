package dataLayer

import (
	"database/sql"
	"github.com/gomodule/redigo/redis"
	"phantom/dataLayer/cacheDaos"
	"phantom/dataLayer/databasDaos"
)

func PopulateCacheLayer(db *sql.DB, pool *redis.Pool) error {
	productIdCachePopulateErr := populateAllProductIdsCache(db, pool)
	if productIdCachePopulateErr != nil {
		return productIdCachePopulateErr
	}

	categoryIdCachePopulateErr := populateAllCategoryIdsCache(db, pool)
	if categoryIdCachePopulateErr != nil {
		return categoryIdCachePopulateErr
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

func populateAllCategoryIdsCache(db *sql.DB, pool *redis.Pool) error {
	databaseDao := databasDaos.CategoryDao{DB: db}
	cacheDao := cacheDaos.AllCategoryIdsCacheDao{Pool: pool}
	cacheDelErr := cacheDao.DeleteWholeCache()
	if cacheDelErr != nil {
		return cacheDelErr
	}

	categoriesFromDb, dbErr := databaseDao.ReadAllCategories()
	if dbErr != nil {
		return dbErr
	}

	cacheSetArr := cacheDao.SetCategoryIds(categoriesFromDb)
	if cacheSetArr != nil {
		return cacheSetArr
	}
	return nil
}
