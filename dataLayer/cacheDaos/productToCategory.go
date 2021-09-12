package cacheDaos

import (
	"github.com/gomodule/redigo/redis"
	"github.com/hashicorp/go-multierror"
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
		_, delErr := conn.Do("DEL", CategoryIdToProductIdCacheName+":"+strconv.FormatInt(category.Id, 10))
		if delErr != nil {
			return delErr
		}
	}
	return nil
}

func (dao CategoryIdToProductIdRedisDao) SetCategoryIdsToProductIdsMap(productArr *[]dbModels.Product) error {
	conn := dao.Pool.Get()
	var err error
	for _, product := range *productArr {
		key := CategoryIdToProductIdCacheName + ":" + strconv.FormatInt(product.CategoryId, 10)
		_, categoryIdSetErr := conn.Do("SADD", key, product.Id)
		if categoryIdSetErr != nil {
			err = multierror.Append(err, categoryIdSetErr)
		}
	}
	return err
}

func (dao CategoryIdToProductIdRedisDao) ReadNProductsOfCategoryId(categoryId *int64, n int) (*[]int64, error) {
	conn := dao.Pool.Get()
	key := CategoryIdToProductIdCacheName + ":" + strconv.FormatInt(*categoryId, 10)
	productIdsArr, readCacheErr := redis.Int64s(conn.Do("SRANDMEMBER", key, n))
	if readCacheErr != nil {
		return nil, readCacheErr
	}
	return &productIdsArr, nil
}
