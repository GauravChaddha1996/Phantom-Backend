package dataLayer

import (
	"database/sql"
	"github.com/gomodule/redigo/redis"
	"phantom/dataLayer/cacheDaos"
	"phantom/dataLayer/databasDaos"
	"phantom/dataLayer/dbModels"
)

func PopulateCacheLayer(db *sql.DB, pool *redis.Pool) error {

	productDao, categoryDao, propertyDao, propertyValueDao,
		productToPropertyDao, categoryToPropertyDao := createDatabaseDaos(db)

	productsFromDb, categoriesFromDb, propertyValuesArr,
		propertyArr, productToPropertyArr, categoryToPropertyArr,
		readFromDbErr := readDataFromDatabase(productDao, categoryDao, propertyValueDao,
		propertyDao, productToPropertyDao, categoryToPropertyDao)

	if readFromDbErr != nil {
		return readFromDbErr
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

	propertyValueIdToProductIdPopulateErr := populatePropertyValueIdsToProductIdsCache(pool, propertyValuesArr, productToPropertyArr)
	if propertyValueIdToProductIdPopulateErr != nil {
		return propertyValueIdToProductIdPopulateErr
	}

	categoryIdToPropertyIdPopulateErr := populateCategoryIdToPropertyIdsCache(pool, categoriesFromDb, categoryToPropertyArr)
	if categoryIdToPropertyIdPopulateErr != nil {
		return categoryIdToPropertyIdPopulateErr
	}

	propertyIdToPropertyValueIdPopulateErr := populatePropertyIdToPropertyValueIdCache(pool, propertyArr, propertyValuesArr)
	if propertyIdToPropertyValueIdPopulateErr != nil {
		return propertyIdToPropertyValueIdPopulateErr
	}

	productIdToPropertyValueIdPopulateErr := populateProductIdsToPropertyValueIdsCache(pool, productsFromDb, productToPropertyArr)
	if productIdToPropertyValueIdPopulateErr != nil {
		return productIdToPropertyValueIdPopulateErr
	}

	propertyValueIdToPropertyIdPopulateErr := populateProductValueIdsToPropertyIdsCache(pool, propertyValuesArr)
	if propertyValueIdToPropertyIdPopulateErr != nil {
		return propertyValueIdToPropertyIdPopulateErr
	}

	return nil
}

func readDataFromDatabase(productDao databasDaos.ProductSqlDao, categoryDao databasDaos.CategorySqlDao, propertyValueDao databasDaos.PropertyValueSqlDao, propertyDao databasDaos.PropertySqlDao, productToPropertyDao databasDaos.ProductToPropertySqlDao, categoryToPropertyDao databasDaos.CategoryToPropertySqlDao) (*[]dbModels.Product, *[]dbModels.Category, *[]dbModels.PropertyValue, *[]dbModels.Property, *[]dbModels.ProductToProperty, *[]dbModels.CategoryToProperty, error) {
	productsFromDb, dbErr := productDao.ReadAllProducts()
	if dbErr != nil {
		return nil, nil, nil, nil, nil, nil, dbErr
	}

	categoriesFromDb, dbErr := categoryDao.ReadAllCategories()
	if dbErr != nil {
		return nil, nil, nil, nil, nil, nil, dbErr
	}

	propertyValuesArr, dbErr := propertyValueDao.ReadAllPropertyValues()
	if dbErr != nil {
		return nil, nil, nil, nil, nil, nil, dbErr
	}

	propertyArr, dbErr := propertyDao.ReadAllProperty()
	if dbErr != nil {
		return nil, nil, nil, nil, nil, nil, dbErr
	}

	productToPropertyArr, dbErr := productToPropertyDao.ReadAllProductToPropertyMapping()
	if dbErr != nil {
		return nil, nil, nil, nil, nil, nil, dbErr
	}

	categoryToPropertyArr, dbErr := categoryToPropertyDao.ReadAllCategoryToPropertyMapping()
	if dbErr != nil {
		return nil, nil, nil, nil, nil, nil, dbErr
	}
	return productsFromDb, categoriesFromDb, propertyValuesArr, propertyArr, productToPropertyArr, categoryToPropertyArr, nil
}

func createDatabaseDaos(db *sql.DB) (databasDaos.ProductSqlDao, databasDaos.CategorySqlDao, databasDaos.PropertySqlDao, databasDaos.PropertyValueSqlDao, databasDaos.ProductToPropertySqlDao, databasDaos.CategoryToPropertySqlDao) {
	productDao := databasDaos.ProductSqlDao{DB: db}
	categoryDao := databasDaos.CategorySqlDao{DB: db}
	propertyDao := databasDaos.PropertySqlDao{DB: db}
	propertyValueDao := databasDaos.PropertyValueSqlDao{DB: db}
	productToPropertyDao := databasDaos.ProductToPropertySqlDao{DB: db}
	categoryToPropertyDao := databasDaos.CategoryToPropertySqlDao{DB: db}
	return productDao, categoryDao, propertyDao, propertyValueDao, productToPropertyDao, categoryToPropertyDao
}

func populateAllProductIdsCache(pool *redis.Pool, productsFromDb *[]dbModels.Product) error {
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

func populateAllCategoryIdsCache(db *sql.DB, pool *redis.Pool) error {
	databaseDao := databasDaos.CategorySqlDao{DB: db}
	cacheDao := cacheDaos.AllCategoryIdsRedisDao{Pool: pool}
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

func populatePropertyValueIdsToProductIdsCache(pool *redis.Pool, propertyValueArr *[]dbModels.PropertyValue, productToPropertyArr *[]dbModels.ProductToProperty) error {
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

func populateCategoryIdToPropertyIdsCache(pool *redis.Pool, categoriesFromDb *[]dbModels.Category, categoryToPropertyFromDbArr *[]dbModels.CategoryToProperty) error {
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

func populatePropertyIdToPropertyValueIdCache(pool *redis.Pool, propertyArr *[]dbModels.Property, propertyValueArr *[]dbModels.PropertyValue) error {
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

func populateProductIdsToPropertyValueIdsCache(pool *redis.Pool, products *[]dbModels.Product, productToPropertyArr *[]dbModels.ProductToProperty) error {
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

func populateProductValueIdsToPropertyIdsCache(pool *redis.Pool, propertyValueArr *[]dbModels.PropertyValue) error {
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
