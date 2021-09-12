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
	propertyDao := databaseDaos.PropertySqlDao{DB: db}
	propertyValueDao := databaseDaos.PropertyValueSqlDao{DB: db}
	productToPropertyDao := databaseDaos.ProductToPropertySqlDao{DB: db}
	categoryToPropertyDao := databaseDaos.CategoryToPropertySqlDao{DB: db}

	// Read from Database
	var dbReadError error
	productsFromDb, productReadErr := productDao.ReadAllProducts()
	categoriesFromDb, categoryReadErr := categoryDao.ReadAllCategories()
	propertyArr, propertyReadErr := propertyDao.ReadAllProperty()
	propertyValuesArr, propertyValueReadErr := propertyValueDao.ReadAllPropertyValues()
	productToPropertyArr, productToPropertyReadErr := productToPropertyDao.ReadAllProductToPropertyMapping()
	categoryToPropertyArr, categoryToProductReadErr := categoryToPropertyDao.ReadAllCategoryToPropertyMapping()

	// Handle any db error
	dbReadError = multierr.Combine(
		productReadErr,
		categoryReadErr,
		propertyReadErr,
		propertyValueReadErr,
		productToPropertyReadErr,
		categoryToProductReadErr,
	)
	if dbReadError != nil {
		return productReadErr
	}

	// Populate cache from db results
	productIdCachePopulateErr := populateAllProductIdsCache(pool, productsFromDb)
	categoryIdCachePopulateErr := populateAllCategoryIdsCache(pool, categoriesFromDb)
	categoryIdToProductIdPopulateErr := populateCategoryIdsToProductIdsCache(pool, productsFromDb, categoriesFromDb)
	propertyValueIdToProductIdPopulateErr := populatePropertyValueIdsToProductIdsCache(pool, propertyValuesArr, productToPropertyArr)
	categoryIdToPropertyIdPopulateErr := populateCategoryIdToPropertyIdsCache(pool, categoriesFromDb, categoryToPropertyArr)
	propertyIdToPropertyValueIdPopulateErr := populatePropertyIdToPropertyValueIdCache(pool, propertyArr, propertyValuesArr)
	productIdToPropertyValueIdPopulateErr := populateProductIdsToPropertyValueIdsCache(pool, productsFromDb, productToPropertyArr)
	propertyValueIdToPropertyIdPopulateErr := populateProductValueIdsToPropertyIdsCache(pool, propertyValuesArr)

	// Handle any cache populate error
	cachePopulateErr := multierr.Combine(
		productIdCachePopulateErr,
		categoryIdCachePopulateErr,
		categoryIdToProductIdPopulateErr,
		propertyValueIdToProductIdPopulateErr,
		categoryIdToPropertyIdPopulateErr,
		propertyIdToPropertyValueIdPopulateErr,
		productIdToPropertyValueIdPopulateErr,
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

func populatePropertyValueIdsToProductIdsCache(
	pool *redis.Pool,
	propertyValueArr *[]dbModels.PropertyValue,
	productToPropertyArr *[]dbModels.ProductToProperty,
) error {
	cacheDao := cacheDaos.PropertyValueToProductRedisDao{Pool: pool}
	cacheDelErr := cacheDao.DeleteWholeCache(propertyValueArr)
	if cacheDelErr != nil {
		return cacheDelErr
	}

	cacheSetArr := cacheDao.SetPropertyValuesToProductIds(productToPropertyArr)
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

func populatePropertyIdToPropertyValueIdCache(
	pool *redis.Pool,
	propertyArr *[]dbModels.Property,
	propertyValueArr *[]dbModels.PropertyValue,
) error {
	cacheDao := cacheDaos.PropertyIdToPropertyValueIdRedisDao{Pool: pool}
	cacheDelErr := cacheDao.DeleteWholeCache(propertyArr)
	if cacheDelErr != nil {
		return cacheDelErr
	}

	cacheSetArr := cacheDao.SetPropertyIdToPropertyValueIdCache(propertyValueArr)
	if cacheSetArr != nil {
		return cacheSetArr
	}
	return nil
}

func populateProductIdsToPropertyValueIdsCache(
	pool *redis.Pool,
	products *[]dbModels.Product,
	productToPropertyArr *[]dbModels.ProductToProperty,
) error {
	cacheDao := cacheDaos.ProductToPropertyValueRedisDao{Pool: pool}
	cacheDelErr := cacheDao.DeleteWholeCache(products)
	if cacheDelErr != nil {
		return cacheDelErr
	}

	cacheSetArr := cacheDao.SetProductIdsToPropertyValues(productToPropertyArr)
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
	cacheDelErr := cacheDao.DeleteWholeCache(propertyValueArr)
	if cacheDelErr != nil {
		return cacheDelErr
	}

	cacheSetArr := cacheDao.SetPropertyValueIdToPropertyIdCache(propertyValueArr)
	if cacheSetArr != nil {
		return cacheSetArr
	}
	return nil
}
