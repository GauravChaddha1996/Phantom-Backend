package cacheDaos

import (
	"github.com/gomodule/redigo/redis"
	"github.com/hashicorp/go-multierror"
	"phantom/dataLayer/dbModels"
)

const AllProductIdCacheName = "all_product_ids_cache"

type AllProductIdsRedisDao struct {
	Pool *redis.Pool
}

func (dao AllProductIdsRedisDao) DeleteWholeCache() error {
	conn := dao.Pool.Get()
	_, delErr := conn.Do("DEL", AllProductIdCacheName)
	if delErr != nil {
		return delErr
	}
	return nil
}

func (dao AllProductIdsRedisDao) SetProductIdsCache(products *[]dbModels.Product) error {
	conn := dao.Pool.Get()
	var err error
	for _, product := range *products {
		_, setErr := conn.Do("ZADD", AllProductIdCacheName, product.CreatedAt.Unix(), product.Id)
		if setErr != nil {
			err = multierror.Append(err, setErr)

		}
	}
	return err
}

func (dao AllProductIdsRedisDao) ReadAllProductIds() (*[]int64, error) {
	conn := dao.Pool.Get()
	productIds, readErr := redis.Int64s(conn.Do("ZRANGE", AllProductIdCacheName, 0, -1, "REV"))
	if readErr != nil {
		return nil, readErr
	}
	return &productIds, nil
}

func (dao AllProductIdsRedisDao) ReadFirstNProductIds(n int64) (*[]int64, error) {
	conn := dao.Pool.Get()
	productIds, readErr := redis.Int64s(conn.Do("ZRANGE", AllProductIdCacheName, 0, n-1, "REV"))
	if readErr != nil {
		return nil, readErr
	}
	return &productIds, nil
}

func (dao AllProductIdsRedisDao) ReadRandomProduct() (*int64, error) {
	conn := dao.Pool.Get()
	productId, readErr := redis.Int64(conn.Do("ZRANDMEMBER", AllProductIdCacheName))
	if readErr != nil {
		return nil, readErr
	}
	return &productId, nil
}
