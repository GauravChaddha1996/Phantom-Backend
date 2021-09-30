package dataLayer

import (
	"database/sql"
	"github.com/gomodule/redigo/redis"
	"go.uber.org/multierr"
	"phantom/dataLayer/cacheDaos"
	"phantom/dataLayer/databaseDaos"
	"phantom/dataLayer/dbModels"
)

func PopulateCacheLayer(
	db *sql.DB,
	pool *redis.Pool,
) error {

	// Make DAOs
	productDao := databaseDaos.ProductSqlDao{DB: db}
	categoryDao := databaseDaos.CategorySqlDao{DB: db}
	propertyValueDao := databaseDaos.PropertyValueSqlDao{DB: db}
	categoryToPropertyDao := databaseDaos.CategoryToPropertySqlDao{DB: db}

	// Read from Database
	var dbReadError error
	productsFromDb, productReadErr := productDao.ReadAllProducts()
	categoriesFromDb, categoryReadErr := categoryDao.ReadAllCategories()
	propertyValuesArr, propertyValueReadErr := propertyValueDao.ReadAllPropertyValues()
	categoryToPropertyArr, categoryToProductReadErr := categoryToPropertyDao.ReadAllCategoryToPropertyMapping()

	// Handle any db error
	dbReadError = multierr.Combine(
		productReadErr,
		categoryReadErr,
		propertyValueReadErr,
		categoryToProductReadErr,
	)
	if dbReadError != nil {
		return productReadErr
	}

	// Populate cache from db results
	productIdCachePopulateErr := populateAllProductIdsCache(pool, productsFromDb)
	categoryIdCachePopulateErr := populateAllCategoryIdsCache(pool, categoriesFromDb)
	categoryIdToProductIdPopulateErr := populateCategoryIdsToProductIdsCache(pool, productsFromDb, categoriesFromDb)
	categoryIdToPropertyIdPopulateErr := populateCategoryIdToPropertyIdsCache(pool, categoriesFromDb, categoryToPropertyArr)
	propertyValueIdToPropertyIdPopulateErr := populateProductValueIdsToPropertyIdsCache(pool, propertyValuesArr)

	// Handle any cache populate error
	cachePopulateErr := multierr.Combine(
		productIdCachePopulateErr,
		categoryIdCachePopulateErr,
		categoryIdToProductIdPopulateErr,
		categoryIdToPropertyIdPopulateErr,
		propertyValueIdToPropertyIdPopulateErr,
	)

	if cachePopulateErr != nil {
		return cachePopulateErr
	}
	return nil
}

func populateAllProductIdsCache(
	pool *redis.Pool,
	productsFromDb *[]dbModels.Product,
) error {
	cacheDao := cacheDaos.AllProductIdsRedisDao{Pool: pool}
	cacheDelErr := cacheDao.DeleteWholeCache()
	if cacheDelErr != nil {
		return cacheDelErr
	}

	cacheSetError := cacheDao.SetProductIdsCache(productsFromDb)
	if cacheSetError != nil {
		return cacheSetError
	}

	return nil
}

func populateAllCategoryIdsCache(
	pool *redis.Pool,
	categoriesFromDb *[]dbModels.Category,
) error {
	cacheDao := cacheDaos.AllCategoryIdsRedisDao{Pool: pool}
	cacheDelErr := cacheDao.DeleteWholeCache()
	if cacheDelErr != nil {
		return cacheDelErr
	}

	cacheSetArr := cacheDao.SetCategoryIds(categoriesFromDb)
	if cacheSetArr != nil {
		return cacheSetArr
	}
	return nil
}

func populateCategoryIdsToProductIdsCache(
	pool *redis.Pool,
	productsFromDb *[]dbModels.Product,
	categoriesFromDb *[]dbModels.Category,
) error {
	cacheDao := cacheDaos.CategoryIdToProductIdRedisDao{Pool: pool}
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

func populateCategoryIdToPropertyIdsCache(
	pool *redis.Pool,
	categoriesFromDb *[]dbModels.Category,
	categoryToPropertyFromDbArr *[]dbModels.CategoryToProperty,
) error {
	cacheDao := cacheDaos.CategoryIdToPropertyIdRedisDao{Pool: pool}
	cacheDelErr := cacheDao.DeleteWholeCache(categoriesFromDb)
	if cacheDelErr != nil {
		return cacheDelErr
	}

	cacheSetArr := cacheDao.SetCategoryToPropertyCache(categoryToPropertyFromDbArr)
	if cacheSetArr != nil {
		return cacheSetArr
	}
	return nil
}

func populateProductValueIdsToPropertyIdsCache(
	pool *redis.Pool,
	propertyValueArr *[]dbModels.PropertyValue,
) error {
	cacheDao := cacheDaos.PropertyValueIdToPropertyIdRedisDao{Pool: pool}
	cacheDelErr := cacheDao.DeleteWholeCache()
	if cacheDelErr != nil {
		return cacheDelErr
	}

	cacheSetArr := cacheDao.SetPropertyValueIdsCache(propertyValueArr)
	if cacheSetArr != nil {
		return cacheSetArr
	}
	return nil
}
