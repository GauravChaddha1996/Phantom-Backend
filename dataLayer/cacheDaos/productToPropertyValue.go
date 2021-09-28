package cacheDaos

import (
	"github.com/gomodule/redigo/redis"
	"github.com/hashicorp/go-multierror"
	"github.com/spf13/cast"
	"phantom/dataLayer/dbModels"
)

const ProductIdToPropertyValueIdCacheName = "product_id_to_property_value_id_cache"

type ProductToPropertyValueRedisDao struct {
	Pool *redis.Pool
}

func (dao ProductToPropertyValueRedisDao) DeleteWholeCache(products *[]dbModels.Product) error {
	conn := dao.Pool.Get()

	for _, product := range *products {
		_, delErr := conn.Do("DEL", ProductIdToPropertyValueIdCacheName+":"+cast.ToString(product.Id))
		if delErr != nil {
			return delErr
		}
	}
	return nil
}

func (dao ProductToPropertyValueRedisDao) SetProductIdsToPropertyValues(dataArr *[]dbModels.ProductToProperty) error {
	conn := dao.Pool.Get()
	var err error
	for _, productToProperty := range *dataArr {
		key := ProductIdToPropertyValueIdCacheName + ":" + cast.ToString(productToProperty.ProductId)
		_, setErr := conn.Do("SADD", key, productToProperty.ValueId)
		if setErr != nil {
			err = multierror.Append(err, setErr)
		}
	}
	return err
}

func (dao ProductToPropertyValueRedisDao) ReadPropertyValueIdsOfProduct(productId int64) (*[]int64, error) {
	conn := dao.Pool.Get()
	key := ProductIdToPropertyValueIdCacheName + ":" + cast.ToString(productId)

	propertyValueIdsArr, cacheReadErr := redis.Int64s(conn.Do("SMEMBERS", key))
	if cacheReadErr != nil {
		return nil, cacheReadErr
	}
	return &propertyValueIdsArr, nil
}
