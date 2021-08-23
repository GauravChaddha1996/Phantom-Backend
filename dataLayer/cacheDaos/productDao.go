package cacheDaos

import (
	"github.com/gomodule/redigo/redis"
	"phantom/dataLayer/dbModels"
)

const AllProductIdCacheName = "all_product_ids_cache"

type AllProductIdsRedisDao struct {
	Pool *redis.Pool
}

func (dao AllProductIdsRedisDao) DeleteWholeCache() error {
	conn := dao.Pool.Get()
	_, err := conn.Do("DEL", AllProductIdCacheName)
	if err != nil {
		return err
	}
	return nil
}

func (dao AllProductIdsRedisDao) SetProductIdsCache(products *[]dbModels.Product) error {
	conn := dao.Pool.Get()
	for _, product := range *products {
		_, err := conn.Do("SADD", AllProductIdCacheName, product.Id)
		if err != nil {
			return err
		}
	}
	return nil
}

func (dao AllProductIdsRedisDao) ReadAllProductIds() (*[]string, error) {
	conn := dao.Pool.Get()
	productIds, err := redis.Strings(conn.Do("SMEMBER", AllProductIdCacheName, 0, -1))
	if err != nil {
		return nil, err
	}
	return &productIds, nil
}
