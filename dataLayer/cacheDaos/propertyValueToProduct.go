package cacheDaos

import (
	"github.com/gomodule/redigo/redis"
	"github.com/hashicorp/go-multierror"
	"github.com/spf13/cast"
	"phantom/dataLayer/dbModels"
)

const PropertyValueIdToProductIdCacheName = "property_value_id_to_product_id_cache"

type PropertyValueToProductRedisDao struct {
	Pool *redis.Pool
}

func (dao PropertyValueToProductRedisDao) GetCacheName(propertyId int64) string {
	return PropertyValueIdToProductIdCacheName + ":" + cast.ToString(propertyId)
}

func (dao PropertyValueToProductRedisDao) DeleteWholeCache(propertyValues *[]dbModels.PropertyValue) error {
	conn := dao.Pool.Get()

	for _, propertyValue := range *propertyValues {
		_, delErr := conn.Do("DEL", dao.GetCacheName(propertyValue.Id))
		if delErr != nil {
			return delErr
		}
	}
	return nil
}

func (dao PropertyValueToProductRedisDao) SetPropertyValuesToProductIds(dataArr *[]dbModels.ProductToProperty) error {
	conn := dao.Pool.Get()
	var err error
	for _, productToProperty := range *dataArr {
		key := dao.GetCacheName(productToProperty.ValueId)
		_, setErr := conn.Do("SADD", key, productToProperty.ProductId)
		if setErr != nil {
			err = multierror.Append(err, setErr)
		}
	}
	return err
}

func (dao PropertyValueToProductRedisDao) ReadProductIdsOfPropertyValue(valueId int64) (*[]int64, error) {
	conn := dao.Pool.Get()
	key := dao.GetCacheName(valueId)

	productIdsArr, cacheReadErr := redis.Int64s(conn.Do("SMEMBERS", key))
	if cacheReadErr != nil {
		return nil, cacheReadErr
	}
	return &productIdsArr, nil
}
