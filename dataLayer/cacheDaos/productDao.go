package cacheDaos

import (
	"github.com/gomodule/redigo/redis"
	"phantom/dataLayer/dbModels"
)

const AllProductIdSetName = "all_product_ids_cache"

type AllProductIdsCacheDao struct {
	Pool *redis.Pool
}

func (dao AllProductIdsCacheDao) DeleteWholeCache() error {
	conn := dao.Pool.Get()
	_, err := conn.Do("DEL", AllProductIdSetName)
	if err != nil {
		return err
	}
	return nil
}

func (dao AllProductIdsCacheDao) SetProductId(product *dbModels.Product) error {
	conn := dao.Pool.Get()
	_, err := conn.Do("SADD", AllProductIdSetName, product.Id)
	if err != nil {
		return err
	}
	return nil
}

func (dao AllProductIdsCacheDao) ReadAllProductIds() (*[]string, error) {
	conn := dao.Pool.Get()
	productIds, err := redis.Strings(conn.Do("SMEMBERS", AllProductIdSetName))
	if err != nil {
		return nil, err
	}
	return &productIds, nil
}
