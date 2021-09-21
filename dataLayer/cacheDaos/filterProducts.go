package cacheDaos

import (
	"github.com/gomodule/redigo/redis"
)

type FilterProductsDao struct {
	Pool *redis.Pool
}

func (dao FilterProductsDao) FindProductsForFilter(categoryId int64, propertyValueIds []int64) (*[]int64, error) {
	conn := dao.Pool.Get()

	// Find cache name for category id
	categoryIdToProductIdRedisDao := CategoryIdToProductIdRedisDao{dao.Pool}
	categoryIdCacheName := categoryIdToProductIdRedisDao.GetCacheName(categoryId)

	// Find cache names for property value ids
	var propertyValueCacheNameArr []string
	propertyValueToProductRedisDao := PropertyValueToProductRedisDao{dao.Pool}
	for _, propertyValueId := range propertyValueIds {
		propertyValueCacheName := propertyValueToProductRedisDao.GetCacheName(propertyValueId)
		propertyValueCacheNameArr = append(propertyValueCacheNameArr, propertyValueCacheName)
	}

	// Take intersection and find product ids
	finalCacheNameArr := combineAllCacheNames(categoryIdCacheName, propertyValueCacheNameArr)
	productIds, intersectionErr := redis.Int64s(conn.Do("SINTER", finalCacheNameArr...))
	if intersectionErr != nil {
		return nil, intersectionErr
	}
	return &productIds, nil
}

func combineAllCacheNames(
	categoryIdCacheName string, propertyValueCacheNameArr []string,
) []interface{} {
	var finalCacheNameArr []interface{}
	finalCacheNameArr = append(finalCacheNameArr, categoryIdCacheName)
	for _, name := range propertyValueCacheNameArr {
		finalCacheNameArr = append(finalCacheNameArr, name)
	}
	return finalCacheNameArr
}
