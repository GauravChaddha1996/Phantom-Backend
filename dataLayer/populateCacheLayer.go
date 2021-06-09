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
	propertyDao := databasDaos.PropertyDao{DB: db}
	propertyValueDao := databasDaos.PropertyValueDao{DB: db}
	productToPropertyDao := databasDaos.ProductToPropertyDao{DB: db}
	categoryToPropertyDao := databasDaos.CategoryToPropertyDao{DB: db}

	productsFromDb, dbErr := productDao.ReadAllProducts()
	if dbErr != nil {
		return dbErr
	}

	categoriesFromDb, dbErr := categoryDao.ReadAllCategories()
	if dbErr != nil {
		return dbErr
	}

	propertyValuesArr, dbErr := propertyValueDao.ReadAllPropertyValues()
	if dbErr != nil {
		return dbErr
	}

	propertyArr, dbErr := propertyDao.ReadAllProperty()
	if dbErr != nil {
		return dbErr
	}

	productToPropertyArr, dbErr := productToPropertyDao.ReadAllProductToPropertyMapping()
	if dbErr != nil {
		return dbErr
	}

	categoryToPropertyArr, dbErr := categoryToPropertyDao.ReadAllCategoryToPropertyMapping()
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

func populatePropertyValueIdsToProductIdsCache(pool *redis.Pool, propertyValueArr *[]dbModels.PropertyValue, productToPropertyArr *[]dbModels.ProductToProperty) error {
	cacheDao := cacheDaos.PropertyValueToProductDao{Pool: pool}
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
	cacheDao := cacheDaos.CategoryIdToPropertyIdDao{Pool: pool}
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
	cacheDao := cacheDaos.PropertyIdToPropertyValueIdDao{Pool: pool}
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
