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
		_, err := conn.Do("ZADD", AllProductIdCacheName, product.CreatedAt.Unix(), product.Id)
		if err != nil {
			return err
		}
	}
	return nil
}

func (dao AllProductIdsRedisDao) ReadAllProductIds() (*[]int64, error) {
	conn := dao.Pool.Get()
	productIds, err := redis.Int64s(conn.Do("ZRANGE", AllProductIdCacheName, 0, -1, "REV"))
	if err != nil {
		return nil, err
	}
	return &productIds, nil
}

func (dao AllProductIdsRedisDao) ReadFirstNProductIds(n int64) (*[]int64, error) {
	conn := dao.Pool.Get()
	productIds, err := redis.Int64s(conn.Do("ZRANGE", AllProductIdCacheName, 0, n-1, "REV"))
	if err != nil {
		return nil, err
	}
	return &productIds, nil
}

func (dao AllProductIdsRedisDao) ReadRandomProduct() (*int64, error) {
	conn := dao.Pool.Get()
	productId, err := redis.Int64(conn.Do("ZRANDMEMBER", AllProductIdCacheName))
	if err != nil {
		return nil, err
	}
	return &productId, nil
}
