package cacheDaos

import (
	"github.com/gomodule/redigo/redis"
	"phantom/dataLayer/dbModels"
	"strconv"
)

const CategoryIdToProductIdCacheName = "category_id_to_product_id_cache"

type CategoryIdToProductIdRedisDao struct {
	Pool *redis.Pool
}

func (dao CategoryIdToProductIdRedisDao) DeleteWholeCache(categoriesFromDb *[]dbModels.Category) error {
	conn := dao.Pool.Get()
	for _, category := range *categoriesFromDb {
		_, err := conn.Do("DEL", CategoryIdToProductIdCacheName+":"+strconv.FormatInt(category.Id, 10))
		if err != nil {
			return err
		}
	}
	return nil
}

func (dao CategoryIdToProductIdRedisDao) SetCategoryIdsToProductIdsMap(productArr *[]dbModels.Product) error {
	conn := dao.Pool.Get()
	for _, product := range *productArr {
		key := CategoryIdToProductIdCacheName + ":" + strconv.FormatInt(product.CategoryId, 10)
		_, categoryIdSetErr := conn.Do("SADD", key, product.Id)
		if categoryIdSetErr != nil {
			return categoryIdSetErr
		}
	}
	return nil
}

func (dao CategoryIdToProductIdRedisDao) ReadAllProductsOfCategoryId(categoryId int64) (*[]int64, error) {
	conn := dao.Pool.Get()
	key := CategoryIdToProductIdCacheName + ":" + strconv.FormatInt(categoryId, 10)
	productIdsArr, readCacheErr := redis.Int64s(conn.Do("SMEMBERS", key))
	if readCacheErr != nil {
		return nil, readCacheErr
	}
	return &productIdsArr, nil
}
