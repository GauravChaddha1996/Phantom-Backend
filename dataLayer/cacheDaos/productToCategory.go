package cacheDaos

import (
	"github.com/gomodule/redigo/redis"
	"github.com/hashicorp/go-multierror"
	"github.com/spf13/cast"
	"phantom/dataLayer/dbModels"
)

const CategoryIdToProductIdCacheName = "category_id_to_product_id_cache"

type CategoryIdToProductIdRedisDao struct {
	Pool *redis.Pool
}

func (dao CategoryIdToProductIdRedisDao) GetCacheName(id int64) string {
	return CategoryIdToProductIdCacheName + ":" + cast.ToString(id)
}

func (dao CategoryIdToProductIdRedisDao) DeleteWholeCache(categoriesFromDb *[]dbModels.Category) error {
	conn := dao.Pool.Get()
	for _, category := range *categoriesFromDb {
		_, delErr := conn.Do("DEL", dao.GetCacheName(category.Id))
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
		key := dao.GetCacheName(product.CategoryId)
		_, categoryIdSetErr := conn.Do("SADD", key, product.Id)
		if categoryIdSetErr != nil {
			err = multierror.Append(err, categoryIdSetErr)
		}
	}
	return err
}

func (dao CategoryIdToProductIdRedisDao) ReadNProductsOfCategoryId(categoryId *int64, n int) (*[]int64, error) {
	conn := dao.Pool.Get()
	key := dao.GetCacheName(*categoryId)
	productIdsArr, readCacheErr := redis.Int64s(conn.Do("SRANDMEMBER", key, n))
	if readCacheErr != nil {
		return nil, readCacheErr
	}
	return &productIdsArr, nil
}
