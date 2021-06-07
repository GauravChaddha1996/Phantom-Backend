package dataLayer

import (
	"database/sql"
	"github.com/gomodule/redigo/redis"
	"phantom/dataLayer/cacheDaos"
	"phantom/dataLayer/databasDaos"
	"phantom/dataLayer/dbModels"
)

func PopulateCacheLayer(db *sql.DB, pool *redis.Pool) error {
	productDao := databasDaos.ProductDao{DB: db}
	categoryDao := databasDaos.CategoryDao{DB: db}

	productsFromDb, dbErr := productDao.ReadAllProducts()
	if dbErr != nil {
		return dbErr
	}

	categoriesFromDb, dbErr := categoryDao.ReadAllCategories()
	if dbErr != nil {
		return dbErr
	}

	productIdCachePopulateErr := populateAllProductIdsCache(pool, productsFromDb)
	if productIdCachePopulateErr != nil {
		return productIdCachePopulateErr
	}

	categoryIdCachePopulateErr := populateAllCategoryIdsCache(db, pool)
	if categoryIdCachePopulateErr != nil {
		return categoryIdCachePopulateErr
	}

	categoryIdToProductIdPopulateErr := populateCategoryIdsToProductIdsCache(pool, productsFromDb, categoriesFromDb)
	if categoryIdToProductIdPopulateErr != nil {
		return categoryIdToProductIdPopulateErr
	}

	return nil
}

func populateAllProductIdsCache(pool *redis.Pool, productsFromDb *[]dbModels.Product) error {
	cacheDao := cacheDaos.AllProductIdsCacheDao{Pool: pool}
	cacheDelErr := cacheDao.DeleteWholeCache()
	if cacheDelErr != nil {
		return cacheDelErr
	}

	for _, product := range *productsFromDb {
		cacheSetError := cacheDao.SetProductId(&product)
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

func populateCategoryIdsToProductIdsCache(pool *redis.Pool, productsFromDb *[]dbModels.Product, categoriesFromDb *[]dbModels.Category) error {
	cacheDao := cacheDaos.CategoryIdToProductIdDao{Pool: pool}
	cacheDelErr := cacheDao.DeleteWholeCache(categoriesFromDb)
	if cacheDelErr != nil {
		return cacheDelErr
	}

	cacheSetArr := cacheDao.SetCategoryIdsToProductIdsMap(productsFromDb)
	if cacheSetArr != nil {
		return cacheSetArr
	}
	return nil
}
